# -- coding: utf-8 --

from config.omc import omc

def GetModelChild(model_name):
    child_list = omc.getClassNames(model_name)
    data_list = []
    if child_list != ['']:
        for i in child_list:
            haschild = omc.getClassNames(".".join([model_name, i]))
            mn_data = {
                "model_name": i,
                "haschild": True if haschild != [""] else False,
                "image": "",
                }
            data_list.append(mn_data)
    return data_list



def GetModelHasChild(model_name):
    child_list = omc.getClassNames(model_name)
    if child_list == ['']:
        return False
    return True

