import threading
import time
import subprocess
import json
from config.redis_config import R
from libs.function.defs import update_records, new_another_name


class OmcRunThread(threading.Thread):
    def __init__(self, request, absolute_path):
        self.state = "init"
        self.absolute_path = absolute_path
        self.uuid = request.uuid
        self.run_pid = None
        self.request = request
        threading.Thread.__init__(self)

    def run(self):
        # 仿真
        self.state = "running"
        cmd = [self.absolute_path + "result"]
        process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        self.run_pid = process.pid
        # 获取命令行输出结果
        output, error = process.communicate()
        if error:
            update_records(uuid=self.uuid,
                           simulate_status="3",
                           simulate_result_str="编译失败",
                           simulate_start="0",
                           # simulate_start_time=str(self.processStartTime),
                           simulate_end_time=int(time.time())
                           )

        else:
            simulate_result_str = output.decode('utf-8')
            if "successfully" in simulate_result_str:
                json_data = {"message": self.request.simulateModelName + " 模型仿真完成"}
                R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                update_records(uuid=self.uuid,
                               simulate_model_result_path=self.request.resultFilePath,
                               simulate_result_str=simulate_result_str,
                               simulate_status="4",
                               simulate_start="0",
                               # simulate_start_time=str(self.processStartTime),
                               simulate_end_time=int(time.time()),
                               another_name=new_another_name(self.request.userName,
                                                             self.request.simulateModelName,
                                                             self.request.userSpaceId)
                               )

            else:
                update_records(uuid=self.uuid,
                               simulate_result_str=simulate_result_str,
                               simulate_status="3",
                               simulate_start="0",
                               # simulate_start_time=str(self.processStartTime),
                               simulate_end_time=int(time.time())
                               )
        self.state = "stopped"
