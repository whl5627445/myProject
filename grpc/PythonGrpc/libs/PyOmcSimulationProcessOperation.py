import psutil
from config.db_config import Session, YssimSimulateRecords
import os
import signal


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


def suspend_process(multiprocessing_id, process_list):
    for i in process_list:
        if i.uuid == multiprocessing_id:
            if i.state == "compiling":
                proc = psutil.Process(i.omc_obj._omc_process.pid)  # 传入子进程的pid
                proc.suspend()  # 暂停子进程
                with Session() as session:
                    processDetails = session.query(YssimSimulateRecords).filter(
                        YssimSimulateRecords.id == multiprocessing_id).first()
                    processDetails.simulate_status = "7"  # 暂停
                    session.commit()
                return {"msg": True}
            if i.state == "running":
                proc = psutil.Process(i.run_pid)   # 传入子进程的pid
                proc.suspend()  # 暂停子进程
                with Session() as session:
                    processDetails = session.query(YssimSimulateRecords).filter(
                        YssimSimulateRecords.id == multiprocessing_id).first()
                    processDetails.simulate_status = "7"  # 暂停
                    session.commit()
                return {"msg": True}
    return {"msg": False}


def resume_process(multiprocessing_id, process_list):
    for i in process_list:
        if i.uuid == multiprocessing_id:
            if i.state == "compiling":
                proc = psutil.Process(i.omc_obj._omc_process.pid)  # 传入任意子进程的pid
                proc.resume()  # 恢复子进程
                with Session() as session:
                    processDetails = session.query(YssimSimulateRecords).filter(
                        YssimSimulateRecords.id == multiprocessing_id).first()
                    processDetails.state = "8"  # 恢复运行
                    session.commit()
                return {"msg": True}
            if i.state == "running":
                proc = psutil.Process(i.run_pid)  # 传入子进程的pid
                proc.resume()  # 恢复子进程
                with Session() as session:
                    processDetails = session.query(YssimSimulateRecords).filter(
                        YssimSimulateRecords.id == multiprocessing_id).first()
                    processDetails.state = "8"  # 恢复运行
                    session.commit()
                return {"msg": True}

    return {"msg": False}


def kill_py_omc_process(multiprocessing_id, process_list):
    for i in process_list:
        if i.uuid == multiprocessing_id:
            i.state = "stopped"
            try:
                os.kill(i.omc_obj.omc_process.pid, 9)
                # os.killpg(os.getpgid(i.omc_obj.omc_process.pid), signal.SIGUSR1)

            # i.omc_obj.sendExpression("quit()")
            except OSError as e:
                print(f"Error: {e}")
            process_list.remove(i)
            del i
            print("杀死进程")
            with Session() as session:
                processDetails = session.query(YssimSimulateRecords).filter(
                    YssimSimulateRecords.id == multiprocessing_id).first()
                processDetails.simulate_status = "3"  # 杀死进程
                session.commit()
            return {"msg": "End Process:{}".format(multiprocessing_id)}
    return {"msg": "The process is not found or has ended or failed:{}".format(multiprocessing_id)}
