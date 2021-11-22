# -- coding: utf-8 --
from library.OMPython.OMCSessionZMQ import OMCSessionZMQ

omc = OMCSessionZMQ()
omc.sendExpression("loadModel(Modelica, {\"3.2.3\"},true,\"\",false)")

def OmcFactory():
    omc_once = OMCSessionZMQ(once=True)
    return omc_once
