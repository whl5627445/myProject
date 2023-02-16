import shutil
import time
import zarr
import psutil
import zipfile
import os

from db_config.omc import omc
from db_config.config import Session, YssimSimulateRecords


def buildFMU(moPath, className, userName, resPath):
    # if omc.loadFile("/yssim-go/public/UserFiles/UploadFile/xuqingda/ChillerStage/20230216111124/ChillerStage.mo"):
    #     return False
    fileNamePrefix =userName + time.strftime('%H%M%S', time.localtime(time.time()))
    if omc.loadFile("/yssim-go/" + moPath):
        print("加载" + moPath + "失败！")
        return False
    fmuPath = omc.buildModelFmu(className=className, fileNamePrefix=fileNamePrefix)
    if fmuPath != "":
        print("导出fmu" + className + "失败！")
        return False
    if omc.deleteClass(className):
        print("卸载" + className + "失败！")
        return False
    try:
        newFmuPath = "/yssim-go/" + resPath + className.replace(".", "_") + ".fmu"
        copyPath = shutil.copy(fmuPath, newFmuPath)
        print("copyPath:", copyPath)
        movePath = shutil.move(fmuPath, newFmuPath + ".zip")
        print("movePath:", movePath)
        zip_file = zipfile.ZipFile(movePath)
        zip_file.extractall("/yssim-go/" + resPath)
        zip_file.close()
        os.rename("/yssim-go/" + resPath + "modelDescription.xml", "/yssim-go/" + resPath + "result_init.xml")
        dirname = os.path.dirname(fmuPath)
        for filename in os.listdir(dirname):
            if filename.startswith(fileNamePrefix):
                os.remove(os.path.join(dirname, filename))
    except:
        return False
    return newFmuPath


def initOmc():
    print("clearProgram:", omc.sendExpression("clearProgram()"))
    print("setModelicaPath:", omc.sendExpression("setModelicaPath(\"/usr/lib/omlibrary\")"))
    print("Buildings9.1.0初始化:", omc.sendExpression("loadModel(Buildings, {\"9.1.0\"},true,\"\",false)"))
    print("Modelica4.0.0初始化:", omc.sendExpression("loadModel(Modelica, {\"4.0.0\"},true,\"\",false)"))
    print("SolarPower初始化:", omc.sendExpression("loadModel(SolarPower, {\"\"},true,\"\",false)"))
    print("WindPowerSystem初始化:", omc.sendExpression("loadModel(WindPowerSystem, {\"\"},true,\"\",false)"))


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
                processDetails = session.query(YssimSimulateRecords).filter(
                    YssimSimulateRecords.id == multiprocessing_id).first()
                processDetails.simulate_status = "2"  # 暂停
                session.commit()
            return {"msg": True}
    return {"msg": False}


def resumeProcess(multiprocessing_id, processList):
    for i in processList:
        if i.uuid == multiprocessing_id:
            proc = psutil.Process(i.pid)  # 传入任意子进程的pid
            proc.resume()  # 恢复子进程
            with Session() as session:
                processDetails = session.query(YssimSimulateRecords).filter(
                    YssimSimulateRecords.id == multiprocessing_id).first()
                processDetails.state = "2"  # 恢复运行
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
                processDetails = session.query(YssimSimulateRecords).filter(
                    YssimSimulateRecords.id == multiprocessing_id).first()
                processDetails.state = "4"  # 退出
                session.commit()

                if multiprocessing_id in managerResDict:
                    zarr.save(processDetails.resPath, managerResDict[multiprocessing_id])
                    del managerResDict[multiprocessing_id]

            return {"msg": "End {} Process:{}".format(processState, multiprocessing_id)}

    return {"msg": "The process is not found or has ended or failed:{}".format(multiprocessing_id)}
