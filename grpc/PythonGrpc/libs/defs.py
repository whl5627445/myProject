import re
from config.db_config import Session, YssimSimulateRecords, YssimModels
import router_pb2
import json


def initTask():
    dataList = []
    with Session() as session:
        # 查询所有为完成的仿真任务,
        task_record_list = []
        record_list = session.query(YssimSimulateRecords).filter(
            YssimSimulateRecords.simulate_status.in_(["2", "1", "6"]),
            YssimSimulateRecords.deleted_at.is_(None)
        ).all()
        for i in record_list:
            print(i.id)
            yssim_model = session.query(YssimModels).filter(
                YssimModels.userspace_id.in_([i.userspace_id, '0']),
                YssimModels.id == i.package_id,
                YssimModels.deleted_at.is_(None)
            ).first()
            if yssim_model:
                print("未完成", yssim_model.id)
                task_record_list.append(i)
            else:
                i.simulate_status = "3"
                session.commit()
    # 返回任务列表
    for i in task_record_list:
        a = router_pb2.PyOmcSimulationRequest(
            uuid=i.id,
            userSpaceId=i.userspace_id,
            userName=i.username,
            simulateModelName=i.simulate_model_name,
            resultFilePath=i.simulate_model_result_path,
            simulationPraData={
                "startTime": i.start_time,
                "stopTime": i.stop_time,
                "method": i.method,
                "numberOfIntervals": i.number_intervals,
                "tolerance": i.tolerance,
            },
            envModelData=i.env_model_data
        )
        dataList.append(a)
    return dataList


def new_another_name(username: str, simulate_model_name: str, userspace_id: str) -> str:
    # 生产新的数据库结果别名
    another_name_list = []
    with Session() as session:
        record_list = session.query(YssimSimulateRecords).filter(
            YssimSimulateRecords.username == username,
            YssimSimulateRecords.simulate_model_name == simulate_model_name,
            YssimSimulateRecords.userspace_id == userspace_id,
            YssimSimulateRecords.simulate_status == "4",
            YssimSimulateRecords.deleted_at.is_(None),
        ).all()

    for record in record_list:
        another_name_list.append(record.another_name)
    max_suffix = 0
    suffix_pattern = re.compile(r"\s(\d+)\s*$")
    for another_name in another_name_list:
        matches = suffix_pattern.findall(another_name)
        if len(matches) > 0:
            suffix = int(matches[0])
            if suffix > max_suffix:
                max_suffix = suffix

    return "结果 " + str(max_suffix + 1)


def update_records(uuid, simulate_status=None, simulate_result_str=None, simulate_start=None, simulate_start_time=None,
                   simulate_end_time=None, simulate_model_result_path=None, another_name=None
                   ):
    with Session() as session:
        processDetails = session.query(YssimSimulateRecords).filter(
            YssimSimulateRecords.id == uuid).first()
        if simulate_status:
            processDetails.simulate_status = simulate_status  # 更改状态
        if simulate_result_str:
            processDetails.simulate_result_str = simulate_result_str  # 更改仿真结果字符串
        if simulate_start:
            processDetails.simulate_start = simulate_start  # 仿真开始标致
        if simulate_start_time:
            if processDetails.simulate_start_time is not None:  # 仿真开始时间只会设置一次
                processDetails.simulate_start_time = simulate_start_time  # 仿真开始时间
        if simulate_end_time:
            processDetails.simulate_end_time = simulate_end_time  # 仿真结束时间
        if simulate_model_result_path:
            processDetails.simulate_model_result_path = simulate_model_result_path  # 仿真结果文件路径
        if another_name:
            processDetails.another_name = another_name  # 结果记录别名
        session.commit()

# def initOmc():
#     print("clearProgram:", omc.sendExpression("clearProgram()"))
#     print("setModelicaPath:", omc.sendExpression("setModelicaPath(\"/usr/lib/omlibrary\")"))
#     print("Buildings9.1.0初始化:", omc.sendExpression("loadModel(Buildings, {\"9.1.0\"},true,\"\",false)"))
#     print("Modelica4.0.0初始化:", omc.sendExpression("loadModel(Modelica, {\"4.0.0\"},true,\"\",false)"))
#     print("SolarPower初始化:", omc.sendExpression("loadModel(SolarPower, {\"\"},true,\"\",false)"))
#     print("WindPowerSystem初始化:", omc.sendExpression("loadModel(WindPowerSystem, {\"\"},true,\"\",false)"))
#     fmuPath = omc.buildModelFmu(className="Modelica.Blocks.Examples.PID_Controller", fileNamePrefix="xxx")
#     print("testBuildFMU:", fmuPath)
#     dirname = os.path.dirname(fmuPath)
#     for filename in os.listdir(dirname):
#         if filename.startswith("xxx"):
#             os.remove(os.path.join(dirname, filename))
#             print("删除文件：", filename)
#
