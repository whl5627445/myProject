# -- coding: utf-8 --


def GetTreeData(data, tree_data_dict, parent_name=None):
    data_list = []
    if parent_name and "der(" not in parent_name:
        for i in parent_name.split("."):
            data = data.get(i, {})
    elif parent_name and "der(" in parent_name:
        data = data.get(parent_name, {})

    for k, v in data.items():
        init_tree_data_dict = ["", "", "", ""]
        data_dict = {
            "haschild": True,
            "Variables": k,
            "unit": tree_data_dict.get(k, init_tree_data_dict)[1],
            "description": tree_data_dict.get(k, init_tree_data_dict)[2],
            "start": tree_data_dict.get(k, init_tree_data_dict)[3],
        }
        if type(v) is not dict:
            data_dict["haschild"] = False
        data_list.append(data_dict)
    return data_list
