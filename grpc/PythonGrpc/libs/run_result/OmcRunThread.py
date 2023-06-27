import threading
import os
import subprocess
import json
import time

import DyMat
import pandas as pd
from libs.function.xml_input import write_xml
from libs.function.run_result_json import update_item_to_json, delete_item_from_json
from libs.function.defs import update_app_pages_records, omc_convert_dict_to_list, update_app_spaces_records,page_release_component_freeze
from libs.function.grpc_log import log
import shutil


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
        self.input_data = omc_convert_dict_to_list(self.inputValData, self.request.pageId)

        if len(self.input_data) == 1:  # 仿真任务
            update_app_pages_records(self.request.pageId, simulate_state=1)
        else:  # 发布任务
            update_app_pages_records(self.request.pageId, release_state=1)
        threading.Thread.__init__(self)

    def run(self):
        message = ""
        run_steps = 0
        if len(self.input_data) == 1:  # 仿真任务
            update_app_pages_records(self.request.pageId, simulate_state=2)
        else:  # 发布任务
            update_app_pages_records(self.request.pageId, release_state=2)
        # 进行多轮仿真
        if self.request.mulResultPath:
            mul_output_path = r"/home/simtek/code/" + self.request.mulResultPath
            if os.path.exists(mul_output_path):
                shutil.rmtree(mul_output_path)
            # 创建新的文件夹
            os.mkdir(mul_output_path)
        else:
            log.info("(OMC)mulResultPath路径不存在")
            if len(self.input_data) == 1:
                # 更新数据库
                update_app_pages_records(self.request.pageId, simulate_state=3)
            else:
                # 更新数据库
                update_app_pages_records(self.request.pageId, release_state=3)
            self.state = "stopped"
            update_app_pages_records(self.request.pageId, update_time=time.time())
            return
        log.info("(OMC)一共需要执行{}轮".format(len(self.input_data)))
        for i in self.input_data:
            log.info("(OMC)进行第{}轮仿真".format(run_steps))
            # 修改xml文件
            log.info("(OMC)修改参数：" + str(i))
            if write_xml(r"/home/simtek/code/" + self.absolute_path, i):
                # 解析文件失败
                message = "仿真失败，模型由于未知原因损坏，请重新导出"
                break

            # 运行可执行文件result
            self.state = "running"
            cmd = [r"/home/simtek/code/" + self.absolute_path + "result"]
            process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            self.run_pid = process.pid
            # 获取命令行输出结果
            output, error = process.communicate()
            if error:
                log.info("(OMC)多轮仿真出错" + str(error))
                message = str(error)
                break

            else:
                run_result_str = output.decode('utf-8')
                message = str(run_result_str)
                if "successfully" in run_result_str:
                    # json_data = {"message": self.request.simulateModelName + "仿真到第{}轮".format(run_steps)}
                    # R.lpush(self.request.userName + "_" + "notification", json.dumps(json_data))
                    log.info("(OMC)successfully")
                    # 从mat中读取数据
                    d = DyMat.DyMatFile(r"/home/simtek/code/" + self.absolute_path + "result_res.mat")

                    dictCsv = {"time": list(d.abscissa("2", True))
                               # "time2": list(d.abscissa("1", True))
                               }

                    dictCsv["time"] = dictCsv["time"][:50]

                    for j in self.outputValNames:
                        d_data = list(d.data(j))[:50]
                        if len(d_data) == 2 and d_data[0] == d_data[1]:
                            d_data = [d_data[0] for i in range(50)]
                        dictCsv[j] = d_data
                    df = pd.DataFrame(pd.DataFrame.from_dict(dictCsv, orient='index').values.T,
                                      columns=list(dictCsv.keys()))
                    if len(self.input_data) == 1:
                        pass
                    else:
                        # 多轮仿真每轮一个scv文件

                        csv_file_name = ""
                        for s in i.values():
                            s = round(s, 4)
                            csv_file_name = csv_file_name + "_" + str(s)
                        df.to_csv(mul_output_path + '{}.csv'.format(csv_file_name),
                                  index=False,
                                  encoding='utf-8')
                else:
                    log.info("(OMC)fail")
                    break
            run_steps += 1

        if run_steps == len(self.input_data):  # 如果每轮都成功
            if len(self.input_data) == 1:
                # 更新数据库
                update_app_pages_records(self.request.pageId, simulate_state=4)
            else:
                # 更新数据库
                update_app_pages_records(self.request.pageId,
                                         release_state=4, is_release=True)
                update_app_spaces_records(self.request.pageId)
                page_release_component_freeze(self.request.pageId)

        else:
            if len(self.input_data) == 1:
                # 更新数据库
                update_app_pages_records(self.request.pageId, simulate_state=3)
            else:
                # 更新数据库
                update_app_pages_records(self.request.pageId, release_state=3)
        if len(self.input_data) == 1:  # 仿真任务
            update_app_pages_records(self.request.pageId, simulate_time=time.time())
            update_app_pages_records(self.request.pageId, simulate_message_read=False)
            update_app_pages_records(self.request.pageId, simulate_err=message)
        else:
            update_app_pages_records(self.request.pageId, release_time=time.time())
            update_app_pages_records(self.request.pageId, release_message_read=False)
            update_app_pages_records(self.request.pageId, release_err=message)


        self.state = "stopped"
        delete_item_from_json(self.uuid)
