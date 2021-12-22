# -- coding: utf-8 --
from library.OMPython import OMCSessionHelper, OMCSessionBase, logger
import subprocess
import sys, os, time, zmq
from library.OMPython.cdata_to_pydata import CdataToPYdata


class OMCSessionZMQ(OMCSessionHelper, OMCSessionBase):

    def __init__ (self, readonly=False, timeout=10.00, docker=None, dockerContainer=None, dockerExtraArgs=[],
                  dockerOpenModelicaPath="omc", dockerNetwork=None, address="127.0.0.1", port=23456, random_string="simtek", once=False):
        OMCSessionHelper.__init__(self)
        OMCSessionBase.__init__(self, readonly, interactivePort = port, random_string=random_string)
        # Locating and using the IOR
        if sys.platform != 'win32' or docker or dockerContainer:
            self._port_file = "openmodelica." + self._currentUser + ".port." + self._random_string
        else:
            self._port_file = "openmodelica.port." + self._random_string
        self._docker = docker
        self._dockerContainer = dockerContainer
        self._dockerExtraArgs = dockerExtraArgs
        self._dockerOpenModelicaPath = dockerOpenModelicaPath
        self._dockerNetwork = dockerNetwork
        self._create_omc_log_file("port")
        self._timeout = timeout
        self._port_file = os.path.join("/tmp" if docker else self._temp_dir, self._port_file).replace("\\", "/")
        self._interactivePort = port
        self._serverIPAddress = address
        self._once = once
        # set omc executable path and args
        self._set_omc_command([
            "--interactive=zmq",
            "--locale=C",
            "-z={0}".format(self._random_string)
        ])
        # start up omc executable, which is waiting for the ZMQ connection
        # 如果不是一次性链接，可判断是服务启动，则启动omc，
        if not self._once:
            self._start_omc_process(timeout)
            # connect to the running omc instance using ZMQ
        self._connect_to_omc(timeout)

    def __del__ (self):
        OMCSessionBase.__del__(self)

    def _connect_to_omc (self, timeout):
        self._omc_zeromq_uri = "file:///" + self._port_file
        # See if the omc server is running
        attempts = 0
        self._port = None
        while True:
            if self._dockerCid:
                try:
                    self._port = subprocess.check_output(["docker", "exec", self._dockerCid, "cat", self._port_file],
                                                         stderr=subprocess.DEVNULL if (sys.version_info > (
                                                             3, 0)) else subprocess.STDOUT).decode().strip()
                    break
                except:
                    pass
            else:
                # if os.path.isfile(self._port_file):
                #     # Read the port file
                #     with open(self._port_file, 'r') as f_p:
                #         self._port = f_p.readline()
                #     os.remove(self._port_file)
                #     break
                self._port = "tcp://" + self._serverIPAddress + ":" + str(self._interactivePort)
                break
            attempts += 1
            if attempts == 80.0:
                name = self._omc_log_file.name
                self._omc_log_file.close()
                logger.error("OMC Server did not start. Please start it! Log-file says:\n%s" % open(name).read())
                raise Exception(
                        "OMC Server did not start (timeout=%f). Could not open file %s" % (timeout, self._port_file))
            time.sleep(timeout / 80.0)

        self._port = self._port.replace("0.0.0.0", self._serverIPAddress)
        # logger.info(
        #         "OMC Server is up and running at {0} pid={1} cid={2}".format(self._omc_zeromq_uri,
        #                                                                      self._omc_process.pid,
        #                                                                      self._dockerCid))

        # Create the ZeroMQ socket and connect to OMC server
        import zmq
        context = zmq.Context.instance()
        self._omc = context.socket(zmq.REQ)
        self._omc.setsockopt(zmq.LINGER, 0)  # Dismisses pending messages if closed
        self._omc.setsockopt(zmq.IMMEDIATE, True)  # Queue messages only to completed connections
        self._omc.connect(self._port)

    def execute (self, command):
        ## check for process is running
        return self.sendExpression(command, parsed=False)

    def sendExpression (self, command, parsed=True):
        ## check for process is running
        # p = self._omc_process.poll()
        # if (p == None):
            # self._connect_to_omc(10)
        attempts = 0
        while True:
            try:
                self._omc.send_string(str(command), flags=zmq.NOBLOCK)
                break
            except zmq.error.Again:
                pass
            attempts += 1
            if attempts == 50.0:
                name = self._omc_log_file.name
                self._omc_log_file.close()
                raise Exception(
                        "No connection with OMC (timeout=%f). Log-file says: \n%s" % (
                            self._timeout, open(name).read()))
            time.sleep(self._timeout / 50.0)
        if command == "quit()":
            self._omc.close()
            self._omc = None
            return None
        else:
            result = self._omc.recv_string()
            if self._once:
                self._omc.close()
            if parsed is True:
                answer = CdataToPYdata(result)
                return answer
            else:
                return result
        # else:
        #     raise Exception("Process Exited, No connection with OMC. Create a new instance of OMCSession")

    def getComponents (self, class_name):
        return self.sendExpression("getComponents(" + class_name + ")")

    def getClassInformation (self, class_name):
        return self.sendExpression("getClassInformation(" + class_name + ")")

    def getComponentsList(self, class_name_list):
        data_list = []
        for i in class_name_list:
            Components_data = self.sendExpression("getComponents(" + i + ", useQuotes = true)")
            if Components_data != [''] and Components_data != "Error":
                data_list.extend(Components_data)
        return data_list

    def getComponentAnnotations (self, class_name):
        return self.sendExpression("getComponentAnnotations(" + class_name + ")")

    def getComponentAnnotationsList(self, class_name_list):
        data_list = []
        for i in class_name_list:
            Components_data = self.sendExpression("getComponentAnnotations(" + i + ", useQuotes = true)")
            if Components_data != [''] and Components_data != "Error":
                data_list.extend(Components_data)
        return data_list

    def getParameterNames (self, class_name):
        return self.sendExpression("getParameterNames(" + class_name + ")")

    def getComponentModifierNames (self, class_name, component_name):
        return self.sendExpression("getComponentModifierNames(" + class_name + ",\"" + component_name + "\")")

    def getComponentModifierNamesList (self, class_name_list, component_name):
        data_list = []
        for class_name in class_name_list:
            InheritedClasses_data = self.sendExpression("getComponentModifierNames(" + class_name + ",\"" + component_name + "\")")
            if InheritedClasses_data != ['']:
                data_list.extend(InheritedClasses_data)
        return data_list

    def getClassComment (self, class_name):
        return self.sendExpression("getClassComment(" + class_name + ")")

    def getInheritedClasses(self, class_name):
        InheritedClasses_data = self.sendExpression("getInheritedClasses(" + class_name + ")")
        return InheritedClasses_data

    def getInheritedClassesList(self, class_name_list):
        data_list = []
        for i in class_name_list:
            InheritedClasses_data = self.sendExpression("getInheritedClasses(" + i + ")")
            if InheritedClasses_data != ['']:
                data_list.extend(InheritedClasses_data)
        return data_list

    def getInheritedClassesListAll(self,class_name):
        data_list = class_name
        name_list = class_name
        while True:
            InheritedClassesData = self.getInheritedClassesList(name_list)
            if InheritedClassesData:
                data_list.extend(InheritedClassesData)
                name_list = InheritedClassesData
            else:
                break
        return data_list

    def getParameterNamesALL(self, class_name_list):
        data_list = []
        for i in class_name_list:
            getParameterNames_data = self.sendExpression("getParameterNames(" + i + ")")
            if getParameterNames_data != ['']:
                data_list.extend(getParameterNames_data)
        return data_list

    def getDiagramAnnotationList(self, class_name_list):
        data_list = []
        for i in class_name_list:
            DiagramAnnotation_data = self.sendExpression("getDiagramAnnotation(" + i + ")")
            if DiagramAnnotation_data != ['']:
                data_list.extend(DiagramAnnotation_data)
        return data_list

    def getIconAnnotationList(self, class_name_list):
        data_list = []
        for i in class_name_list:
            IconAnnotation_data = self.sendExpression("getIconAnnotation(" + i + ")")
            if IconAnnotation_data != [''] and len(IconAnnotation_data) >= 8:
                data_list.extend(IconAnnotation_data[-1])
        return data_list

    def getConnectionCountList(self, class_name_list):
        data_list = []
        for i in class_name_list:
            ConnectionCount_num = self.sendExpression("getConnectionCount(" + i + ")")
            if not ConnectionCount_num:
                data_list.append(0)
                continue
            data = int(ConnectionCount_num)
            data_list.append(data)
        return data_list

    def getNthConnection(self, class_name, num):
        data_list = []
        NthConnection_data = self.sendExpression("getNthConnection(" + class_name + "," + str(num) + ")")
        data_list.extend(NthConnection_data)
        return data_list

    def getNthConnectionAnnotation(self, class_name, num):
        data_list = []
        NthConnectionAnnotation_data = self.sendExpression("getNthConnectionAnnotation(" + class_name + "," + str(num) + ")")
        data_list.extend(NthConnectionAnnotation_data)
        return data_list

    def getNthConnectionList(self, class_name, num):
        data_list = []
        for i in range(num):
            NthConnection_data = self.sendExpression("getNthConnection(" + class_name + "," + str(i+1) + ")")
            data_list.extend(NthConnection_data)
        return data_list

    def getNthConnectionAnnotationList(self, class_name, num):
        data_list = []
        for i in range(num):
            NthConnectionAnnotation_data = self.sendExpression("getNthConnectionAnnotation(" + class_name + "," + str(i + 1) + ")")
            data_list.extend(NthConnectionAnnotation_data)
        return data_list

    def getInheritedClassesAllList(self, class_name):
        namelist = [class_name]
        w_name = namelist
        while True:
            ic_data = self.getInheritedClasses(w_name)
            if ic_data or ic_data != []:
                namelist.extend(ic_data)
            else:
                break
            w_name = ic_data
        return namelist

    def getDerivedClassModifierNames(self, class_name):
        getDerivedClassModifierNames_data = self.sendExpression("getDerivedClassModifierNames(" + class_name + ")")
        return getDerivedClassModifierNames_data

    def getDerivedClassModifierValue(self, class_name, modifier_name):
        getDerivedClassModifierValue_data = self.sendExpression("getDerivedClassModifierValue(" + class_name + "," + modifier_name + ")")
        return getDerivedClassModifierValue_data

    def getDerivedClassModifierValueList(self, class_name, modifier_name_list):
        data_list = []
        for modifier_name in modifier_name_list:
            getDerivedClassModifierValue_data = self.sendExpression("getDerivedClassModifierValue(" + class_name + "," + modifier_name + ")")
            data_list.insert(0, getDerivedClassModifierValue_data)
        return data_list

    def isEnumeration(self, parameter_name):
        isEnumeration_data = self.sendExpression("isEnumeration(" + parameter_name + ")")
        return isEnumeration_data

    def getEnumerationLiterals(self, parameter_name):
        getEnumerationLiterals_data = self.sendExpression("getEnumerationLiterals(" + parameter_name + ")")
        return getEnumerationLiterals_data

    def getParameterValue(self, component_name, modifier_name):
        getDerivedClassModifierValue_data = self.sendExpression("getParameterValue(" + component_name + ",\"" + modifier_name + "\")")
        return getDerivedClassModifierValue_data

    def getComponentModifierValue(self, class_name, modifier_name):
        cmd = "getComponentModifierValue(" + class_name + "," + modifier_name + ")"
        getDerivedClassModifierValue_data = self.sendExpression(cmd)
        return getDerivedClassModifierValue_data

    def setComponentModifierValue(self, class_name, parameter, value):
        code = "=" + value + ""
        if not value:
            code = "()"
        cmd = "setComponentModifierValue(" + class_name + ", " + parameter + ", $Code(" + code + "))"
        result = self.sendExpression(cmd)
        return result

    def setComponentProperties(self, class_name, component_name, final="false", protected="false", replaceable="false", variabilty="", inner="false", outer="false", causality=""):
        cmd_parameter_list = [class_name, ",", component_name, ",{", final, ",false,", protected, ",", replaceable, "},{\"", variabilty, "\"}", ",{", inner, ",", outer, "},{\"", causality, "\"}"]
        cmd = "setComponentProperties(" + "".join(cmd_parameter_list) + ")"
        cmd = cmd.replace("False", "false")
        cmd = cmd.replace("True", "true")
        result = self.sendExpression(cmd)
        return result

    def setComponentComment(self, class_name, component_name, annotation):
        cmd = "setComponentComment(" + class_name + "," + component_name + ",\"" + annotation + "\")"
        result = self.sendExpression(cmd)
        return result

    def setComponentDimensions(self, class_name, component_name, dimension):
        cmd = "setComponentDimensions(" + class_name + "," + component_name + ",{" + dimension + "})"
        result = self.sendExpression(cmd)
        return result

    def existClass(self, class_name):
        cmd = "existClass(" + class_name + ")"
        result = self.sendExpression(cmd)
        return result

    def copyClass(self, copied_class_name, class_name, parent_name):
        cmd = "copyClass(" + copied_class_name + ",\"" + class_name + "\"," + parent_name + ")"
        result = self.sendExpression(cmd)
        return result

    def deleteClass(self, class_name):
        cmd = "deleteClass(" + class_name + ")"
        result = self.sendExpression(cmd)
        return result

    def list(self, class_name, parsed=False):
        cmd = "list(" + class_name  + ")"
        data = self.sendExpression(cmd, parsed=parsed)
        return data

    def loadString(self, model_str, path="", merge="false"):
        cmd = "loadString(\"" + model_str + "\",\"" + path + "\",\"UTF-8\"" + "," + merge + ")"
        result = self.sendExpression(cmd)
        return result

    def listFile(self, package_name):
        cmd = "listFile(" + package_name + ",true)"
        result = self.sendExpression(cmd)
        return result

    def parseString(self, model_str, path=""):
        cmd = "parseString(\"" + model_str+ "\",\"" + path + "\")"
        result = self.sendExpression(cmd)
        return result

    def addClassAnnotation(self, class_name_all, annotate_str):
        cmd = "addClassAnnotation(" + class_name_all + ", annotate=" + annotate_str + ")"
        result = self.sendExpression(cmd)
        return result

    def addConnection(self, class_name_all, connect_start, connect_end, line_points, color="0,0,127"):
        # addConnection(integrator.y,PI.u_ff,qq.PID_Controller,annotate=Line(points={{-42,30},{-40,30},{-40,-22}},color={0,0,127}))
        line_points = ",".join(["{" + i + "}" for i in line_points])
        annotate = "annotate=Line(points={" + line_points + "},color={" + color + "}))"
        cmd = "addConnection(" + connect_start + ","+ connect_end + "," + class_name_all + "," + annotate
        result = self.sendExpression(cmd)
        return result

    def updateConnectionNames(self, class_name_all, from_name, to_name, from_name_new, to_name_new):
        cmd = "updateConnectionNames(\"" + class_name_all + "\",\""+ from_name + "\",\"" + to_name + "\",\"" + from_name_new + "\",\"" + to_name_new + "\")"
        result = self.sendExpression(cmd)
        return result

    def updateConnectionAnnotation (self, class_name_all, connect_start, connect_end, line_points, color="0,0,127"):
        # updateConnectionAnnotation(qq.Scenario1_Status, "gasGlassFurnance.pro_out", "productSink.port_a","annotate=$annotation(Line(points={{-40,14},{-24,14},{-24,32},{-12,32}},color={255,170,255}))")
        line_points = ",".join(["{" + i + "}" for i in line_points])
        annotate = "$annotation(Line(points={" + line_points + "},color={" + color + "}))\""
        cmd = "updateConnectionAnnotation(" + class_name_all + ", \"" + connect_start + "\", \"" + connect_end +  "\", \"annotate=" + annotate + ")"
        result = self.sendExpression(cmd)
        return result

    def deleteConnection(self, class_name_all, connect_start, connect_end):
        cmd = "deleteConnection(" + connect_start + ","+ connect_end + "," + class_name_all + ")"
        result = self.sendExpression(cmd)
        return result

    def addComponent(self, new_component_name, old_component_name, class_name_all, origin, extent, rotation):
        # annotate=Placement(visible=true, transformation=transformation(origin={-72,-64}, extent={{-10,-10},{10,10}}, rotation=0))
        annotate = "annotate=Placement(visible=true, transformation=transformation(origin={" + origin + "}, extent={{" + extent[0] + "},{" + extent[1] + "}}, rotation=" + rotation + "))"
        cmd = "addComponent(" + new_component_name + ","+ old_component_name + "," + class_name_all + "," + annotate + ")"
        result = self.sendExpression(cmd)
        return result

    def deleteComponent(self, component_name, class_name_all):
        cmd = "deleteComponent(" + component_name + ","+ class_name_all + ")"
        result = self.sendExpression(cmd)
        return result

    def updateComponent(self, component_name, component_model_name, class_name_all, origin, extent, rotation):
        # updateComponent(PI,Modelica.Blocks.Continuous.LimPID,ENN.Examples.PID_Controller10086,annotate=Placement(visible=true, transformation=transformation(origin={-46,-10}, extent={{10,10},{-10,-10}}, rotation=180)))
        annotate = "annotate=Placement(visible=true, transformation=transformation(origin={" + origin + "}, extent={{" + extent[0] + "},{" + extent[1] + "}}, rotation=" + rotation + "))"
        cmd = "updateComponent(" + component_name + "," + component_model_name + ","+ class_name_all + "," + annotate + ")"
        result = self.sendExpression(cmd)
        return result


if __name__ == '__main__':
    def loadString (model_str, path, merge="false"):
        cmd = "loadString(\"" + model_str + "\",\"" + path + "\",\"UTF-8\"" + ",false)"
        # result = self.sendExpression(cmd)
        return cmd
    loadString("within ENN.Examples.PID_Controller; partial model q123456 extends ENN.Examples.Scenario1_Status;  end q123456;", "/a/b/c")
