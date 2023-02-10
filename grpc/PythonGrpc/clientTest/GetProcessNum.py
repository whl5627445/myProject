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
    # 测试获取进程状态数量接口
    response = stub.GerAllProcessNumber(router_pb2.GerAllProcessNumberRequest())
    print(response.totalTasks)
    time2 = time.time()
    print("请求耗时：{}".format(time2-time1))

if __name__ == '__main__':
    run()
