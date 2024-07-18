from config.db_config import (
    Session,
    YssimSimulateRecords,
    YssimModels,
    AppDataSources,
    ParameterCalibrationRecord,
)
import router_pb2
from libs.function.run_result_json import read_json_file
from libs.function.grpc_log import log
from sqlalchemy import or_


def initOmTask():
    omcDataList = []
    # 查询所有为完成的仿真任务
    with Session() as session:
        task_record_list = []
        record_list = (
            session.query(YssimSimulateRecords)
            .filter(
                YssimSimulateRecords.simulate_status.in_(["2", "1", "6"]),
                YssimSimulateRecords.simulate_type == "OM",
                YssimSimulateRecords.deleted_at.is_(None),
            )
            .all()
        )
        log.info("(OMC)未完成的仿真记录：" + str([k.id for k in record_list]))
        for i in record_list:
            yssim_model = (
                session.query(YssimModels)
                .filter(
                    YssimModels.userspace_id.in_([i.userspace_id, "0"]),
                    YssimModels.id == i.package_id,
                    YssimModels.deleted_at.is_(None),
                )
                .first()
            )
            if yssim_model:
                log.info("(OMC)重新开始仿真的模型：" + yssim_model.package_name)
                i.package_name = yssim_model.package_name
                i.package_file_path = yssim_model.file_path
                task_record_list.append(i)
            else:
                log.info("(OMC)重新开始仿真，模型不存在：" + i.simulate_model_name + "，退出线程。")
                i.simulate_status = "3"
                session.commit()

    # 仿真任务列表
    for i in task_record_list:
        a = router_pb2.SubmitTaskRequest(
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
            taskType="simulate",
        )
        omcDataList.append(a)
    # 查询所有未完成的数据源任务
    with Session() as session:
        task_record_list = []
        record_list = (
            session.query(AppDataSources)
            .filter(
                AppDataSources.compile_status.in_(["2", "1"]),
                AppDataSources.compile_type == "OM",
                AppDataSources.deleted_at.is_(None),
            )
            .all()
        )
        log.info("(OMC)未完成的导出数据源记录：" + str([k.id for k in record_list]))
        for i in record_list:
            yssim_model = (
                session.query(YssimModels)
                .filter(
                    YssimModels.userspace_id.in_([i.user_space_id, "0"]),
                    YssimModels.id == i.package_id,
                    YssimModels.deleted_at.is_(None),
                )
                .first()
            )
            if yssim_model:
                log.info("(OMC)重新开始导出数据源的模型：" + yssim_model.package_name)
                i.package_name = yssim_model.package_name
                i.package_file_path = yssim_model.file_path
                task_record_list.append(i)
            else:
                log.info("(OMC)重新开始导出数据源，模型不存在：" + i.simulate_model_name + "，退出线程。")
                i.compile_status = "3"
                session.commit()
    # 编译任务列表
    for i in task_record_list:
        a = router_pb2.SubmitTaskRequest(
            uuid=i.id,
            userSpaceId=i.user_space_id,
            userName=i.username,
            simulateModelName=i.model_name,
            resultFilePath=i.compile_path,
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
            taskType="translate",
        )

        omcDataList.append(a)

    # # 多轮仿真任务
    # json_data = read_json_file()
    # for item in json_data:
    #     data = router_pb2.SubmitTaskRequest()
    #     data.uuid=item["id"]
    #     data.userSpaceId = item["userSpaceId"],
    #     data.userName = item["userName"],
    #     data.simulateModelName = item["simulateModelName"],
    #     data.resultFilePath = item["resultFilePath"],
    #     data.simulationPraData = item["simulationPraData"],
    #     data.envModelData = item["envModelData"],
    #     data.simulateType = item["simulateType"],
    #
    #     data.packageName = item["packageName"],
    #     data.packageFilePath = item["packageFilePath"],
    #
    #     data.taskType = item["taskType"],
    #
    #     data.pageId = item["pageId"],
    #     dataDict = {}
    #     for key,val in item["inputValData"].items():
    #         dataDict[key] = data.inputObj
    #     data.inputValData = item[""],
    #
    #     data.outputValNames = item["outputValNames"]
    #
    #
    #     omcDataList.append(data)

    return omcDataList


