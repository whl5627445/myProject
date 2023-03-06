from concurrent import futures
import os.path
import pandas as pd
from multiprocessing import Manager
from libs.defs import getSate, suspendProcess, resumeProcess, killProcess, initOmc, buildFMU
from libs.FmpySimulation import MyProcess
import threading
import zarr
from config.db_config import Session, YssimSimulateRecords
import time
import grpc
import router_pb2
import router_pb2_grpc
import DyMat
from fmpy import write_csv

adsPath = "/home/simtek/code/"

if __name__ == '__main__':
    # 进程通信，用于保存每个进程的仿真结果
    managerResDict = Manager().dict()
    managerLock = Manager().Lock()
    # 进程列表，用于保存每个进程对象
    processList = []

    # 实现 proto 文件中定义的 GreeterServicer
    class Greeter(router_pb2_grpc.GreeterServicer):
        # 实现 proto 文件中定义的 rpc 调用
        # 仿真接口
        def FmuSimulation(self, request, context):

            if not os.path.exists(adsPath + request.moPath):
                return router_pb2.FmuSimulationReply(log="No such file or directory!")
                # 最大任务数
            if not buildFMU(request.moPath, request.className, request.userName, request.resPath):
                return router_pb2.FmuSimulationReply(log="buildFMU error!")
            if len(processList) < 10:
                processOne = MyProcess(request, managerResDict)
                processList.append(processOne)
                return router_pb2.FmuSimulationReply(log="Task submitted successfully.")
            else:
                return router_pb2.FmuSimulationReply(
                    log="The total number of system tasks has exceeded 8. Please wait and request again!")

        # 获取某个进程状态信息
        def GetProcessStatus(self, request, context):
            uuid = request.uuid
            with Session() as session:
                processDetails = session.query(YssimSimulateRecords).filter(YssimSimulateRecords.id == uuid).first()
            if processDetails:
                return router_pb2.GetProcessStatusReply(log=str(processDetails.simulate_result_str),
                                                        exception=int(processDetails.simulate_status),
                                                        progress=int(processDetails.simulate_status),
                                                        processStartTime=str(processDetails.simulate_start_time),
                                                        state=str(processDetails.simulate_status),
                                                        processRunTime=str(processDetails.simulate_end_time),
                                                        resPath=str(processDetails.simulate_model_result_path)
                                                        )

            else:
                return router_pb2.GetProcessStatusReply(log="not found",
                                                        exception=1,
                                                        progress=0,
                                                        processStartTime="",
                                                        state="unknown",
                                                        processRunTime="",
                                                        resPath="")

        # 获取所有进程的数量
        def GerAllProcessNumber(self, request, context):
            running_num = 0
            pending_num = 0
            for i in processList:
                processState = getSate(i.__repr__())
                if processState == 'initial':
                    pending_num += 1
                if processState == 'started':
                    running_num += 1
            return router_pb2.GerAllProcessNumberReply(totalTasks=len(processList),
                                                       numOfRunningProcess=running_num,
                                                       numOfPendingProcess=pending_num
                                                       )

        # 获取变量结果
        def GetResult(self, request, context):
            print("GetResult被调用。")
            with Session() as session:
                processDetails = session.query(YssimSimulateRecords).filter(
                    YssimSimulateRecords.id == request.uuid).first()
            if processDetails is None:
                return router_pb2.GetResultReply(data=[], log="not found in db")
            if processDetails.simulate_model_result_path:
                zarrPath = adsPath + processDetails.simulate_model_result_path + "zarr_res.zarr"
                if processDetails.simulate_status in ["3", "4"]:
                    res = zarr.load(zarrPath)
                    if (res is not None) and (request.variable in res.dtype.names):
                        data = res[request.variable].tolist()
                        return router_pb2.GetResultReply(data=data, log="true")
                    else:
                        return router_pb2.GetResultReply(data=[], log="not found var(end)")
                else:
                    if request.uuid in managerResDict:
                        zarr.save(zarrPath, managerResDict[request.uuid])
                        res = zarr.load(zarrPath)
                        if (res.size != 0) and (request.variable in res.dtype.names):
                            data = res[request.variable].tolist()
                            return router_pb2.GetResultReply(data=data, log="true")
                        else:
                            return router_pb2.GetResultReply(data=[], log="not found var(running)")
                    else:
                        return router_pb2.GetResultReply(data=[], log="not found dict in mana")
            else:
                return router_pb2.GetResultReply(data=[], log="not found resPath")

        # 进程操作
        def ProcessOperation(self, request, context):
            print("ProcessOperation被调用。")
            multiprocessing_id = request.uuid
            operationName = request.operationName
            if operationName == "kill":
                killProcessReply = killProcess(multiprocessing_id, processList, managerResDict)
                return router_pb2.ProcessOperationReply(msg=killProcessReply["msg"])
            if operationName == "suspend":
                suspendProcessReply = suspendProcess(multiprocessing_id, processList)
                return router_pb2.ProcessOperationReply(msg=suspendProcessReply["msg"])
            if operationName == "resume":
                resumeProcessReply = resumeProcess(multiprocessing_id, processList)
                return router_pb2.ProcessOperationReply(msg=resumeProcessReply["msg"])
            return router_pb2.ProcessOperationReply(msg="Unknown operation!")

        def SaveFilterResultToCsv(self, request, context):
            print("SaveFilterResultToCsv被调用。")
            res = zarr.load(adsPath + request.resultPath)
            resDict = {"time": res["time"].tolist()}
            if res is not None:
                for i in request.Vars:
                    resDict[i] = res[i].tolist()
                df = pd.DataFrame(resDict)
                dirnamePath = os.path.dirname(adsPath + request.newFileName)
                if not os.path.exists(dirnamePath):
                    print("创建csv路径", dirnamePath)
                    os.makedirs(dirnamePath)
                df.to_csv(adsPath + request.newFileName, index=False, encoding='utf-8')
                return router_pb2.SaveFilterResultToCsvReply(ok=True)
            else:
                return router_pb2.SaveFilterResultToCsvReply(ok=False)

        def MatToCsv(self, request, context):
            print("MatToCsv被调用。")
            try:
                d = DyMat.DyMatFile(adsPath + request.matPath)
                namesList = list(d.names())
                dictCsv = {"time1": list(d.abscissa("2", True)),
                           "time2": list(d.abscissa("1", True))
                           }
                if len(dictCsv["time1"]) > 1000:
                    dictCsv["time1"] = dictCsv["time1"][:1000]
                    for i in namesList:
                        dictCsv[i] = list(d.data(i))[:1000]
                else:
                    for i in namesList:
                        dictCsv[i] = list(d.data(i))
                df = pd.DataFrame(pd.DataFrame.from_dict(dictCsv, orient='index').values.T,
                                  columns=list(dictCsv.keys()))
                df.to_csv(os.path.dirname(adsPath + request.matPath) + "/result_res.csv", index=False, encoding='utf-8')
            except Exception as e:
                print(e)
                return router_pb2.MatToCsvReply(ok=False)
            else:
                return router_pb2.MatToCsvReply(ok=True)

        def ZarrToCsv(self, request, context):
            print("ZarrToCsv被调用")
            try:
                d = zarr.load(adsPath + request.zarrPath)
                if d.shape[0] > 1000:
                    d = d[:1000]
                write_csv(os.path.dirname(adsPath + request.zarrPath) + "/result_res.csv", d)
            except Exception as e:
                print(e)
                return router_pb2.ZarrToCsvReply(ok=False)
            else:
                return router_pb2.ZarrToCsvReply(ok=True)

        def CheckVarExist(self, request, context):
            print("CheckVarExist被调用")
            resMap = {}
            zarrPath = os.path.dirname(adsPath + request.Path) + "/zarr_res.zarr"

            if os.path.exists(zarrPath):
                d = zarr.load(zarrPath)

                if d is not None:
                    for i in request.Names:
                        if i in d.dtype.names:
                            resMap[i] = True
                        else:
                            resMap[i] = False
            else:
                for i in request.Names:
                    resMap[i] = True
            return router_pb2.CheckVarExistReply(Res=resMap)


    def action():
        while True:
            time.sleep(1)
            for i in processList:
                processState = getSate(i.__repr__())
                if processState in ["closed", "unknown", "stopped"]:
                    processList.remove(i)
            for i in processList[:8]:
                if not i.is_alive():
                    i.start()


    startProcessList = threading.Thread(target=action)
    startProcessList.start()

    # 启动 rpc 服务
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    router_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('0.0.0.0:50051')
    initOmc()
    print("服务开启成功！")
    server.start()
    try:
        while True:
            time.sleep(24 * 3600)  # one day in seconds
    except KeyboardInterrupt:
        server.stop(0)
