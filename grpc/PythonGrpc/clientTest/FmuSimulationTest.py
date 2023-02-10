import time
import grpc
import router_pb2
import router_pb2_grpc


def run():
    # 连接 rpc 服务器
    channel = grpc.insecure_channel('localhost:50051')
    # 调用 rpc 服务
    stub = router_pb2_grpc.GreeterStub(channel)
    # 测试仿真接口
    time1 = time.time()
    response = stub.FmuSimulation(router_pb2.FmuSimulationRequest(uuid="xqd04",
                                                                  startTime=0,
                                                                  stopTime=1000,
                                                                  fmuPath="./static/xuqingda/VTM_FWD_CS/202301130927/VTM_FWD_CS.fmu",
                                                                  resPath="./static/xuqingda/VTM_FWD_CS/zarRes/xqd04res.zarr",
                                                                  params={"SOC_ini": "101"
                                                                  },
                                                                  outputInterval=1))
    time2 = time.time()
    print("请求耗时：{}".format(time2 - time1))
    print("log: " + response.log)


if __name__ == '__main__':
    run()
