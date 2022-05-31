# -- coding: utf-8 --
from config.omc import omc


def ClearProgram():
    omc.sendExpression("clearProgram()")
    omc.sendExpression("loadModel(Modelica, {\"3.2.3\"},true,\"\",false)")
    omc.sendExpression("loadModel(Complex, {\"3.2.3\"},true,\"\",false)")
    omc.sendExpression("loadModel(ModelicaServices, {\"3.2.3\"},true,\"\",false)")
    omc.sendExpression("loadModel(ModelicaReference, {\"3.2.3\"},true,\"\",false)")
    omc.sendExpression("loadModel(Buildings, {\"8.0.0\"},true,\"\",false)")
    omc.sendExpression("loadModel(SolarPower, {\"\"},true,\"\",false)")
    omc.sendExpression("loadModel(WindPowerSystem, {\"\"},true,\"\",false)")
    return
