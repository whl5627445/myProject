import time
import grpc
import router_pb2
import router_pb2_grpc


def run():
    # 连接 rpc 服务器
    time1 = time.time()
    channel = grpc.insecure_channel('localhost:50051')
    # 调用 rpc 服务
    stub = router_pb2_grpc.GreeterStub(channel)

    response = stub.ProcessOperation(router_pb2.ProcessOperationRequest(uuid="xqd04", operationName="kill"))
    print(response.msg)
    time2 = time.time()
    print("请求耗时：{}".format(time2 - time1))


if __name__ == '__main__':
    run()
