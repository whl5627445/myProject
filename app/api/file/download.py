# -- coding: utf-8 --
from router.download_router import router
from fastapi import HTTPException
from app.BaseModel.respose_model import ResponseModel, InitResponseModel
from app.model.ModelsPackage.ModelsInformation import ModelsInformation
from fastapi import File, UploadFile, Request
from config.DB_config import DBSession
from app.model.Simulate.SimulateResult import SimulateResult
from fastapi.responses import FileResponse
from app.BaseModel.simulate import SimulateResultExportModel
from library.file_operation import FileOperation
import pandas as pd
import random


session = DBSession()


@router.get("/getfilelist", response_model=ResponseModel)
async def GetFileListView(request: Request):
    """
    # 用户获取mo文件信息接口， 可以根据url进行下载
    ## return: 包名， 上传时间， 下载路径
    """
    res = InitResponseModel()
    username = request.user.username
    package_list = session.query(ModelsInformation).filter_by(sys_or_user=username).all()
    for i in range(len(package_list)):
        data_dict = {
            "id": i,
            "package_name": package_list[i].package_name,
            "create_time": package_list[i].create_time,
            # "url": "static/" + package_list[i].file_path,
            }
        res.data.append(data_dict)
    return res


@router.post("/getsimulateresultfile")
async def GetSimulateResultFileView(request: Request, items: SimulateResultExportModel):
    """
    # 用户获取仿真结果文件接口，返回对应仿真记录的仿真结果
    ## record_id：仿真记录接口获取的id
    ## export_type：导出数据文件的类型， 目前支持txt，csv，xlsx(txt暂时不用)
    ## var_list：需要导出的变量数组， 变量名称需是全名
    ## return: 返回对应记录的文件地址
    """
    res = InitResponseModel()
    username = request.user.username
    export_type = items.export_type
    data = session.query(SimulateResult.model_variable_name, SimulateResult.model_variable_data_abscissa,
                         SimulateResult.model_variable_data).filter_by(simulate_record_id=items.record_id, username=username).filter(SimulateResult.model_variable_name.in_(items.var_list)).all()
    name_dict = {}
    for i in data:
        name_dict[i[0]] = {
            "model_variable_name": i[0],
            "model_variable_data_abscissa": i[1],
            "model_variable_data": i[2],
            }
    pd_data = pd.DataFrame(columns=["time"])
    try:
        for k, v in name_dict.items():
            columns = ["time", k]
            var_data = pd.DataFrame(columns=columns)
            var_data["time"] = v["model_variable_data_abscissa"]
            var_data[k] = v["model_variable_data"]
            if pd_data is None:
                pd_data = var_data
            else:
                pd_data = pd.merge(pd_data, var_data, on="time", how="outer")
            pd_data = pd_data.drop_duplicates()
    except Exception as e:
        print(e)
        return {"msg": "参数有误","status": 2}
    file_path = username +"/static/" + username + "/"
    file_name = "".join(random.sample('zyxwvutsrqponmlkjihgfedcba0123456789',20)) + "." + export_type
    data_file = file_path + file_name
    FileOperation.touth_file(file_path, file_name)
    if export_type == "csv":
        pd_data.to_csv(data_file, index=False)
    elif export_type == "xlsx":
        pd_data.to_excel(data_file, index=False)
    else:
        raise HTTPException(status_code=400, detail="not found")
    return FileResponse(data_file, filename=file_name)