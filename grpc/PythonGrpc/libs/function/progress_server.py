import re
import socket  # 导入socket库
import threading
import time
from libs.function.grpc_log import log
from libs.function.find_port import findPort
from libs.function.defs import update_simulate_records, find_max_number


class TcpServer(threading.Thread):  # TCP服务
    def __init__(self, socket_port, db_id):
        super().__init__()
        self.socket_port = socket_port
        log.info("socket服务端口为:{}".format(socket_port))
        self.db_id = db_id
        self.socket_host = ''
        self.socket_addr = (self.socket_host, self.socket_port)
        self.buff_size = 2048  # 定义一次从socket缓冲区最多读入2048个字节
        self.max_listen = 1
        self.stop_flag = True
        # AF_INET表示socket网络层使用IP协议，SOCK_STREAM表示socket传输层使用tcp协议
        self.s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.conn = None
        self.number_of_requests = 0
        self.percentage = [0]

    def update_db(self):
        percent_flag = 0
        while self.stop_flag:
            # 将仿真百分数转换为30-100之间
            percent = int((self.percentage[-1]/100)*70+30)
            if percent != percent_flag:
                percent_flag = percent
                log.info("(OMC)更新数据库进度："+str(percent))
                update_simulate_records(uuid=self.db_id, percentage=percent)
                # log.info("更新数据库进度为:{}".format(percent))
                time.sleep(0.5)

    def stop(self):
        log.info("关闭读取进度的socket,进度更新到:"+str(self.percentage[-1]))
        self.stop_flag = False
        self.s.close()
        if self.conn:
            self.conn.close()
        else:
            print(type(self.conn))

    def run(self):
        t = threading.Thread(target=self.update_db)
        t.start()
        # 设置超时时间为600秒
        self.s.settimeout(600)

        # 绑定服务器地址和端口
        self.s.bind(self.socket_addr)
        # 启动服务监听

        self.s.listen(self.max_listen)
        log.info('等待仿真进度传入!')

        while self.stop_flag:
            # 等待客户端连接请求,获取connSock
            try:
                self.conn, addr = self.s.accept()
            except Exception as e:
                log.info("等待客户端连接请求出错:" + str(e))
                return
            log.info('客户端:{} 接入！！！'.format(addr))
            # with conn:
            while self.stop_flag:
                self.number_of_requests += 1
                # log.info('接收请求信息')
                # 接收请求信息
                try:
                    data = self.conn.recv(self.buff_size)  # 读取的数据一定是bytes类型，需要解码成字符串类型
                    if not data:
                        break
                    info = data.decode()
                    percent = find_max_number(info)
                    self.percentage.append(percent/100)

                    # 发送请求数据
                    self.conn.send(f'服务端接收到信息{info}'.encode())
                    # log.info('发送返回完毕！！！')
                except Exception as e:
                    log.info("读取进度socket断开："+str(e))
                    return


# if __name__ == '__main__':
#     port_ = findPort(50000)
#     print(port_)
#     tcpServer = TcpServer(port_)
#     tcpServer.start()
    # time.sleep(325)
    # print("stop")
    # tcpServer.stop()
