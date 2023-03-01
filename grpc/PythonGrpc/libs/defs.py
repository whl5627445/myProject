import shutil
import time
import zarr
import psutil
import zipfile
import os

from config.omc import omc
from config.db_config import Session, YssimSimulateRecords


def buildFMU(moPath, className, userName, resPath):
    # if omc.loadFile("/yssim-go/public/UserFiles/UploadFile/xuqingda/ChillerStage/20230216111124/ChillerStage.mo"):
    #     return False
    adsPath = "/home/simtek/code/"
    fileNamePrefix = userName + time.strftime('%H%M%S', time.localtime(time.time()))
    print(moPath, className, userName, resPath)
    if moPath == "":  # 如果是系统模型,直接导出fmu
        fmuPath = omc.buildModelFmu(className=className, fileNamePrefix=fileNamePrefix)
        if fmuPath == "":
            print("moPath为空被判定为系统模型，导出fmu" + className + "失败！")
            return False
    else:  # 如果不是系统模型，先判断mo文件是否存在，再加载mo文件，导出fmu，卸载模型
        if os.path.exists(adsPath + moPath):
            print("mo文件存在:", adsPath + moPath)
        else:
            print("mo文件不存在:", adsPath + moPath)
            return False
        if not omc.loadFile(adsPath + moPath):
            print("加载" + moPath + "失败！")
            return False
        fmuPath = omc.buildModelFmu(className=className, fileNamePrefix=fileNamePrefix)
        print("fmupath:", fmuPath)
        if fmuPath == "":
            print("moPath不为空被判定为用户模型，导出fmu" + className + "失败！")
            return False
        if not omc.deleteClass(className):
            print("卸载" + className + "失败！")
            return False
    try:
        newFmuPath = adsPath + resPath + className.replace(".", "_") + ".fmu"
        shutil.move(fmuPath, newFmuPath)
        # copyPath = shutil.copy(fmuPath, newFmuPath)
        # movePath = shutil.move(fmuPath, newFmuPath + ".zip")
        # zip_file = zipfile.ZipFile(movePath)
        # zip_file.extractall(adsPath + resPath)
        # zip_file.close()
        # os.rename(adsPath + resPath + "modelDescription.xml", adsPath + resPath + "result_init.xml")
        dirname = os.path.dirname(fmuPath)
        for filename in os.listdir(dirname):
            if filename.startswith(fileNamePrefix):
                shutil.move(os.path.join(dirname, filename), adsPath + resPath + filename)
                print("移动文件：", filename)
    except Exception as e:
        return False
    return newFmuPath


def initOmc():
    print("clearProgram:", omc.sendExpression("clearProgram()"))
    print("setModelicaPath:", omc.sendExpression("setModelicaPath(\"/usr/lib/omlibrary\")"))
    print("Buildings9.1.0初始化:", omc.sendExpression("loadModel(Buildings, {\"9.1.0\"},true,\"\",false)"))
    print("Modelica4.0.0初始化:", omc.sendExpression("loadModel(Modelica, {\"4.0.0\"},true,\"\",false)"))
    print("SolarPower初始化:", omc.sendExpression("loadModel(SolarPower, {\"\"},true,\"\",false)"))
    print("WindPowerSystem初始化:", omc.sendExpression("loadModel(WindPowerSystem, {\"\"},true,\"\",false)"))
    fmuPath = omc.buildModelFmu(className="Modelica.Blocks.Examples.PID_Controller", fileNamePrefix="xxx")
    print("testBuildFMU:", fmuPath)
    dirname = os.path.dirname(fmuPath)
    for filename in os.listdir(dirname):
        if filename.startswith("xxx"):
            os.remove(os.path.join(dirname, filename))
            print("删除文件：", filename)


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
