# -- coding: utf-8 --
from config.omc import omc


def SetSimulationOptions(model_name, StartTime, StopTime, Tolerance,Interval):
    pass
    annotate_str = "experiment(StartTime=" + str(StartTime) + ",StopTime=" + str(StopTime) + ",Tolerance=" + str(Tolerance) + ",Interval=" + str(Interval) + ")"
    result = omc.addClassAnnotation(model_name, annotate_str)
    return result
