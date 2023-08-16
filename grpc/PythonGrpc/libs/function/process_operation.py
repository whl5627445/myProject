import psutil
from config.db_config import Session, YssimSimulateRecords,AppPages
import os
import signal
from libs.function.grpc_log import log
import requests
import json
import configparser


# def getSate(string1):
#     if 'started' in string1:
#         return 'started'
#     if 'closed' in string1:
#         return 'closed'
#     if 'unknown' in string1:
#         return 'unknown'
#     if 'initial' in string1:
#         return 'initial'
#     if 'stopped' in string1:
#         return 'stopped'
#     return "unknown"


# def suspend_process(multiprocessing_id, process_list):
#     for i in process_list:
#         if i.uuid == multiprocessing_id:
#             if i.state == "compiling":
#                 proc = psutil.Process(i.omc_obj._omc_process.pid)  # 传入子进程的pid
#                 proc.suspend()  # 暂停子进程
#                 with Session() as session:
#                     processDetails = session.query(YssimSimulateRecords).filter(
#                         YssimSimulateRecords.id == multiprocessing_id).first()
#                     processDetails.simulate_status = "7"  # 暂停
#                     session.commit()
#                 return {"msg": True}
#             if i.state == "running":
#                 proc = psutil.Process(i.run_pid)   # 传入子进程的pid
#                 proc.suspend()  # 暂停子进程
#                 with Session() as session:
#                     processDetails = session.query(YssimSimulateRecords).filter(
#                         YssimSimulateRecords.id == multiprocessing_id).first()
#                     processDetails.simulate_status = "7"  # 暂停
#                     session.commit()
#                 return {"msg": True}
#     return {"msg": False}
#
#
# def resume_process(multiprocessing_id, process_list):
#     for i in process_list:
#         if i.uuid == multiprocessing_id:
#             if i.state == "compiling":
#                 proc = psutil.Process(i.omc_obj._omc_process.pid)  # 传入任意子进程的pid
#                 proc.resume()  # 恢复子进程
#                 with Session() as session:
#                     processDetails = session.query(YssimSimulateRecords).filter(
#                         YssimSimulateRecords.id == multiprocessing_id).first()
#                     processDetails.state = "8"  # 恢复运行
#                     session.commit()
#                 return {"msg": True}
#             if i.state == "running":
#                 proc = psutil.Process(i.run_pid)  # 传入子进程的pid
#                 proc.resume()  # 恢复子进程
#                 with Session() as session:
#                     processDetails = session.query(YssimSimulateRecords).filter(
#                         YssimSimulateRecords.id == multiprocessing_id).first()
#                     processDetails.state = "8"  # 恢复运行
#                     session.commit()
#                 return {"msg": True}
#
#     return {"msg": False}


def kill_process(multiprocessing_id, om_process_list, dm_process_list, om_task_mark_dict, dymola_task_mark_dict):
    
    for i in om_process_list:
        if i.uuid == multiprocessing_id:
            try:
                if hasattr(i, 'omc_obj'):
                    log.info("(OMC)杀死的进程id："+str(i.omc_obj.omc_process.pid))
                    parent_proc = psutil.Process(i.omc_obj.omc_process.pid)
                    for child_proc in parent_proc.children(recursive=True):
                        log.info("(OMC)关闭子进程："+str(child_proc.pid))
                        os.kill(child_proc.pid, 9)
                    os.kill(i.omc_obj.omc_process.pid, 9)
                    log.info("(OMC)关闭omc进程成功")
                # os.killpg(os.getpgid(i.omc_obj.omc_process.pid), signal.SIGUSR1)
                if i.run_pid:
                    os.kill(i.run_pid, 9)
            # i.omc_obj.sendExpression("quit()")
            except OSError as e:
                log.info(f"(OMC)Error: {e}")
            i.state = "stopped"

            del om_task_mark_dict[i.request.userName]
            om_process_list.remove(i)
            del i
            log.info("(OMC)杀死线程，数据库id:"+multiprocessing_id)
            with Session() as session:
                simulate_records = session.query(YssimSimulateRecords).filter(
                    YssimSimulateRecords.id == multiprocessing_id).first()
                if simulate_records:
                    simulate_records.simulate_status = "3"  # 杀死进程
                app_pages_record = session.query(AppPages).filter(
                    AppPages.id == multiprocessing_id).first()
                if app_pages_record:
                    app_pages_record.release_state = 3  # 杀死进程
                session.commit()
            return {"msg": "End Process:{}".format(multiprocessing_id)}

    for i in dm_process_list:
        if i.uuid == multiprocessing_id:
            config = configparser.ConfigParser()
            config.read('./config/grpc_config.ini')
            DymolaSimulationConnect = config['dymola']['DymolaSimulationConnect']

            url = DymolaSimulationConnect + "/dymola/stopDymola"
            data = {
                "taskId": multiprocessing_id
            }
            headers = {
                "Content-Type": "application/json"
            }
            timeout = 10 * 60 # 10分钟

            response = requests.post(url, data=json.dumps(data), headers=headers, timeout=timeout)
            log.info("(Dymola)发送请求体："+str(response))
            if response.status_code == 200:
                result = response.json()
                log.info("(Dymola)请求返回的结果："+str(result))

                i.state = "stopped"
                del dymola_task_mark_dict[i.request.userName]
                dm_process_list.remove(i)
                del i
                log.info("(Dymola)杀死线程，数据库id:" + multiprocessing_id)
                with Session() as session:
                    processDetails = session.query(YssimSimulateRecords).filter(
                        YssimSimulateRecords.id == multiprocessing_id).first()
                    if processDetails:
                        processDetails.simulate_status = "3"  # 杀死进程
                    app_pages_record = session.query(AppPages).filter(
                        AppPages.id == multiprocessing_id).first()
                    if app_pages_record:
                        app_pages_record.simulate_state = 3  # 杀死进程
                    session.commit()
                return {"msg": "End Process:{}".format(multiprocessing_id)}
    return {"msg": "The process is not found or has ended or failed:{}".format(multiprocessing_id)}
