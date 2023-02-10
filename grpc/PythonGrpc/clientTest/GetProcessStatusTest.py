import time
import grpc
import router_pb2
import router_pb2_grpc


def run():
    time1 = time.time()
    # 连接 rpc 服务器

    channel = grpc.insecure_channel('localhost:50051')
    # 调用 rpc 服务
    stub = router_pb2_grpc.GreeterStub(channel)

    # 测试获取单个进程接口
    response = stub.GetProcessStatus(router_pb2.GetProcessStatusRequest(uuid="xqd04"))
    print(response.progress)
    print(response.state)
    print( response.processStartTime)
    time2 = time.time()
    print("请求耗时：{}".format(time2 - time1))


if __name__ == '__main__':
    run()
