# -- coding: utf-8 --


def GetTreeData(data, tree_data_dict, parent_name=""):
    data_list = []
    name_list = []
    if parent_name:
        name_list.append(parent_name)
        for i in parent_name.split("."):
            data = data.get(i, {})
    data_id = 0

    for k, v in data.items():
        name_list.append(k)
        var_name = tree_data_dict.get(".".join(name_list), None)
        data_dict = {
            "id":data_id,
            "haschild": True,
            "Variables": k,
            "unit": var_name.unit if var_name else "",
            "description":var_name.description  if var_name else "",
            "start": var_name.start  if var_name else "",
        }
        if type(v) is not dict:
            data_dict["haschild"] = False
        data_list.append(data_dict)
        data_id += 1
    return data_list
