from fastapi.testclient import TestClient
from main import app
from ..result_data.model_result import *


client = TestClient(app)

# def test_GetRootModel():
#     response = client.get("/simulatemodel/test", params={"modelname": "123"})
#     assert response.status_code == 200
#     # assert response.json() == {"msg": "Hello World"}

def test_listrootlibrary():
    response = client.get("/simulatemodel/listrootlibrary")
    assert response.status_code == 200
    # print ("左侧模型根目录数据： ", response.json())


def test_listlibrary():
    response = client.get("/simulatemodel/listlibrary", params={"model_name": "Modelica", "sys_user": "sys"})
    assert response.status_code == 200
    # print ("模型子目录数据： ", response.json())


def test_getgraphicsdata():
    response = client.get("/simulatemodel/getgraphicsdata", params={"model_name": "Modelica.Blocks.Examples.Filter", "sys_user": "sys"})
    assert response.status_code == 200
    # print(response.json()["data"])
    assert response.json()["data"] == test_getgraphicsdata_res
    # print ("画图数据： ", response.json()["data"])


def test_getmodelcode():
    response = client.get("/simulatemodel/getmodelcode", params={"model_name": "Modelica.Blocks.Examples.Filter", "sys_user": "sys"})
    assert response.status_code == 200
    assert response.json()["data"] == test_getmodelcode_res
    # print("源码数据： ", response.json()["data"])

def test_getmodelparameters():
    response = client.get("/simulatemodel/getmodelparameters", params={"model_name": "Modelica.Blocks.Examples.Filter", "sys_user": "sys", "components_name": "Modelica.Blocks.Sources.Step", "name": "step"})
    assert response.status_code == 200
    assert response.json()["data"] == test_getmodelparameters_res
    # print("模型组件的参数数据 ", response.json()["data"])


def test_setmodelparameters():
    response = client.get("/simulatemodel/setmodelparameters", params={"model_name": "Modelica.Blocks.Examples.Filter", "sys_user": "sys", "components_name": "Modelica.Blocks.Sources.Step", "name": "step"})
    assert response.status_code == 200
    assert response.json()["data"] == test_getmodelparameters_res
    # print("模型组件的参数数据 ", response.json()["data"])
