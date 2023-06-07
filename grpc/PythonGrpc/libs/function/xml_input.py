import xml.etree.ElementTree as ET
from libs.function.grpc_log import log


def write_xml(xml_path, var_dict):
    # 解析XML文件为一个树
    result_xml_path = xml_path+"result_init.xml"
    log.info("(OMC)解析xml文件地址："+result_xml_path)

    try:
        tree = ET.parse(result_xml_path)
        # 获取根节点，并从中获取ModelVariables元素
        root = tree.getroot()
        model_variables = root.find('ModelVariables')
    except ET.ParseError as err:
        log.info(f"(OMC)解析XML文档时出现错误：{err}")
        return 1

    # 迭代ScalarVariable元素
    for scalar_var in model_variables.findall('ScalarVariable'):
        # 获取元素的名称和值引用
        name = scalar_var.get('name')
        if name in var_dict:
            log.info("(OMC)修改xml变量："+name)
            real_node = scalar_var.find('Real')
            # 设置startLine属性为100
            real_node.set("start", str(var_dict[name]))


    # 将修改后的XML保存到文件中
    tree.write(xml_path+'result_init.xml')
