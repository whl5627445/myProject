import re
import re
import time
import random
import string

from config.db_config import Session, YssimSimulateRecords, AppDataSources, AppPages, AppSpaces, AppPagesComponent, \
    AppPagesComponentRelease
from config.redis_config import R
import json
import os
import zipfile
import itertools
from libs.function.grpc_log import log


def new_another_name(username: str, simulate_model_name: str, package_id: str, userspace_id: str) -> str:
    # 生产新的数据库结果别名
    with Session() as session:
        record_list = session.query(YssimSimulateRecords).filter(
            YssimSimulateRecords.package_id == package_id,
            YssimSimulateRecords.username == username,
            YssimSimulateRecords.userspace_id == userspace_id,
            YssimSimulateRecords.simulate_model_name == simulate_model_name,
            YssimSimulateRecords.simulate_status == "4",
            YssimSimulateRecords.deleted_at.is_(None),
        ).all()
    max_suffix = 0
    for another_name in [i.another_name for i in record_list]:
        # 使用正则表达式提取数字部分
        match = re.search(r"\s*(\d+)\s*$", another_name)
        # 如果能找到数字
        if match:
            # 将匹配到的数字转换为整数
            num = int(match.group())
            # 如果当前数字比最大值大或者最大值为空，则更新最大值
            if max_suffix is None or num > max_suffix:
                max_suffix = num

    return "结果 " + str(max_suffix + 1)


def update_app_spaces_records(page_id):
    # 发布完成更改app_space的发布状态is_release为True
    with Session() as session:
        query = session.query(AppPages, AppSpaces).join(
            AppSpaces, AppPages.app_space_id == AppSpaces.id
        ).filter(
            AppPages.id == page_id
        ).first()
        app_space = query[1]
        app_space.is_release = True
        session.commit()