def initDmTask():
    dmDataList = []
    # 查询所有为完成的仿真任务
    with Session() as session:
        task_record_list = []
        record_list = (
            session.query(YssimSimulateRecords)
            .filter(
                YssimSimulateRecords.simulate_status.in_(["2", "1", "6"]),
                YssimSimulateRecords.simulate_type == "DM",
                YssimSimulateRecords.deleted_at.is_(None),
            )
            .all()
        )
        log.info("(Dymola)未完成的仿真记录：" + str([k.id for k in record_list]))
        for i in record_list:
            yssim_model = (
                session.query(YssimModels)
                .filter(
                    YssimModels.userspace_id.in_([i.userspace_id, "0"]),
                    YssimModels.id == i.package_id,
                    YssimModels.deleted_at.is_(None),
                )
                .first()
            )
            if yssim_model:
                log.info("(Dymola)重新开始仿真的模型：" + yssim_model.package_name)
                i.package_name = yssim_model.package_name
                i.package_file_path = yssim_model.file_path
                task_record_list.append(i)
            else:
                log.info("(Dymola)重新开始仿真，模型不存在：" + i.simulate_model_name + "，退出线程。")
                i.simulate_status = "3"
                session.commit()

    # 仿真任务列表
    for i in task_record_list:
        a = router_pb2.SubmitTaskRequest(
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
            taskType="simulate",
        )
        dmDataList.append(a)

        # dm_threading = DmSimulation(a)
        # dm_threading.start()
    # 查询所有为完成的数据源任务
    with Session() as session:
        task_record_list = []
        record_list = (
            session.query(AppDataSources)
            .filter(
                AppDataSources.compile_status.in_(["2", "1"]),
                AppDataSources.compile_type == "DM",
                AppDataSources.deleted_at.is_(None),
            )
            .all()
        )
        log.info("(Dymola)未完成的数据源记录：" + str([k.id for k in record_list]))
        for i in record_list:
            yssim_model = (
                session.query(YssimModels)
                .filter(
                    YssimModels.userspace_id.in_([i.user_space_id, "0"]),
                    YssimModels.id == i.package_id,
                    YssimModels.deleted_at.is_(None),
                )
                .first()
            )
            if yssim_model:
                log.info("(Dymola)重新开始导出数据源的模型：" + yssim_model.package_name)
                i.package_name = yssim_model.package_name
                i.package_file_path = yssim_model.file_path
                task_record_list.append(i)
            else:
                log.info("(Dymola)重新开始导出数据源，模型不存在：" + i.simulate_model_name + "，退出线程。")
                i.compile_status = "3"
                session.commit()
    # 编译任务列表
    for i in task_record_list:
        a = router_pb2.SubmitTaskRequest(
            uuid=i.id,
            userSpaceId=i.user_space_id,
            userName=i.username,
            simulateModelName=i.model_name,
            resultFilePath=i.compile_path,
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
            taskType="translate",
        )
        dmDataList.append(a)
    return dmDataList


def initCalibrationTask():
    compileList = []
    simulateList = []
    # 查询所有为完成的仿真任务
    with Session() as session:
        task_record_list = []
        record_list = (
            session.query(ParameterCalibrationRecord)
            .filter(
                ParameterCalibrationRecord.compile_status.in_(["6"]),
                ParameterCalibrationRecord.deleted_at.is_(None),
            )
            .all()
        )
    # 编译任务列表
    record_list.extend(
        session.query(ParameterCalibrationRecord)
        .filter(
            ParameterCalibrationRecord.simulate_status.in_(["2", "1", "6"]),
            ParameterCalibrationRecord.deleted_at.is_(None),
        )
        .all()
    )
    for i in record_list:
        yssim_model = (
            session.query(YssimModels)
            .filter(
                YssimModels.userspace_id.in_([i.userspace_id, "0"]),
                YssimModels.id == i.package_id,
                YssimModels.deleted_at.is_(None),
            )
            .first()
        )
        if yssim_model:
            log.info("(标定)重新开始仿真的模型：" + yssim_model.package_name)
            i.package_name = yssim_model.package_name
            # i.package_file_path = yssim_model.file_path
            task_record_list.append(i)
        else:
            log.info("(标定)重新开始仿真，模型不存在：" + i.model_name + "，任务状态变更为失败。")
            i.simulate_status = "3"
            session.commit()
    for i in task_record_list:
        if i.compile_status == "6" or i.compile_status == "1":
            compile_task = router_pb2.SubmitTaskRequest(
                uuid=i.id,
                userSpaceId=i.userspace_id,
                userName=i.username,
                packageName=i.package_name,
                taskType="compile",
                simulateModelName=i.model_name,
                resultFilePath=i.compile_path,
                envModelData=i.compile_Dependencies,
                simulateType="OM",
            )
            compileList.append(compile_task)
    # 仿真任务列表
    for i in task_record_list:
        if (
            i.simulate_status == "6"
            or i.simulate_status == "1"
            or i.simulate_status == "2"
        ):
            a = router_pb2.SubmitTaskRequest(
                uuid=i.id,
                simulateType="OM",
                taskType="simulate",
            )

            simulateList.append(a)
    return compileList, simulateList
