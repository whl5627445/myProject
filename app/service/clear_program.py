# -- coding: utf-8 --
from config.omc import omc


def ClearProgram():
    omc.sendExpression("clearProgram()")
    return