def update_simulate_records(uuid, simulate_status=None, simulate_result_str=None, simulate_start=None,
                            simulate_start_time=None,
                            simulate_end_time=None, simulate_model_result_path=None, another_name=None,
                            result_run_time=None
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
            if processDetails.simulate_start_time == 0:  # 仿真开始时间只会设置一次
                processDetails.simulate_start_time = simulate_start_time  # 仿真开始时间
        if simulate_end_time:
            processDetails.simulate_end_time = simulate_end_time  # 仿真结束时间
        if simulate_model_result_path:
            processDetails.simulate_model_result_path = simulate_model_result_path  # 仿真结果文件路径
        if another_name:
            processDetails.another_name = another_name  # 结果记录别名
        if result_run_time:
            processDetails.result_run_time = result_run_time  # 可执行文件的运行时间

        session.commit()


def update_compile_records(uuid,
                           compile_status=None,
                           compile_start_time=None,
                           compile_stop_time=None,
                           result_run_time=None
                           ):
    with Session() as session:
        data_sources_record = session.query(AppDataSources).filter(
            AppDataSources.id == uuid).first()
        if compile_status:
            data_sources_record.compile_status = compile_status  # 更改状态
        if compile_start_time:
            # if data_sources_record.compile_start_time == 0 or data_sources_record.compile_start_time is None :  # 仿真开始时间只会设置一次
            data_sources_record.compile_start_time = compile_start_time  # 仿真开始时间
        if compile_stop_time:
            data_sources_record.compile_stop_time = compile_stop_time  # 仿真结束时间
        if result_run_time:
            data_sources_record.result_run_time = result_run_time
        session.commit()


def update_app_pages_records(pages_id, mul_result_path=None, simulate_state=None, release_state=None, release_time=None,
                             simulate_time=None, release_message_read=None, simulate_message_read=None,
                             simulate_err=None, release_err=None, is_release=None,naming_order=None):
    with Session() as session:
        app_pages_record = session.query(AppPages).filter(
            AppPages.id == pages_id).first()
        if is_release is not None:
            app_pages_record.is_release = is_release
        if mul_result_path:
            app_pages_record.mul_result_path = mul_result_path
        if simulate_state:
            app_pages_record.simulate_state = simulate_state
        if release_state:
            app_pages_record.release_state = release_state
        if release_time:
            app_pages_record.release_time = release_time
        if simulate_time:
            app_pages_record.simulate_time = simulate_time
        if release_message_read is not None:
            app_pages_record.release_message_read = release_message_read
        if simulate_message_read is not None:
            app_pages_record.simulate_message_read = simulate_message_read
        if simulate_err:
            app_pages_record.simulate_err = simulate_err
        if release_err:
            app_pages_record.release_err = release_err
        if naming_order:
            app_pages_record.naming_order = naming_order
        session.commit()


def page_release_component_freeze(pages_id):
    with Session() as session:
        components_list = session.query(AppPagesComponent).filter(
            AppPagesComponent.page_id == pages_id, AppPagesComponent.deleted_at == None).all()
        new_components_list = []
        session.query(AppPagesComponentRelease).filter(
            AppPagesComponentRelease.page_id == pages_id).delete()
        for i in components_list:
            component_release = AppPagesComponentRelease(
                id=i.id,
                page_id=i.page_id,
                type=i.type,
                width=i.width,
                height=i.height,
                position_x=i.position_x,
                position_y=i.position_y,
                angle=i.angle,
                horizontal_flip=i.horizontal_flip,
                vertical_flip=i.vertical_flip,
                opacity=i.opacity,
                other_configuration=i.other_configuration,
                z_index=i.z_index,
                styles=i.styles,
                events=i.events,
                chart_config=i.chart_config,
                option=i.option,
                component_path=i.component_path,
                hide=i.hide,
                lock=i.lock,
                is_group=i.is_group,
                create_time=i.create_time,
                deleted_at=i.deleted_at,
                input_name=i.input_name,
                output=i.output,
                max=i.max,
                min=i.min,
                interval=i.interval
            )

            new_components_list.append(component_release)
        session.add_all(new_components_list)
        session.commit()


def sendMessage(omc_obj, username):
    if omc_obj.omc_process.poll() is not None:
        return
    message_str = omc_obj.getMessagesStringInternal()
    data_list = message_str.split(";,")
    message_list = []
    for i in data_list:
        dl = i.strip().split(",\n")
        message_dict = {}
        for p in dl:
            pl = p.strip()
            if "MODELICAPATH" in pl or "installPackage" in pl or "Downloaded" in pl or "Automatically loaded package" in pl:
                continue
            # elif "Automatically " in pl or "Lexer " in pl:
            #     continue
            elif pl.startswith("message"):
                mes = pl.replace("message = ", "", -1)
                message_dict["message"] = mes[1:-1]
                # print("mes", mes)
            elif pl.startswith("level"):
                level = pl.split(".")
                message_dict["type"] = level[-1]
                # print("level", level)
        if len(message_dict) > 1:
            message_list.append(message_dict)
    for i in message_list:
        message_notice(username, i)
    return message_list


def message_notice(username, mes):
    R.lpush(username + "_" + "notification", json.dumps(mes))


def zip_folders(folders, output_path):
    if len(folders) > 0:
        os.makedirs(os.path.dirname(output_path), exist_ok=True)
        with zipfile.ZipFile(output_path, mode='w') as archive:
            for folder in folders:
                parent_folder = os.path.dirname(folder)
                for root, dirs, files in os.walk(folder):
                    for file in files:
                        file_path = os.path.join(root, file)
                        arcname = os.path.relpath(file_path, parent_folder)
                        archive.write(file_path, arcname=arcname)


def omc_convert_dict_to_list(dict_obj, page_id):
    # 定义待返回的结果列表
    result = []

    # 定义一个字典的 key 的列表
    keys = list(dict_obj.keys())
    log.info("(OMC)需要修改的参数:" + str(keys))
    # with Session() as session:
    #     app_pages_record = session.query(AppPages).filter(
    #         AppPages.id == page_id).first()
    #     app_pages_record.naming_order = keys
    #     session.commit()

    # 获取字典的值的列表
    values = [dict_obj[k].inputObjList for k in keys]
    log.info("(OMC)需要修改的参数:" + str(values))

    # 生成所有可能的元素组合，并将它们转换为字典
    for combination in itertools.product(*values):
        d = {keys[i]: combination[i] for i in range(len(keys))}
        result.append(d)

    return result


def dymola_convert_list(dict_obj, page_id):
    # 定义一个字典的 key 的列表
    keys = list(dict_obj.keys())
    log.info("(Dymola)需要修改的参数:" + str(keys))
    # with Session() as session:
    #     app_pages_record = session.query(AppPages).filter(
    #         AppPages.id == page_id).first()
    #     app_pages_record.naming_order = keys
    #     session.commit()
    # 获取字典的值的列表
    values = [dict_obj[k].inputObjList for k in keys]
    # 使用 itertools.product() 函数生成所有元素组合，并转换为结果列表
    result = list(itertools.product(*values))
    res = [list(t) for t in result]
    return res


def dymola_res_list_to_csv_dict(input_data, input_names):
    result = []
    for i in range(0, len(input_data), len(input_names)):
        row = {}
        for j in range(len(input_names)):
            row[input_names[j]] = input_data[i + j]
        result.append(row)
    return result


def result_step(arr):
    if len(arr) <= 50:
        return arr
    else:
        step = len(arr) // 50  # 计算步长
        new_arr = [arr[i] for i in range(0, len(arr), step)]   # 等间距取值
        return new_arr
