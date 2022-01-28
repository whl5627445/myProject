# -- coding: utf-8 --


def GetTreeData(data, variable_name):
    data_list = []
    data_id = 0
    data_dict_all = {}
    for i in data:
        var_parent_list = i[0].split('.')
        if i[1].startswith("_") or i[1].startswith('der('):
            continue
        if not variable_name:
            var_prefix = ""
            var = i[0]
            variable_name_len = len(i[1].split('.'))
        else:
            var_prefix = variable_name
            var = i[1]
            variable_name_len = len(variable_name.split('.'))
        if variable_name_len == len(var_parent_list):
            data_dict_all[var] = {
                "id": data_id,
                "haschild": False,
                "Variables": i[1] if i[1].startswith('der(') else i[1].split('.')[-1],
                "unit": i[2],
                "description": i[3],
                "start": i[4],
                }
        else:
            Variables = i[0] if i[0].startswith('der(') else i[0].removeprefix(var_prefix + ".").split('.')[0]
            if data_dict_all.get(Variables, None):
                continue
            data_dict_all[Variables] = {
                "id": data_id,
                "haschild": True,
                "Variables": Variables,
                "unit": "",
                "description": "",
                "start": "",
                }
        data_id += 1
    for i in data_dict_all.values():
        data_list.append(i)
    return data_list
