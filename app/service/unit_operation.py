# -- coding: utf-8 --

from config.omc import omc


def ConvertUnits(s1, s2):
    result = omc.convertUnits(s1, s2)
    return result
