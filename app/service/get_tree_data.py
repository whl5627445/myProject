# -- coding: utf-8 --


def GetTreeData(data, tree_data_dict, parent_name=None):
    data_list = []
    if parent_name:
        for i in parent_name.split("."):
            data = data.get(i, {})

    data_id = 0
    for k, v in data.items():
        init_tree_data_dict = ["", "", "", "", ""]
        data_dict = {
            "id":data_id,
            "haschild": True,
            "Variables": k,
            "unit": tree_data_dict.get(k, init_tree_data_dict)[1],
            "description": tree_data_dict.get(k, init_tree_data_dict)[2],
            "start": tree_data_dict.get(k, init_tree_data_dict)[3],
        }
        if type(v) is not dict:
            data_dict["haschild"] = False
        data_list.append(data_dict)
        data_id += 1
    return data_list
