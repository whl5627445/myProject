import threading
import time
import subprocess
import json
import DyMat
import pandas as pd
from libs.function.xml_input import write_xml
# from config.redis_config import R
from libs.function.run_result_json import add_item_to_json, update_json_item


class OmcRunThread(threading.Thread):
    def __init__(self, request):
        self.state = "init"
        self.absolute_path = request.resultFilePath
        self.uuid = request.uuid
        self.run_pid = None
        self.inputValData = request.inputValData
        self.outputValNames = request.outputValNames
        self.request = request
        self.csv_data = []
        add_item_to_json({"id": self.uuid, "run_states": "init"})
        threading.Thread.__init__(self)

    def run(self):
        run_steps = 0

        keys = list(self.inputValData.keys())  # 获取所有键
        result = []  # 存储结果
        length = len(self.inputValData[keys[0]])  # 获取列表长度，也可以使用len(data["J1.J"])
        for i in range(length):
            item = {}
            for key in keys:
                item[key] = self.inputValData[key][i]
            result.append(item)
        # 进行多轮仿真
        for i in result:
            run_steps += run_steps
            # 修改xml文件
            write_xml(self.absolute_path, i)
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
                                  "run_percentage": run_steps/length,
                                  "run_status": "3",
                                  "run_end_time": int(time.time())})
                break

            else:
                run_result_str = output.decode('utf-8')
                if "successfully" in run_result_str:
                    json_data = {"message": self.request.simulateModelName + "仿真到第{}轮".format(run_steps)}
                    # R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                    update_json_item(self.uuid, {
                        "run_steps": run_steps,
                        "run_percentage": run_steps/length,
                        "run_result_str": run_result_str,
                        "run_status": 4,
                        "run_end_time": int(time.time()), }
                                     )
                    # 从mat中读取数据
                    d = DyMat.DyMatFile(self.absolute_path + "result_res.mat")
                    self.csv_data.append(list(d.abscissa("2", True)))
                    for j in self.outputValNames:
                        self.csv_data.append(list(d.data(j)))
                else:
                    update_json_item(self.uuid,
                                     {"run_steps": run_steps,
                                      "run_percentage": run_steps/length,
                                      "run_result_str": run_result_str,
                                      "run_status": 3,
                                      "run_end_time": int(time.time())}
                                     )
                    break
        df = pd.DataFrame(self.csv_data)
        # 将DataFrame对象保存为CSV文件
        df.to_csv('output.csv', index=False)
        self.state = "stopped"
