import re
from config.db_config import Session, YssimSimulateRecords, AppDataSources, AppPages
from config.redis_config import R
import json
import os
import zipfile
import itertools


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


def update_simulate_records(uuid, simulate_status=None, simulate_result_str=None, simulate_start=None,
                            simulate_start_time=None,
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
            if processDetails.simulate_start_time == 0:  # 仿真开始时间只会设置一次
                processDetails.simulate_start_time = simulate_start_time  # 仿真开始时间
        if simulate_end_time:
            processDetails.simulate_end_time = simulate_end_time  # 仿真结束时间
        if simulate_model_result_path:
            processDetails.simulate_model_result_path = simulate_model_result_path  # 仿真结果文件路径
        if another_name:
            processDetails.another_name = another_name  # 结果记录别名
        session.commit()


def update_compile_records(uuid,
                           compile_status=None,
                           compile_start_time=None,
                           compile_stop_time=None,
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
        session.commit()


def update_app_pages_records(pages_id, mul_result_path=None, simulate_state=None, release_state=None):
    with Session() as session:
        app_pages_record = session.query(AppPages).filter(
            AppPages.id == pages_id).first()
        if mul_result_path:
            app_pages_record.mul_result_path = mul_result_path
        if simulate_state:
            app_pages_record.simulate_state = simulate_state
        if release_state:
            app_pages_record.release_state = release_state
        session.commit()


def sendMessage(omc_obj, username):
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
                        archive.write(file_path, arcname=os.path.relpath(file_path, parent_folder))


def convert_dict_to_list(dict_obj):
    # 定义待返回的结果列表
    result = []

    # 定义一个字典的 key 的列表
    keys = list(dict_obj.keys())

    # 获取字典的值的列表
    values = [dict_obj[k].inputObjList for k in keys]

    # 生成所有可能的元素组合，并将它们转换为字典
    for combination in itertools.product(*values):
        d = {keys[i]: combination[i] for i in range(len(keys))}
        result.append(d)
    return result

def convert_list(lst):
    # 使用 itertools.product() 函数生成所有元素组合，并转换为结果列表
    result = list(itertools.product(*lst))
    res = [list(t) for t in result]
    return res

