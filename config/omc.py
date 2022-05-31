# -- coding: utf-8 --
from library.OMPython.OMCSessionZMQ import OMCSessionZMQ


def omc_init():
    omc_obj = OMCSessionZMQ()
    # omc_obj.sendExpression("setCommandLineOptions(\"-d=nfAPI\")")
    return omc_obj
omc = omc_init()

