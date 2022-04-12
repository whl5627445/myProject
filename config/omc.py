# -- coding: utf-8 --
from library.OMPython.OMCSessionZMQ import OMCSessionZMQ
# from config.DB_config import DBSession
# from app.model.User.User import User


def omc_init():
    omc_obj = OMCSessionZMQ()
    omc_obj.sendExpression("loadModel(Modelica, {\"3.2.3\"},true,\"\",false)")
    return omc_obj
omc = omc_init()
# omc_dict = {}
#
# session = DBSession()
# OmcPortBase = 50000
# user_all = session.query(User).all()
# for i in user_all:
#     user_name = i.uuser_name
#     port = OmcPortBase + i.port
#     docker_ID = os.popen(f"sudo docker ps -aq --filter name=\"^{user_name}$\"").read().replace("\n", "")
#     if docker_ID:
#         omc = OMCSessionZMQ(dockerContainer=docker_ID,dockerNetwork="host", port=port)
#     else:
#         omc = OMCSessionZMQ(docker="openmodelica/openmodelica:v1.18.0-minimal",dockerNetwork="host", port=port)
#
#     omc_dict[user_name] = OMCSessionZMQ(port=OmcPortBase + i.port)
