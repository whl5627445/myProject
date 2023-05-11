import threading
import time
import subprocess
import json
from libs.function.xml_input import write_xml
from config.redis_config import R
from libs.function.run_result_json import add_item_to_json, update_json_item


class OmcRunThread(threading.Thread):
    def __init__(self, request):
        self.state = "init"
        self.absolute_path = request.resultFilePath
        self.uuid = request.uuid
        self.run_pid = None
        self.inputValData = request.inputObjList
        self.outputValNames = request.outputValNames
        self.request = request
        add_item_to_json({"id": self.uuid, "run_states": "init"})
        threading.Thread.__init__(self)

    def run(self):
        run_steps = 0
        for i in self.inputValData:
            run_steps += run_steps
            # 修改xml文件
            write_xml(self.absolute_path, i.inputValData)
            # 运行可执行文件result
            self.state = "running"
            cmd = [self.absolute_path + "result"]
            process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            self.run_pid = process.pid
            # 获取命令行输出结果
            output, error = process.communicate()
            if error:
                update_json_item(self.uuid,
                                 {"run_steps": run_steps,
                                  "run_status": "3",
                                  "run_end_time": int(time.time())})

            else:
                run_result_str = output.decode('utf-8')
                if "successfully" in run_result_str:
                    json_data = {"message": self.request.simulateModelName + "仿真到第{}轮".format(run_steps)}
                    R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                    update_json_item(self.uuid, {
                        "run_steps": run_steps,
                        "run_result_str": run_result_str,
                        "run_status": 4,
                        "run_end_time": int(time.time()), }
                                     )

                else:
                    update_json_item(self.uuid,
                                     {"run_steps": run_steps,
                                      "run_result_str": run_result_str,
                                      "run_status": 3,
                                      "run_end_time": int(time.time())}
                                     )
        self.state = "stopped"
