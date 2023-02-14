from concurrent import futures
import os.path
from multiprocessing import Manager
from libs.defs import getSate, suspendProcess, resumeProcess, killProcess
from libs.FmpySimulation import MyProcess
import threading
import zarr
from db_config.config import Session, YssimSimulateRecords
import time
import grpc
import router_pb2
import router_pb2_grpc

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
            if not os.path.exists(request.fmuPath):
                return router_pb2.FmuSimulationReply(log="No such file or directory!")
                # 最大任务数
            if len(processList) < 10:
                processOne = MyProcess(request, managerResDict)
                processList.append(processOne)
                return router_pb2.FmuSimulationReply(log="Task submitted successfully.")
            else:
                return router_pb2.FmuSimulationReply(
                    log="The total number of system tasks has exceeded 4. Please wait and request again!")

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
            with Session() as session:
                processDetails = session.query(YssimSimulateRecords).filter(YssimSimulateRecords.id == request.uuid).first()
            if processDetails is None:
                return router_pb2.GetResultReply(log="not found in db")
            if processDetails.simulate_model_result_path:
                if processDetails.simulate_status in ["3", "4"]:
                    res = zarr.load(processDetails.simulate_model_result_path)
                    if (res is not None) and (request.variable in res.dtype.names):
                        data = res[request.variable].tolist()
                        return router_pb2.GetResultReply(data=data)
                    else:
                        return router_pb2.GetResultReply(log="not found var(end)")
                else:
                    if request.uuid in managerResDict:
                        zarr.save(processDetails.simulate_model_result_path, managerResDict[request.uuid])
                        res = zarr.load(processDetails.simulate_model_result_path)
                        if (res.size != 0) and (request.variable in res.dtype.names):
                            data = res[request.variable].tolist()
                            return router_pb2.GetResultReply(data=data)
                        else:
                            return router_pb2.GetResultReply(log="not found var(running)")
                    else:
                        return router_pb2.GetResultReply(log="not found dict in mana")
            else:
                return router_pb2.GetResultReply(log="not found resPath")

        # 进程操作
        def ProcessOperation(self, request, context):
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


    def action():
        while True:
            time.sleep(2)
            for i in processList:
                processState = getSate(i.__repr__())
                if processState in ["closed", "unknown", "stopped"]:
                    processList.remove(i)
            for i in processList[:4]:
                if not i.is_alive():
                    i.start()


    startProcessList = threading.Thread(target=action)
    startProcessList.start()

    # 启动 rpc 服务
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    router_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('0.0.0.0:5001')
    print("服务开启成功！")
    server.start()
    try:
        while True:
            time.sleep(24 * 3600)  # one day in seconds
    except KeyboardInterrupt:
        server.stop(0)
