import xml.etree.ElementTree as ET


def xml_input(xml_path, variable_name_list, variable_val_list):
    # 解析XML文件为一个树
    tree = ET.parse(xml_path)

    # 获取根节点，并从中获取ModelVariables元素
    root = tree.getroot()
    model_variables = root.find('ModelVariables')

    # 迭代ScalarVariable元素
    for scalar_var in model_variables.findall('ScalarVariable'):
        # 获取元素的名称和值引用
        name = scalar_var.get('name')
        value_reference = scalar_var.get('valueReference')

        # 设置startLine属性为100
        scalar_var.set('startLine', '100')

    # 将修改后的XML保存到文件中
    tree.write('output.xml')
