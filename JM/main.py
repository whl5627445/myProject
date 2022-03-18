# -- coding: utf-8 --
"""
    compile_fmu(class_name, file_name=[], compiler='auto', target='me', version='2.0', compiler_options={}, compile_to='.', compiler_log_level='warning',separate_process=True, jvm_args='')
    将 Modelica 模型编译为 FMU。
    必须传递模型类名，所有其他参数都有默认值。不同的情况是：
    * 仅传递 class_name：- 假定类在 MODELICAPATH 中。
    * 传递 Class_name 和 file_name： - file_name 可以是单个路径作为字符串或路径列表（字符串）。路径可以是文件或库路径。
    - 默认编译器设置为“自动”，这意味着将根据模型文件结尾选择合适的编译器，
    即如果在 file_name 列表中找到 .mo 文件则为 ModelicaCompiler，
    如果在 file_name 列表中找到 .mop 文件则为 OptimicaCompiler 编译器目标默认为“me”，
    这意味着共享文件包含用于模型交换 API 的 FMI。
    将此参数设置为 'cs' 将生成包含用于协同仿真 API 的 FMI 的 FMU。
    参数 :: class_name - 模型类的名称。 file_name - 模型文件和/或库的路径（字符串）或路径（字符串列表）。
    默认值：空列表。编译器 - 用于编译模型的编译器。不同的选项是：
    - 'auto'：根据文件结尾自动选择编译器
    - 'modelica'：使用 ModelicaCompiler
    - 'optimica'：使用 OptimicaCompiler 默认：'auto'目标
    - 在 Python15 编译器目标中使用模型。可能的值是“me”、“cs”或“me + cs”。
    默认值：'me' version - FMI 版本。有效选项为“1.0”和“2.0”。默认值：'2.0'
    compiler_options - 编译器选项。默认值：空字典。
    compile_to - 指定目标文件或目录。如果是文件，则将创建任何中间目录（如果它们不存在）。如果是目录，则给定的路径必须存在。默认值：当前目录。
    compiler_log_level - 设置编译器的日志记录。采用带有日志输出的逗号分隔列表。日志输出以一个标志开头：
    'warning' / 'w'、'error' / 'e'、'info' / 'i' 或 'debug' / 'd'。
    可以通过附加带有冒号和文件名的标志将日志写入文件。默认值：'warning'
    separator_process - 在单独的进程中运行模型的编译。检查环境变量（按此顺序）：
    1. SEPARATE_PROCESS_JVM
    2. JAVA_HOME 定位要使用的 Java 安装。例如（在 Windows 上）
    这可能是：SEPARATE_PROCESS_JVM = C:\Program Files\Java\jdk1.6.0_37 默认值：True
    jvm_args - 在单独的进程中编译时传递给 JVM 的参数字符串。默认值：空 stringReturns :: 编译结果，表示已创建的 FMU 的名称以及引发的警告列表。
    """
import logging

from pymodelica import compile_fmu
from pyfmi import load_fmu
import json
import socket
logging.basicConfig(level=logging.DEBUG,#控制台打印的日志级别
                    format='%(asctime)s - %(pathname)s[line:%(lineno)d] - %(levelname)s: %(message)s'
                    #日志格式
                    )

def service():
    s = socket.socket()
    # host = socket.gethostname()
    host = "0.0.0.0"
    port = 56789
    logging.debug(host)
    s.bind((host, port))
    s.listen(5)
    while True:
        print "start ok"
        soc, addr = s.accept()
        data = json.loads(soc.recv(4096))
        logging.debug(data)
        if data:
            try:
                result_file_path = "/" + data["result_file_path"]
                print result_file_path
                if data["type"] == "compile":
                    try:
                        fmu = compile_fmu(class_name=data["modelname"], file_name=data["mo_path"], compile_to=result_file_path)
                        soc.send('ok')
                    except Exception as e:
                        print e
                        soc.send(str(e))
                elif data["type"] == "simulate":
                    try:
                        vdp = load_fmu(result_file_path + data["modelname"] + ".fmu")
                        opts = vdp.simulate_options()
                        opts["ncp"] = data["ncp"]
                        opts["result_file_name"] = result_file_path + "/" + data["result_name"]
                        start_time = data.get("start_time", 0.0)
                        final_time = data.get("final_time", 10.0)
                        res = vdp.simulate(start_time=start_time, final_time=final_time, options=opts)
                    except Exception as e:
                        print e
                        soc.send(str(e))
                    soc.send('ok')
            except Exception as e:
                print e
                soc.send(str(e))
        soc.close()


if __name__ == '__main__':
    service()
