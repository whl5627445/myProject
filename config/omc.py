# -- coding: utf-8 --
from library.OMPython.OMCSessionZMQ import OMCSessionZMQ


def omc_init():
    omc_obj = OMCSessionZMQ()

    # omc_obj.sendExpression("setCommandLineOptions(\"-d=nfAPI\")")
    omc_obj.sendExpression("loadModel(Modelica, {\"3.2.3\"},true,\"\",false)")
    omc_obj.sendExpression("loadModel(Complex, {\"3.2.3\"},true,\"\",false)")
    omc_obj.sendExpression("loadModel(ModelicaServices, {\"3.2.3\"},true,\"\",false)")
    omc_obj.sendExpression("loadModel(ModelicaReference, {\"3.2.3\"},true,\"\",false)")
    return omc_obj
omc = omc_init()

