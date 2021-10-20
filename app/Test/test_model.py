from fastapi.testclient import TestClient
from app.api.modelview import *
from main import app


client = TestClient(app)

def test_GetRootModel():
    response = client.get("/simulatemodel/test", params={"modelname": "123"})
    assert response.status_code == 200
    assert response.json() == {"msg": "Hello World"}

def test_listrootlibrary():
    response = client.get("/simulatemodel/listrootlibrary")
    assert response.status_code == 200
    print ("左侧模型根目录数据： ", response.json())


def test_listlibrary():
    response = client.get("/simulatemodel/listlibrary", params={"modelname": "Modelica", "sys_user": "sys"})
    assert response.status_code == 200
    print ("模型子目录数据： ", response.json())


def test_getgraphicsdata():
    response = client.get("/simulatemodel/getgraphicsdata", params={"modelname": "Modelica.Blocks.Examples.PID_Controller", "sys_user": "user"})
    assert response.status_code == 200
    print ("画图数据： ", response.json())
