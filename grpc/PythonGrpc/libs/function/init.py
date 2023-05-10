from config.db_config import Session, YssimSimulateRecords, YssimModels, AppDataSources
import router_pb2
from libs.simulation.DmSimulationThread import DmSimulation


def initOmTask():
    omcDataList = []
    # 查询所有为完成的仿真任务
    with Session() as session:
        task_record_list = []
        record_list = session.query(YssimSimulateRecords).filter(
            YssimSimulateRecords.simulate_status.in_(["2", "1", "6"]),
            YssimSimulateRecords.simulate_type == "OM",
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

    # 仿真任务列表
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
            taskType="simulate"
        )
        omcDataList.append(a)
    # 查询所有为完成的编译任务
    with Session() as session:
        task_record_list = []
        record_list = session.query(AppDataSources).filter(
            AppDataSources.compiler_status.in_(["2", "1", "6"]),
            AppDataSources.compiler_type == "OM",
            AppDataSources.deleted_at.is_(None)
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
    # 编译任务列表
    for i in task_record_list:
        a = router_pb2.SimulationRequest(
            uuid=i.id,
            userSpaceId=i.user_space_id,
            userName=i.username,
            simulateModelName=i.model_name,
            resultFilePath=i.compiler_path,
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
            taskType="translate"
        )

        omcDataList.append(a)

    return omcDataList


def initDmTask():
    dmDataList = []
    # 查询所有为完成的仿真任务
    with Session() as session:
        task_record_list = []
        record_list = session.query(YssimSimulateRecords).filter(
            YssimSimulateRecords.simulate_status.in_(["2", "1", "6"]),
            YssimSimulateRecords.simulate_type == "DM",
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

    # 仿真任务列表
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
            simulateType="DM",
            packageName=i.package_name,
            packageFilePath=i.package_file_path,
            taskType="simulate"
        )
        dmDataList.append(a)

            # dm_threading = DmSimulation(a)
            # dm_threading.start()
    # 查询所有为完成的编译任务
    with Session() as session:
        task_record_list = []
        record_list = session.query(AppDataSources).filter(
            AppDataSources.compiler_status.in_(["2", "1", "6"]),
            AppDataSources.compiler_type == "DM",
            AppDataSources.deleted_at.is_(None)
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
    # 编译任务列表
    for i in task_record_list:
        a = router_pb2.SimulationRequest(
            uuid=i.id,
            userSpaceId=i.user_space_id,
            userName=i.username,
            simulateModelName=i.model_name,
            resultFilePath=i.compiler_path,
            simulationPraData={
                "startTime": i.start_time,
                "stopTime": i.stop_time,
                "method": i.method,
                "numberOfIntervals": i.number_intervals,
                "tolerance": i.tolerance,
            },
            envModelData=i.env_model_data,
            simulateType="DM",
            packageName=i.package_name,
            packageFilePath=i.package_file_path,
            taskType="translate"
        )
        dmDataList.append(a)
    return dmDataList
