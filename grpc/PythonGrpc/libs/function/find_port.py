import socket
from libs.function.grpc_log import log


def checkPort(port):
    res = False
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    try:
        s.bind(('localhost', port))
    except socket.error as e:
        log.info("(OMC)端口{}已被占用,查找可用端口。".format(port))
        res = True

    s.close()
    return res


def findPort(startPort):
    while checkPort(startPort):
        startPort = startPort + 1
    return startPort
