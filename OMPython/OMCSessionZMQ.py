from OMPython import OMCSessionHelper, OMCSessionBase, logger
import subprocess
import sys, os, time, zmq
from OMPython.cdata_to_pydata import CdataToPYdata


class OMCSessionZMQ(OMCSessionHelper, OMCSessionBase):

    def __init__ (self, readonly=False, timeout=10.00, docker=None, dockerContainer=None, dockerExtraArgs=[],
                  dockerOpenModelicaPath="omc", dockerNetwork=None, port=None, random_string=None):
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
        # set omc executable path and args
        self._set_omc_command([
            "--interactive=zmq",
            "--locale=C",
            "-z={0}".format(self._random_string)
        ])
        # start up omc executable, which is waiting for the ZMQ connection
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
                if os.path.isfile(self._port_file):
                    # Read the port file
                    with open(self._port_file, 'r') as f_p:
                        self._port = f_p.readline()
                    os.remove(self._port_file)
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
        logger.info(
                "OMC Server is up and running at {0} pid={1} cid={2}".format(self._omc_zeromq_uri,
                                                                             self._omc_process.pid,
                                                                             self._dockerCid))

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
        p = self._omc_process.poll()
        if (p == None):
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
                if parsed is True:
                    answer = CdataToPYdata(result)
                    return answer
                else:
                    return result
        else:
            raise Exception("Process Exited, No connection with OMC. Create a new instance of OMCSession")

    def getComponents (self, class_name):
        return self.sendExpression("getComponents(" + class_name + ")")

    def getComponentAnnotations (self, class_name):
        return self.sendExpression("getComponentAnnotations(" + class_name + ")")

    def getParameterNames (self, class_name):
        return self.sendExpression("getParameterNames(" + class_name + ")")

    def getComponentModifierNames (self, class_name, component_name):
        return self.sendExpression("getComponentModifierNames(" + class_name + ",\"" + component_name + "\")")

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

    def getDiagramAnnotationList(self, class_name_list):
        data_list = []
        for i in class_name_list:
            DiagramAnnotation_data = self.sendExpression("getDiagramAnnotation(" + i + ")")
            if DiagramAnnotation_data != ['']:
                data_list.extend(DiagramAnnotation_data)
        return data_list

    def getComponentsList(self, class_name_list):
        data_list = []
        for i in class_name_list:
            Components_data = self.sendExpression("getComponents(" + i + ", useQuotes = true)")
            if Components_data != ['']:
                data_list.extend(Components_data)
        return data_list

    def getComponentAnnotationsList(self, class_name_list):
        data_list = []
        for i in class_name_list:
            ComponentAnnotations_data = self.sendExpression("getComponentAnnotations(" + i + ")")
            if ComponentAnnotations_data != ['']:
                data_list.extend(ComponentAnnotations_data)
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
            data = int(ConnectionCount_num)
            data_list.append(data)
        return data_list

    def getNthConnectionList(self, class_name_list, num):
        data_list = []
        for i in range(num):
            NthConnection_data = self.sendExpression("getNthConnection(" + class_name_list[i] + "," + str(i+1) + ")")
            data_list.extend(NthConnection_data)
        return data_list

    def getNthConnectionAnnotationList(self, class_name_list, num):
        data_list = []
        for i in range(num):
            NthConnectionAnnotation_data = self.sendExpression("getNthConnectionAnnotation(" + class_name_list[i] + "," + str(i + 1) + ")")
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
        getDerivedClassModifierValue_data = self.sendExpression("getComponentModifierValue(" + class_name + "," + modifier_name + ")")
        return getDerivedClassModifierValue_data

    def setComponentModifierValue(self, class_name, parameter, value):
        cmd = "setComponentModifierValue(" + class_name + "," + parameter + ",$Code(=" + value + "))"
        result = self.sendExpression(cmd)
        return result

    def setComponentProperties(self, class_name, component_name, final="false", protected="false", replaceable="false", variabilty="", inner="false", outer="false", causality=""):
        # self.loadFile("/home/simtek/dev/public/UserFiles/UploadFile/tom/1631690039.291318/ENN.mo")
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

