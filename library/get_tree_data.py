

def GetTreeData(data, parent_name=None):
    data_list = []
    if parent_name:
        for i in parent_name.split("."):
            data = data.get(i, {})
    for k, v in data.items():
        data_dict = {
            "haschild": True,
            "Variables": "",
        }
        if type(v) is dict:
            data_dict["Variables"] = k
        else:
            data_dict["Variables"] = k
            data_dict["haschild"] = False
        data_list.append(data_dict)
    return data_list
