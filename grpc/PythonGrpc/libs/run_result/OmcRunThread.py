import threading
import os
import subprocess
import json
import DyMat
import pandas as pd
from libs.function.xml_input import write_xml
from config.redis_config import R
from libs.function.run_result_json import update_item_to_json, delete_item_from_json
from libs.function.defs import update_app_pages_records


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
        update_item_to_json(self.uuid, {"id": self.uuid, "run_states": "init", })
        threading.Thread.__init__(self)

    def run(self):

        run_steps = 0
        keys = list(self.inputValData.keys())  # 获取所有键
        input_data = []  # 存储结果输入字典,[{"name1":1,"name2":2,},{"name1":11,"name2":22,}]
        print(self.inputValData[keys[0]])
        print("type::", type(self.inputValData[keys[0]].inputObjList))
        input_data_length = len(self.inputValData[keys[0]].inputObjList)  # 获取列表长度，也可以使用len(data["J1.J"])
        for i in range(input_data_length):
            item = {}
            for key in keys:
                item[key] = self.inputValData[key].inputObjList[i]
            input_data.append(item)
        # 进行多轮仿真
        for i in input_data:
            run_steps += 1
            print("进行第{}轮仿真".format(run_steps))
            # 修改xml文件
            print("i::: ", i)
            if write_xml(r"/home/simtek/code/" + self.absolute_path, i):
                # 解析文件失败
                break

            # 运行可执行文件result
            self.state = "running"
            cmd = [r"/home/simtek/code/" + self.absolute_path + "result"]
            process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            self.run_pid = process.pid
            # 获取命令行输出结果
            output, error = process.communicate()
            if error:
                break

            else:
                run_result_str = output.decode('utf-8')
                if "successfully" in run_result_str:
                    json_data = {"message": self.request.simulateModelName + "仿真到第{}轮".format(run_steps)}
                    R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))

                    # 从mat中读取数据
                    d = DyMat.DyMatFile(r"/home/simtek/code/" + self.absolute_path + "result_res.mat")
                    # if run_steps == 1:
                    #     self.csv_data.append(list(d.abscissa("2", True)))
                    # for j in self.outputValNames:
                    #     self.csv_data.append(list(d.data(j)))

                    dictCsv = {"time": list(d.abscissa("2", True))
                               # "time2": list(d.abscissa("1", True))
                               }

                    dictCsv["time"] = dictCsv["time"][:500]

                    for j in self.outputValNames:
                        dictCsv[j] = list(d.data(j))[:500]
                    print("self.outputValNames",self.outputValNames)
                    df = pd.DataFrame(pd.DataFrame.from_dict(dictCsv, orient='index').values.T,
                                      columns=list(dictCsv.keys()))
                    if input_data_length == 1:
                        pass
                        # 单次仿真的结果
                        # if not os.path.exists(r"/home/simtek/code/" + self.absolute_path + 'single_output'):
                        #     # 创建新的文件夹
                        #     os.mkdir(r"/home/simtek/code/" + self.absolute_path + 'single_output')
                        # df.to_csv(r"/home/simtek/code/" + self.absolute_path + 'single_output/Single_output.csv',
                        #           index=False,
                        #           encoding='utf-8')
                    else:
                        # 多轮仿真每轮一个scv文件
                        if not os.path.exists(r"/home/simtek/code/" + self.absolute_path + 'mul_output'):
                            # 创建新的文件夹
                            os.mkdir(r"/home/simtek/code/" + self.absolute_path + 'mul_output')
                        df.to_csv(r"/home/simtek/code/" + self.absolute_path + 'mul_output/Mul_output_{}.csv'.format(
                            run_steps),
                                  index=False,
                                  encoding='utf-8')
                else:
                    break

        if run_steps == input_data_length:  # 如果每轮都成功
            if input_data_length == 1:
                # 更新数据库
                update_app_pages_records(self.request.pageId,
                                         single_simulation_result_path=self.absolute_path + 'single_output/')
            else:
                # 更新数据库
                update_app_pages_records(self.request.pageId,
                                         multi_simulation_results_path=self.absolute_path + 'mul_output/')
        self.state = "stopped"
        delete_item_from_json(self.uuid)
