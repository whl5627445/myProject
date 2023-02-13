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
    response = stub.GetResult(router_pb2.GetResultRequest(uuid="xqd04", variable="time"))
    print(response.log)
    print(response.data)
    print(len(response.data))
    time2 = time.time()
    print("请求耗时：{}".format(time2 - time1))


if __name__ == '__main__':
    run()
