import pandas as pd
import os
import zarr
import psutil

from db_config.config import Session, ProcessState


# def TimeStampToTime(timestamp):
#     timeStruct = time.localtime(timestamp)
#     return time.strftime('%Y-%m-%d %H:%M:%S', timeStruct)


def getSate(string1):
    if 'started' in string1:
        return 'started'
    if 'closed' in string1:
        return 'closed'
    if 'unknown' in string1:
        return 'unknown'
    if 'initial' in string1:
        return 'initial'
    if 'stopped' in string1:
        return 'stopped'
    return "unknown"


def suspendProcess(multiprocessing_id, processList):
    for i in processList:
        if i.uuid == multiprocessing_id:
            proc = psutil.Process(i.pid)  # 传入任意子进程的pid
            proc.suspend()  # 暂停子进程
            with Session() as session:
                processDetails = session.query(ProcessState).filter(ProcessState.uuid == multiprocessing_id).first()
                processDetails.state = "已暂停"
                session.commit()
            return {"msg": True}
    return {"msg": False}


def resumeProcess(multiprocessing_id, processList):
    for i in processList:
        if i.uuid == multiprocessing_id:
            proc = psutil.Process(i.pid)  # 传入任意子进程的pid
            proc.resume()  # 恢复子进程
            with Session() as session:
                processDetails = session.query(ProcessState).filter(ProcessState.uuid == multiprocessing_id).first()
                processDetails.state = "恢复运行"
                session.commit()
            return {"msg": True}
    return {"msg": False}


def killProcess(multiprocessing_id, processList, managerResDict):
    for i in processList:
        if i.uuid == multiprocessing_id:
            processState = getSate(i.__repr__())
            if i.is_alive():
                i.kill()
            else:
                i.close()
            with Session() as session:
                processDetails = session.query(ProcessState).filter(ProcessState.uuid == multiprocessing_id).first()
                processDetails.state = "已退出"
                session.commit()

                if multiprocessing_id in managerResDict:
                    zarr.save(processDetails.resPath, managerResDict[multiprocessing_id])
                    del managerResDict[multiprocessing_id]

            return {"msg": "End {} Process:{}".format(processState, multiprocessing_id)}

    return {"msg": "The process is not found or has ended or failed:{}".format(multiprocessing_id)}
