import xml.etree.ElementTree as ET


def write_xml(xml_path, var_dict):
    # 解析XML文件为一个树
    result_xml_path = xml_path+"result_init.xml"
    tree = ET.parse(result_xml_path)

    # 获取根节点，并从中获取ModelVariables元素
    root = tree.getroot()
    model_variables = root.find('ModelVariables')

    # 迭代ScalarVariable元素
    for scalar_var in model_variables.findall('ScalarVariable'):
        # 获取元素的名称和值引用
        name = scalar_var.get('name')
        variability = scalar_var.get('variability')
        causality = scalar_var.get('causality')
        if variability == "parameter" and causality == "parameter" and name in var_dict:
            print(name)
            print(variability)
            print(causality)
            real_node = scalar_var.find('Real')
            # 设置startLine属性为100
            real_node.set("start", var_dict[name])

    # 将修改后的XML保存到文件中
    tree.write(xml_path+'result_init.xml')
