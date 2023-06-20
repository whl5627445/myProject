import copy
from concurrent import futures
import os.path
import pandas as pd
from libs.function.process_operation import kill_py_omc_process

from libs.simulation.OmcSimulationThread import OmcSimulation
from libs.translate.OmcTranslateThread import OmcTranslateThread
from libs.run_result.OmcRunThread import OmcRunThread

from libs.simulation.DmSimulationThread import DmSimulation
from libs.translate.DmcTranslateThread import DmTranslateThread
from libs.run_result.DmRunThread import DmRunThread

from libs.function.find_port import findPort
from libs.function.grpc_log import log
from libs.function.init import initOmTask, initDmTask
import threading
import zarr
from config.db_config import Session, YssimSimulateRecords
import time
import grpc
import router_pb2
import router_pb2_grpc
import DyMat
from fmpy import write_csv
import configparser

config = configparser.ConfigParser()
config.read('./config/grpc_config.ini')
start_port = int(config['app']['start_port'])
max_simulation_num = int(config['app']['max_simulation_num'])

adsPath = "/home/simtek/code/"
# used_ports = []
# adsPath = "/home/xuqingda/GolandProjects/YssimGoService/"

if __name__ == '__main__':
    # omc线程列表，用于保存每个线程对象
    OmSimulationThreadList = []
    # dymola线程列表，用于保存每个线程对象
    DmSimulationThreadList = []
    # 接收仿真任务队列
    # 初始化omc任务队列
    omcTaskList = initOmTask()
    # 初始化dy任务队列
    dymolaTaskList = initDmTask()
    taskMarkDict = {}


    # 实现 proto 文件中定义的 GreeterServicer
    class Greeter(router_pb2_grpc.GreeterServicer):
        # 实现 proto 文件中定义的 rpc 调用
        # 仿真接口
        def SubmitTask(self, request, context):

            newRequest = copy.deepcopy(request)
            data = newRequest
            if data.simulateType == "OM":
                # 如果是OM仿真，将仿真请求体放到omcTaskList中
                omcTaskList.append(data)
            elif data.simulateType == "DM":
                # 如果是DM仿真,将仿真请求体放到dymolaTaskList中
                dymolaTaskList.append(data)
            return router_pb2.SubmitTaskReply(ok=True,
                                              msg="Task submitted successfully."
                                              )

        def ProcessOperation(self, request, context):
            log.info("ProcessOperation被调用。")
            multiprocessing_id = request.uuid
            operationName = request.operationName
            if operationName == "kill":
                if request.simulate_type == "OM":
                    for i in range(len(omcTaskList)):
                        if omcTaskList[i].uuid == multiprocessing_id:
                            omcTaskList.pop(i)
                            break
                if request.simulate_type == "DM":
                    for i in range(len(dymolaTaskList)):
                        if dymolaTaskList[i].uuid == multiprocessing_id:
                            dymolaTaskList.pop(i)
                killProcessReply = kill_py_omc_process(multiprocessing_id, OmSimulationThreadList,
                                                       request.simulate_type, taskMarkDict)
                # del taskMarkDict[request.userName]
                # del taskMarkDict[request.userName]
                return router_pb2.ProcessOperationReply(msg=killProcessReply["msg"])
            # if operationName == "suspend":
            #     suspendProcessReply = suspend_process(multiprocessing_id, OmSimulationThreadList)
            #     return router_pb2.ProcessOperationReply(msg=suspendProcessReply["msg"])
            # if operationName == "resume":
            #     resumeProcessReply = resume_process(multiprocessing_id, OmSimulationThreadList)
            #     return router_pb2.ProcessOperationReply(msg=resumeProcessReply["msg"])
            return router_pb2.ProcessOperationReply(msg="Unknown operation!")

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

        # 获取变量结果 单个mat文件单个变量mat
        def GetResult(self, request, context):
            log.info("GetResult被调用。")
            d = DyMat.DyMatFile(r"/home/simtek/code/" + request.path)
            try:
                data_time = list(d.abscissa("2", True))
                if request.variable == "time":
                    data = data_time
                else:
                    d_data = list(d.data(request.variable))
                    if len(d_data) == 2 and d_data[0] == d_data[1]:
                        d_data = [d_data[0] for i in range(len(data_time))]
                    data = d_data
                return router_pb2.GetResultReply(data=data, log="true")
            except Exception as e:
                return router_pb2.GetResultReply(data=[], log=str(e))

        # 单个记录，多个变量zarr
        def ReadSimulationResult(self, request, context):
            log.info("ReadSimulationResult被调用。")
            res = zarr.load(adsPath + request.resultPath)
            if res is not None:
                resList = [res["time"].tolist()]
                for i in request.Vars:
                    resList.append(res[i].tolist())
                merged_list = list(zip(*resList))
                result_list = [list(item) for item in merged_list]
                sim_result = router_pb2.ReadSimulationResultReply()
                for row_data in result_list:
                    row = router_pb2.ReadSimulationResultReply.ele()
                    row.row.extend(row_data)
                    sim_result.data.append(row)
                sim_result.ok = True
                return sim_result  # 返回 protobuf 消息对象
            else:
                return router_pb2.SaveFilterResultToCsvReply(ok=False)

        def CheckVarExist(self, request, context):
            log.info("CheckVarExist被调用")
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


    class SimulationThread(threading.Thread):
        def __int__(self):
            threading.Thread.__init__(self)
            pass

        def run(self):
            log.info("仿真任务执行线程启动!")
            while True:
                time.sleep(1)
                if len(OmSimulationThreadList) > 0:
                    log.info("(OMC)正在运行的任务数：{}".format(len(OmSimulationThreadList)))
                    log.info("(OMC)正在运行的任务：" + str(
                        [{j.request.simulateModelName: j.state, "user_name": j.request.userName} for j in
                         OmSimulationThreadList]))
                if len(omcTaskList) > 0:
                    log.info("(OMC)正在排队的任务数：{}".format(len(omcTaskList)))
                    log.info("(OMC)正在排队的任务：" + str(
                        [{"model_name": j.simulateModelName, "user_name": j.userName} for j in
                         omcTaskList]))

                for i in OmSimulationThreadList:
                    if i.state == "stopped":
                        log.info("(OMC)" + i.request.simulateModelName + "仿真结束,线程关闭。")
                        OmSimulationThreadList.remove(i)
                        del taskMarkDict[i.request.userName]
                        # del i
                for i in DmSimulationThreadList:
                    if i.state == "stopped":
                        log.info("(Dymola)" + i.request.simulateModelName + "仿真结束,线程关闭。")
                        DmSimulationThreadList.remove(i)
                        del taskMarkDict[i.request.userName]
                if len(DmSimulationThreadList) > 0:
                    log.info("(Dymola)正在运行的任务数：{}".format(len(DmSimulationThreadList)))
                    log.info("(Dymola)正在运行的任务：" + str(
                        [{j.request.simulateModelName: j.state, "user_name": j.request.userName} for j in
                         DmSimulationThreadList]))
                if len(dymolaTaskList) > 0:
                    log.info("(Dymola)未执行任务队列剩余数量：{}".format(len(dymolaTaskList)))
                    log.info("(Dymola)正在排队的任务：" + str(
                        [{"model_name": j.simulateModelName, "user_name": j.userName} for j in
                         dymolaTaskList]))
                i = 0
                while i < len(omcTaskList):
                    userName = omcTaskList[i].userName
                    if userName not in taskMarkDict:
                        data = omcTaskList.pop(i)
                        if data.taskType == "simulate":
                            # 找到空闲的端口号
                            port = findPort(start_port)
                            om_threading = OmcSimulation(data, port)
                            om_threading.start()
                            OmSimulationThreadList.append(om_threading)
                        if data.taskType == "translate":
                            # 找到空闲的端口号
                            port = findPort(start_port)
                            om_threading = OmcTranslateThread(data, port)
                            om_threading.start()
                            OmSimulationThreadList.append(om_threading)
                        if data.taskType == "run":
                            om_threading = OmcRunThread(data)
                            om_threading.start()
                            OmSimulationThreadList.append(om_threading)
                        taskMarkDict[userName] = True
                    i += 1
                while i < len(dymolaTaskList):
                    userName = dymolaTaskList[i].userName
                    if userName not in taskMarkDict:
                        data = dymolaTaskList.pop(i)
                        if data.taskType == "simulate":
                            dm_threading = DmSimulation(data)
                            dm_threading.start()
                            DmSimulationThreadList.append(dm_threading)
                        if data.taskType == "translate":
                            dm_threading = DmTranslateThread(data)
                            dm_threading.start()
                            DmSimulationThreadList.append(dm_threading)
                        if data.taskType == "run":
                            dm_threading = DmRunThread(data)
                            dm_threading.start()
                            DmSimulationThreadList.append(dm_threading)
                        taskMarkDict[userName] = True
                    i += 1
                # if len(OmSimulationThreadList) < max_simulation_num and len(omcTaskList) > 0:
                #     data = omcTaskList.pop(0)
                #     if data.taskType == "simulate":
                #         # 找到空闲的端口号
                #         port = findPort(start_port)
                #         om_threading = OmcSimulation(data, port)
                #         om_threading.start()
                #         OmSimulationThreadList.append(om_threading)
                #     if data.taskType == "translate":
                #         # 找到空闲的端口号
                #         port = findPort(start_port)
                #         om_threading = OmcTranslateThread(data, port)
                #         om_threading.start()
                #         OmSimulationThreadList.append(om_threading)
                #     if data.taskType == "run":
                #         om_threading = OmcRunThread(data)
                #         om_threading.start()
                #         OmSimulationThreadList.append(om_threading)

                # if len(DmSimulationThreadList) < max_simulation_num and len(dymolaTaskList) > 0:
                #     data = dymolaTaskList.pop(0)
                #     if data.taskType == "simulate":
                #         dm_threading = DmSimulation(data)
                #         dm_threading.start()
                #         DmSimulationThreadList.append(dm_threading)
                #     if data.taskType == "translate":
                #         dm_threading = DmTranslateThread(data)
                #         dm_threading.start()
                #         DmSimulationThreadList.append(dm_threading)
                #     if data.taskType == "run":
                #         dm_threading = DmRunThread(data)
                #         dm_threading.start()
                #         DmSimulationThreadList.append(dm_threading)


    # 启动 rpc 服务
    simulation_obj = SimulationThread()
    simulation_obj.start()
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    router_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('0.0.0.0:50051')

    log.info("服务开启成功！0.0.0.0:50051")
    server.start()

    try:
        while True:
            time.sleep(24 * 3600)  # one day in seconds

    except KeyboardInterrupt:
        server.stop(0)
