from config.db_config import Session, YssimSimulateRecords, YssimModels
import router_pb2
from libs.DmSimulation import DmSimulation


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
                i.package_name = yssim_model.package_name
                i.package_file_path = yssim_model.file_path
                task_record_list.append(i)
            else:
                i.simulate_status = "3"
                session.commit()
    # 返回任务列表
    for i in task_record_list:
        a = router_pb2.SimulationRequest(
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
            envModelData=i.env_model_data,
            simulateType="OM",
            packageName=i.package_name,
            packageFilePath=i.package_file_path,
        )
        if i.simulate_type == "OM":
            dataList.append(a)
        if i.simulate_type == "DM":

            dm_threading = DmSimulation(a)
            dm_threading.start()

    return dataList
