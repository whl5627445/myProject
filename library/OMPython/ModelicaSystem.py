# -- coding: utf-8 --
from __future__ import absolute_import
from __future__ import division
from __future__ import print_function
import numpy as np
from library.OMPython import OMCSessionZMQ
import os, sys
import subprocess
from builtins import int, range
from xml.etree.ElementTree import ElementTree as ET
from copy import deepcopy
import csv
from collections import OrderedDict

platform = sys.platform

class ModelicaSystem(object):
    def __init__ (self, fileName=None, modelName=None, lmodel=[], commandLineOptions=None, Port=None, random_string=None):  # 1
        """
        "constructor"
        It initializes to load file and build a model, generating object, exe, xml, mat, and json files. etc. It can be called :
            •without any arguments: In this case it neither loads a file nor build a model. This is useful when a FMU needed to convert to Modelica model
            •with two arguments as file name with ".mo" extension and the model name respectively
            •with three arguments, the first and second are file name and model name respectively and the third arguments is Modelica standard library to load a model, which is common in such models where the model is based on the standard library. For example, here is a model named "dcmotor.mo" below table 4-2, which is located in the directory of OpenModelica at "C:\\OpenModelica1.9.4-dev.beta2\\share\\doc\\omc\\testmodels".
        Note: If the model file is not in the current working directory, then the path where file is located must be included together with file name. Besides, if the Modelica model contains several different models within the same package, then in order to build the specific model, in second argument, user must put the package name with dot(.) followed by specific model name.
        ex: myModel = ModelicaSystem("ModelicaModel.mo", "modelName")
        """

        if fileName is None and modelName is None and not lmodel:  # all None
            self.getconn = OMCSessionZMQ(port=Port, random_string=random_string)
            return

        if fileName is None:
            return "File does not exist"
        self.tree = None

        self.quantitiesList = []
        self.paramlist = {}
        self.inputlist = {}
        self.outputlist = {}
        self.continuouslist = {}
        self.simulateOptions = {}
        self.overridevariables = {}
        self.simoptionsoverride = {}
        self.linearOptions = {'startTime': 0.0, 'stopTime': 1.0, 'numberOfIntervals': 500, 'stepSize': 0.002,
                              'tolerance': 1e-8}
        self.optimizeOptions = {'startTime': 0.0, 'stopTime': 1.0, 'numberOfIntervals': 500, 'stepSize': 0.002,
                                'tolerance': 1e-8}
        self.linearquantitiesList = []  # linearization  quantity list
        self.linearparameters = {}
        self.linearinputs = []  # linearization input list
        self.linearoutputs = []  # linearization output list
        self.linearstates = []  # linearization  states list

        self.getconn = OMCSessionZMQ(port=Port, random_string=random_string)

        ## set commandLineOptions if provided by users
        if commandLineOptions is not None:
            exp = "".join(["setCommandLineOptions(", "\"", commandLineOptions, "\"", ")"])
            self.getconn.sendExpression(exp)

        self.xmlFile = None
        self.lmodel = lmodel  # may be needed if model is derived from other model
        self.modelName = modelName  # Model class name
        self.fileName = fileName  # Model file/package name
        self.inputFlag = False  # for model with input quantity
        self.simulationFlag = False  # if the model is simulated?
        self.linearizationFlag = False
        self.outputFlag = False
        self.csvFile = ''  # for storing inputs condition
        self.resultfile = ""  # for storing result file
        if not os.path.exists(self.fileName):  # if file does not eixt
            print("File Error:" + os.path.abspath(self.fileName) + " does not exist!!!")
            return

        (head, tail) = os.path.split(self.fileName)  # to store directory/path and file)
        self.currDir = os.getcwd()
        self.modelDir = head
        self.fileName_ = tail

        if not self.modelDir:
            file_ = os.path.exists(self.fileName_)
            if (file_):  # execution from path where file is located
                self.__loadingModel()
            else:
                print("Error: File does not exist!!!")

        else:
            os.chdir(self.modelDir)
            file_ = os.path.exists(self.fileName_)
            self.model = self.fileName_[:-3]
            if (self.fileName_):  # execution from different path
                os.chdir(self.currDir)
                self.__loadingModel()
            else:
                print("Error: File does not exist!!!")

    def __del__ (self):
        if self.getconn is not None:
            self.requestApi('quit')

    # for loading file/package, loading model and building model
    def __loadingModel (self):
        # load file
        loadfileError = ''
        loadfileResult = self.requestApi("loadFile", self.fileName)
        loadfileError = self.requestApi("getErrorString")

        # print the notification to users
        if loadfileResult == True and loadfileError:
            print(loadfileError)

        if loadfileResult == False:
            specError = 'Parser error: Unexpected token near: optimization (IDENT)'
            if specError in loadfileError:
                self.requestApi("setCommandLineOptions", '"+g=Optimica"')
                self.requestApi("loadFile", self.fileName)
            else:
                print('loadFile Error: ' + loadfileError)
                return

        # load Modelica standard libraries or Modelica files if needed
        for element in self.lmodel:
            if element is not None:
                loadmodelError = ''
                if element.endswith(".mo"):
                    loadModelResult = self.requestApi("loadFile", element)
                    loadmodelError = self.requestApi('getErrorString')
                else:
                    loadModelResult = self.requestApi("loadModel", element)
                    loadmodelError = self.requestApi('getErrorString')
                if loadmodelError:
                    print(loadmodelError)
        self.buildModel()

    def buildModel (self):
        # buildModelResult=self.getconn.sendExpression("buildModel("+ mName +")")
        buildModelResult = self.requestApi("buildModel", self.modelName)
        buildModelError = self.requestApi("getErrorString")
        if '' in buildModelResult:
            print(buildModelError)
            return
        self.xmlFile = os.path.join(os.path.dirname(buildModelResult[0]), buildModelResult[1]).replace("\\", "/")
        return self.xmlparse()

    def sendExpression (self, expr, parsed=False, new_parsed=True):
        return self.getconn.sendExpression(expr, parsed, new_parsed)


    # request to OMC
    def requestApi (self, apiName, entity=None, properties=None):  # 2
        if (entity is not None and properties is not None):
            exp = '{}({}, {})'.format(apiName, entity, properties)
        elif entity is not None and properties is None:
            if (apiName == "loadFile" or apiName == "importFMU"):
                exp = '{}("{}")'.format(apiName, entity)
            else:
                exp = '{}({})'.format(apiName, entity)
        else:
            exp = '{}()'.format(apiName)
        try:
            res = self.getconn.sendExpression(exp)
        except Exception as e:
            print(e)
            res = None
        return res

    def xmlparse (self):
        if (os.path.exists(self.xmlFile)):
            self.tree = ET.parse(self.xmlFile)
            self.root = self.tree.getroot()
            rootCQ = self.root
            for attr in rootCQ.iter('DefaultExperiment'):
                self.simulateOptions["startTime"] = attr.get('startTime')
                self.simulateOptions["stopTime"] = attr.get('stopTime')
                self.simulateOptions["stepSize"] = attr.get('stepSize')
                self.simulateOptions["tolerance"] = attr.get('tolerance')
                self.simulateOptions["solver"] = attr.get('solver')

            for sv in rootCQ.iter('ScalarVariable'):
                scalar = {}
                scalar["name"] = sv.get('name')
                scalar["changable"] = sv.get('isValueChangeable')
                scalar["description"] = sv.get('description')
                scalar["variability"] = sv.get('variability')
                scalar["causality"] = sv.get('causality')
                scalar["alias"] = sv.get('alias')
                scalar["aliasvariable"] = sv.get('aliasVariable')
                ch = list(sv)
                start = None
                for att in ch:
                    start = att.get('start')
                scalar["start"] = start

                if (self.linearizationFlag == False):
                    if (scalar["variability"] == "parameter"):
                        self.paramlist[scalar["name"]] = scalar["start"]
                    if (scalar["variability"] == "continuous"):
                        self.continuouslist[scalar["name"]] = scalar["start"]
                    if (scalar["causality"] == "input"):
                        self.inputlist[scalar["name"]] = scalar["start"]
                    if (scalar["causality"] == "output"):
                        self.outputlist[scalar["name"]] = scalar["start"]

                if self.linearizationFlag:
                    if scalar["variability"] == "parameter":
                        self.linearparameters[scalar["name"]] = scalar["start"]
                    if scalar["alias"] == "alias":
                        name = scalar["name"]
                        if name[1] == 'x':
                            self.linearstates.append(name[3:-1])
                        if name[1] == 'u':
                            self.linearinputs.append(name[3:-1])
                        if name[1] == 'y':
                            self.linearoutputs.append(name[3:-1])
                    self.linearquantitiesList.append(scalar)
                else:
                    self.quantitiesList.append(scalar)
        else:
            print("Error: ! XML file not generated")
            return

    def getQuantities (self, names=None):  # 3
        """
        变量的信息稍微详细一些
        This method returns list of dictionaries. It displays details of quantities such as name, value, changeable, and description, where changeable means  if value for corresponding quantity name is changeable or not. It can be called :
        usage:
        >>> getQuantities()
        >>> getQuantities("Name1")
        >>> getQuantities(["Name1","Name2"])
        """
        if names is None:
            return self.quantitiesList
        elif isinstance(names, str):
            return [x for x in self.quantitiesList if x["name"] == names]
        elif isinstance(names, list):
            return [x for y in names for x in self.quantitiesList if x["name"] == y]

    def getContinuous (self, names=None):  # 4
        """
        变量名字和变量值
        This method returns dict. The key is continuous names and value is corresponding continuous value.
        usage:
        >>> getContinuous()
        >>> getContinuous("Name1")
        >>> getContinuous(["Name1","Name2"])
        """
        if not self.simulationFlag:
            if names is None:
                return self.continuouslist
            elif isinstance(names, str):
                return [self.continuouslist.get(names, "NotExist")]
            elif isinstance(names, list):
                return [self.continuouslist.get(x, "NotExist") for x in names]
        else:
            if names is None:
                for i in self.continuouslist:
                    try:
                        value = self.getSolutions(i)
                        self.continuouslist[i] = value[0][-1]
                    except Exception:
                        print(i, "could not be computed")
                return self.continuouslist

            elif isinstance(names, str):
                if names in self.continuouslist:
                    value = self.getSolutions(names)
                    self.continuouslist[names] = value[0][-1]
                    return [self.continuouslist.get(names)]
                else:
                    return (names, "  is not continuous")

            elif isinstance(names, list):
                valuelist = []
                for i in names:
                    if i in self.continuouslist:
                        value = self.getSolutions(i)
                        self.continuouslist[i] = value[0][-1]
                        valuelist.append(value[0][-1])
                    else:
                        return (i, "  is not continuous")
                return valuelist

    def getParameters (self, names=None):  # 5
        """
        参数名和参数值
        This method returns dict. The key is parameter names and value is corresponding parameter value.
        If name is None then the function will return dict which contain all parameter names as key and value as corresponding values.
        usage:
        >>> getParameters()
        >>> getParameters("Name1")
        >>> getParameters(["Name1","Name2"])
        """
        if (names == None):
            return self.paramlist
        elif (isinstance(names, str)):
            return [self.paramlist.get(names, "NotExist")]
        elif (isinstance(names, list)):
            return ([self.paramlist.get(x, "NotExist") for x in names])

    def getlinearParameters (self, names=None):  # 5
        """
        This method returns dict. The key is parameter names and value is corresponding parameter value.
        If *name is None then the function will return dict which contain all parameter names as key and value as corresponding values. eg., getParameters()
        Otherwise variable number of arguments can be passed as parameter name in string format separated by commas. eg., getParameters('paraName1', 'paraName2')
        """
        if not names:
            return self.linearparameters
        elif isinstance(names, str):
            return [self.linearparameters.get(names, "NotExist")]
        else:
            return [self.linearparameters.get(x, "NotExist") for x in names]

    def getInputs (self, names=None):  # 6
        """
        This method returns dict. The key is input names and value is corresponding input value.
        If *name is None then the function will return dict which contain all input names as key and value as corresponding values. eg., getInputs()
        Otherwise variable number of arguments can be passed as input name in string format separated by commas. eg., getInputs('iName1', 'iName2')
        """
        if (names == None):
            return self.inputlist
        elif (isinstance(names, str)):
            return [self.inputlist.get(names, "NotExist")]
        elif (isinstance(names, list)):
            return ([self.inputlist.get(x, "NotExist") for x in names])

    def getOutputs (self, names=None):  # 7
        """
        This method returns dict. The key is output names and value is corresponding output value.
        If name is None then the function will return dict which contain all output names as key and value as corresponding values. eg., getOutputs()
        usage:
        >>> getOutputs()
        >>> getOutputs("Name1")
        >>> getOutputs(["Name1","Name2"])
        """
        if not self.simulationFlag:
            if (names == None):
                return self.outputlist
            elif (isinstance(names, str)):
                return [self.outputlist.get(names, "NotExist")]
            else:
                return ([self.outputlist.get(x, "NotExist") for x in names])
        else:
            if (names == None):
                for i in self.outputlist:
                    value = self.getSolutions(i)
                    self.outputlist[i] = value[0][-1]
                return self.outputlist
            elif (isinstance(names, str)):
                if names in self.outputlist:
                    value = self.getSolutions(names)
                    self.outputlist[names] = value[0][-1]
                    return [self.outputlist.get(names)]
                else:
                    return (names, " is not Output")
            elif (isinstance(names, list)):
                valuelist = []
                for i in names:
                    if i in self.outputlist:
                        value = self.getSolutions(i)
                        self.outputlist[i] = value[0][-1]
                        valuelist.append(value[0][-1])
                    else:
                        return (i, "is not Output")
                return valuelist

    def getSimulationOptions (self, names=None):  # 8
        """
        模拟选项的键值
        This method returns dict. The key is simulation option names and value is corresponding simulation option value.
        If name is None then the function will return dict which contain all simulation option names as key and value as corresponding values. eg., getSimulationOptions()
        usage:
        >>> getSimulationOptions()
        >>> getSimulationOptions("Name1")
        >>> getSimulationOptions(["Name1","Name2"])
        """
        if (names == None):
            return self.simulateOptions
        elif (isinstance(names, str)):
            return [self.simulateOptions.get(names, "NotExist")]
        elif (isinstance(names, list)):
            return ([self.simulateOptions.get(x, "NotExist") for x in names])

    def getLinearizationOptions (self, names=None):  # 9
        """
        线性模拟的选项名称
        This method returns dict. The key is linearize option names and value is corresponding linearize option value.
        If name is None then the function will return dict which contain all linearize option names as key and value as corresponding values. eg., getLinearizationOptions()
        usage:
        >>> getLinearizationOptions()
        >>> getLinearizationOptions("Name1")
        >>> getLinearizationOptions(["Name1","Name2"])
        """
        if (names == None):
            return self.linearOptions
        elif (isinstance(names, str)):
            return [self.linearOptions.get(names, "NotExist")]
        elif (isinstance(names, list)):
            return ([self.linearOptions.get(x, "NotExist") for x in names])

    def getOptimizationOptions (self, names=None):  # 10
        """
        usage:
        >>> getOptimizationOptions()
        >>> getOptimizationOptions("Name1")
        >>> getOptimizationOptions(["Name1","Name2"])
        """
        if (names == None):
            return self.optimizeOptions
        elif (isinstance(names, str)):
            return [self.optimizeOptions.get(names, "NotExist")]
        elif (isinstance(names, list)):
            return ([self.optimizeOptions.get(x, "NotExist") for x in names])

    # to simulate or re-simulate model
    def simulate (self, resultfile=None, simflags=None):  # 11
        """
        模拟函数
        This method simulates model according to the simulation options.
        usage
        >>> simulate()
        >>> simulate(resultfile="a.mat")
        >>> simulate(simflags="-noEventEmit -noRestart -override=e=0.3,g=10) set runtime simulation flags")
        """
        if resultfile is None:
            r = ""
            self.resultfile = "".join([self.modelName, "_res.mat"])
        else:
            r = " -r=" + resultfile
            self.resultfile = resultfile

        # allow runtime simulation flags from user input
        if simflags is None:
            simflags = ""
        else:
            simflags = " " + simflags;

        if (self.overridevariables or self.simoptionsoverride):
            tmpdict = self.overridevariables.copy()
            tmpdict.update(self.simoptionsoverride)
            values1 = ','.join("%s=%s" % (key, val) for (key, val) in list(tmpdict.items()))
            override = " -override=" + values1
        else:
            override = ""

        if self.inputFlag:  # if model has input quantities
            for i in self.inputlist:
                val = self.inputlist[i]
                if val == None:
                    val = [(float(self.simulateOptions["startTime"]), 0.0),
                           (float(self.simulateOptions["stopTime"]), 0.0)]
                    self.inputlist[i] = [(float(self.simulateOptions["startTime"]), 0.0),
                                         (float(self.simulateOptions["stopTime"]), 0.0)]
                if float(self.simulateOptions["startTime"]) != val[0][0]:
                    print("!!! startTime not matched for Input ", i)
                    return
                if float(self.simulateOptions["stopTime"]) != val[-1][0]:
                    print("!!! stopTime not matched for Input ", i)
                    return
                if val[0][0] < float(self.simulateOptions["startTime"]):
                    print('Input time value is less than simulation startTime for inputs', i)
                    return
            self.__simInput()  # create csv file
            csvinput = " -csvInput=" + self.csvFile
        else:
            csvinput = ""

        if (platform.system() == "Windows"):
            getExeFile = os.path.join(os.getcwd(), '{}.{}'.format(self.modelName, "exe")).replace("\\", "/")
        else:
            getExeFile = os.path.join(os.getcwd(), self.modelName).replace("\\", "/")

        if (os.path.exists(getExeFile)):
            cmd = getExeFile + override + csvinput + r + simflags
            # print(cmd)
            if (platform.system() == "Windows"):
                omhome = os.path.join(os.environ.get("OPENMODELICAHOME"))
                dllPath = os.path.join(omhome, "bin").replace("\\", "/") + os.pathsep + os.path.join(omhome,
                                                                                                     "lib/omc").replace(
                        "\\", "/") + os.pathsep + os.path.join(omhome, "lib/omc/cpp").replace("\\",
                                                                                              "/") + os.pathsep + os.path.join(
                        omhome, "lib/omc/omsicpp").replace("\\", "/")
                my_env = os.environ.copy()
                my_env["PATH"] = dllPath + os.pathsep + my_env["PATH"]
                p = subprocess.Popen(cmd, env=my_env)
                p.wait()
                p.terminate()
            else:
                os.system(cmd)
            self.simulationFlag = True

        else:
            raise Exception("Error: application file not generated yet")

    # to extract simulation results
    def getSolutions (self, varList=None, resultfile=None):  # 12
        """
        读取模拟结果文件， numpy类型的返回
        This method returns tuple of numpy arrays. It can be called:
            •with a list of quantities name in string format as argument: it returns the simulation results of the corresponding names in the same order. Here it supports Python unpacking depending upon the number of variables assigned.
        usage:
        >>> getSolutions()
        >>> getSolutions("Name1")
        >>> getSolutions(["Name1","Name2"])
        >>> getSolutions(resultfile="c:/a.mat")
        >>> getSolutions("Name1",resultfile=""c:/a.mat"")
        >>> getSolutions(["Name1","Name2"],resultfile=""c:/a.mat"")
        """
        if (resultfile == None):
            resFile = self.resultfile
        else:
            resFile = resultfile

        # check for result file exits
        if (not os.path.exists(resFile)):
            print("Error: Result file does not exist")
            return
            # exit()
        else:
            if (varList == None):
                # validSolution = ['time'] + self.__getInputNames() + self.__getContinuousNames() + self.__getParameterNames()
                validSolution = self.getconn.sendExpression("readSimulationResultVars(\"" + resFile + "\")")
                self.getconn.sendExpression("closeSimulationResultFile()")
                return validSolution
            elif (isinstance(varList, str)):
                if (varList not in [l["name"] for l in self.quantitiesList] and varList != "time"):
                    print('!!! ', varList, ' does not exist\n')
                    return
                exp = "readSimulationResult(\"" + resFile + '",{' + varList + "})"
                res = self.getconn.sendExpression(exp)
                npRes = np.array(res)
                exp2 = "closeSimulationResultFile()"
                self.getconn.sendExpression(exp2)
                return npRes
            elif (isinstance(varList, list)):
                # varList, = varList
                for v in varList:
                    if v == "time":
                        continue
                    if v not in [l["name"] for l in self.quantitiesList]:
                        print('!!! ', v, ' does not exist\n')
                        return
                variables = ",".join(varList)
                exp = "readSimulationResult(\"" + resFile + '",{' + variables + "})"
                res = self.getconn.sendExpression(exp)
                npRes = np.array(res)
                exp2 = "closeSimulationResultFile()"
                self.getconn.sendExpression(exp2)
                return npRes

    def strip_space (self, name):
        if (isinstance(name, str)):
            return name.replace(" ", "")
        elif (isinstance(name, list)):
            return [x.replace(" ", "") for x in name]

    def setMethodHelper (self, args1, args2, args3, args4=None):
        """
        Helper function for setParameter(),setContinuous(),setSimulationOptions(),setLinearizationOption(),setOptimizationOption()
        args1 - string or list of string given by user
        args2 - dict() containing the values of different variables(eg:, parameter,continuous,simulation parameters)
        args3 - function name (eg; continuous, parameter, simulation, linearization,optimization)
        args4 - dict() which stores the new override variables list,
        """
        if (isinstance(args1, str)):
            args1 = self.strip_space(args1)
            value = args1.split("=")
            if value[0] in args2:
                args2[value[0]] = value[1]
                if (args4 != None):
                    args4[value[0]] = value[1]
            else:
                print(value[0], "!is not a", args3, "variable")
                return
        elif (isinstance(args1, list)):
            args1 = self.strip_space(args1)
            for var in args1:
                value = var.split("=")
                if value[0] in args2:
                    args2[value[0]] = value[1]
                    if (args4 != None):
                        args4[value[0]] = value[1]
                else:
                    print(value[0], "!is not a", args3, "variable")
                    return

    def setContinuous (self, cvals):  # 13
        """
        This method is used to set continuous values. It can be called:
        with a sequence of continuous name and assigning corresponding values as arguments as show in the example below:
        usage
        >>> setContinuous("Name=value")
        >>> setContinuous(["Name1=value1","Name2=value2"])
        """
        return self.setMethodHelper(cvals, self.continuouslist, "continuous", self.overridevariables)

    def setParameters (self, pvals):  # 14
        """
        This method is used to set parameter values. It can be called:
        with a sequence of parameter name and assigning corresponding value as arguments as show in the example below:
        usage
        >>> setParameters("Name=value")
        >>> setParameters(["Name1=value1","Name2=value2"])
        """
        return self.setMethodHelper(pvals, self.paramlist, "parameter", self.overridevariables)

    def setSimulationOptions (self, simOptions):  # 16
        """
        This method is used to set simulation options. It can be called:
        with a sequence of simulation options name and assigning corresponding values as arguments as show in the example below:
        usage
        >>> setSimulationOptions("Name=value")
        >>> setSimulationOptions(["Name1=value1","Name2=value2"])
        """
        return self.setMethodHelper(simOptions, self.simulateOptions, "simulation-option", self.simoptionsoverride)

    def setLinearizationOptions (self, linearizationOptions):  # 18
        """
        This method is used to set linearization options. It can be called:
        with a sequence of linearization options name and assigning corresponding value as arguments as show in the example below
        usage
        >>> setLinearizationOptions("Name=value")
        >>> setLinearizationOptions(["Name1=value1","Name2=value2"])
        """
        return self.setMethodHelper(linearizationOptions, self.linearOptions, "Linearization-option", None)

    def setOptimizationOptions (self, optimizationOptions):  # 17
        """
        This method is used to set optimization options. It can be called:
        with a sequence of optimization options name and assigning corresponding values as arguments as show in the example below:
        usage
        >>> setOptimizationOptions("Name=value")
        >>> setOptimizationOptions(["Name1=value1","Name2=value2"])
        """
        return self.setMethodHelper(optimizationOptions, self.optimizeOptions, "optimization-option", None)

    def setInputs (self, name):  # 15
        """
        This method is used to set input values. It can be called:
        with a sequence of input name and assigning corresponding values as arguments as show in the example below:
        usage
        >>> setInputs("Name=value")
        >>> setInputs(["Name1=value1","Name2=value2"])
        """
        if (isinstance(name, str)):
            name = self.strip_space(name)
            value = name.split("=")
            if value[0] in self.inputlist:
                tmpvalue = eval(value[1])
                if (isinstance(tmpvalue, int) or isinstance(tmpvalue, float)):
                    self.inputlist[value[0]] = [(float(self.simulateOptions["startTime"]), float(value[1])),
                                                (float(self.simulateOptions["stopTime"]), float(value[1]))]
                elif (isinstance(tmpvalue, list)):
                    self.checkValidInputs(tmpvalue)
                    self.inputlist[value[0]] = tmpvalue
                self.inputFlag = True
            else:
                print(value[0], "!is not an input")
        elif (isinstance(name, list)):
            name = self.strip_space(name)
            for var in name:
                value = var.split("=")
                if value[0] in self.inputlist:
                    tmpvalue = eval(value[1])
                    if (isinstance(tmpvalue, int) or isinstance(tmpvalue, float)):
                        self.inputlist[value[0]] = [(float(self.simulateOptions["startTime"]), float(value[1])),
                                                    (float(self.simulateOptions["stopTime"]), float(value[1]))]
                    elif (isinstance(tmpvalue, list)):
                        self.checkValidInputs(tmpvalue)
                        self.inputlist[value[0]] = tmpvalue
                    self.inputFlag = True
                else:
                    print(value[0], "!is not an input")

    def checkValidInputs (self, name):
        if name != sorted(name, key=lambda x: x[0]):
            print('Time value should be in increasing order')
            return
        for l in name:
            if isinstance(l, tuple):
                # if l[0] < float(self.simValuesList[0]):
                if l[0] < float(self.simulateOptions["startTime"]):
                    print('Input time value is less than simulation startTime')
                    return
                if len(l) != 2:
                    print('Value for ' + l + ' is in incorrect format!')
                    return
            else:
                print('Error!!! Value must be in tuple format')
                return

    # To create csv file for inputs
    def __simInput (self):
        sl = list()  # Actual timestamps
        skip = False
        # inp = list()
        # inp = deepcopy(self.__getInputValues())
        inp = deepcopy(list(self.inputlist.values()))
        for i in inp:
            cl = list()
            el = list()
            for (t, x) in i:
                cl.append(t)
            for i in cl:
                if skip is True:
                    skip = False
                    continue
                if i not in sl:
                    el.append(i)
                else:
                    elem_no = cl.count(i)
                    sl_no = sl.count(i)
                    if elem_no == 2 and sl_no == 1:
                        el.append(i)
                        skip = True
            sl = sl + el

        sl.sort()
        for t in sl:
            for i in inp:
                for ttt in [tt[0] for tt in i]:
                    if t not in [tt[0] for tt in i]:
                        i.append((t, '?'))
        inpSortedList = list()
        sortedList = list()
        for i in inp:
            sortedList = sorted(i, key=lambda x: x[0])
            inpSortedList.append(sortedList)
        for i in inpSortedList:
            ind = 0
            for (t, x) in i:
                if x == '?':
                    t1 = i[ind - 1][0]
                    u1 = i[ind - 1][1]
                    t2 = i[ind + 1][0]
                    u2 = i[ind + 1][1]
                    nex = 2
                    while (u2 == '?'):
                        u2 = i[ind + nex][1]
                        t2 = i[ind + nex][0]
                        nex += 1
                    x = float(u1 + (u2 - u1) * (t - t1) / (t2 - t1))
                    i[ind] = (t, x)
                ind += 1
        slSet = list()
        slSet = set(sl)
        for i in inpSortedList:
            tempTime = list()
            for (t, x) in i:
                tempTime.append(t)
            inSl = None
            inI = None
            for s in slSet:
                inSl = sl.count(s)
                inI = tempTime.count(s)
                if inSl != inI:
                    test = list()
                    test = [(x, y) for x, y in i if x == s]
                    i.append(test[0])
        newInpList = list()
        tempSorting = list()
        for i in inpSortedList:
            # i.sort() => just sorting might not work so need to sort according to 1st element of a tuple
            tempSorting = sorted(i, key=lambda x: x[0])
            newInpList.append(tempSorting)

        interpolated_inputs_all = list()
        for i in newInpList:
            templist = list()
            for (t, x) in i:
                templist.append(x)
            interpolated_inputs_all.append(templist)

        name_ = 'time'
        # name = ','.join(self.__getInputNames())
        name = ','.join(list(self.inputlist.keys()))
        name = '{},{},{}'.format(name_, name, 'end')

        a = ''
        l = []
        l.append(name)
        for i in range(0, len(sl)):
            a = ("%s,%s" % (
                str(float(sl[i])), ",".join(list(str(float(inppp[i])) for inppp in interpolated_inputs_all)))) + ',0'
            l.append(a)

        self.csvFile = '{}.csv'.format(self.modelName)
        with open(self.csvFile, "w") as f:
            writer = csv.writer(f, delimiter='\n')
            writer.writerow(l)

    # to convert Modelica model to FMU
    def convertMo2Fmu (self, version="2.0", fmuType="me_cs", fileNamePrefix="<default>", includeResources=True):  # 19
        """
        This method is used to generate FMU from the given Modelica model. It creates "modelName.fmu" in the current working directory. It can be called:
        with no arguments
        with arguments of https://build.openmodelica.org/Documentation/OpenModelica.Scripting.translateModelFMU.html
        usage
        >>> convertMo2Fmu()
        >>> convertMo2Fmu(version="2.0", fmuType="me|cs|me_cs", fileNamePrefix="<default>", includeResources=true)
        """
        convertMo2FmuError = ''
        if fileNamePrefix == "<default":
            fileNamePrefix = self.modelName
        if includeResources:
            includeResourcesStr = "true"
        else:
            includeResourcesStr = "false"
        properties = 'version="{}", fmuType="{}", fileNamePrefix="{}", includeResources={}'.format(version, fmuType,
                                                                                                   fileNamePrefix,
                                                                                                   includeResourcesStr)
        translateModelFMUResult = self.requestApi('translateModelFMU', self.modelName, properties)
        if convertMo2FmuError:
            print(convertMo2FmuError)

        return translateModelFMUResult

    # to convert FMU to Modelica model
    def convertFmu2Mo (self, fmuName):  # 20
        """
        In order to load FMU, at first it needs to be translated into Modelica model. This method is used to generate Modelica model from the given FMU. It generates "fmuName_me_FMU.mo".
        Currently, it only supports Model Exchange conversion.
        usage
        >>> convertFmu2Mo("c:/BouncingBall.Fmu")
        """
        convertFmu2MoError = ''
        importResult = self.requestApi('importFMU', fmuName)
        convertFmu2MoError = self.requestApi('getErrorString')
        if convertFmu2MoError:
            print(convertFmu2MoError)

        return importResult

    # to optimize model
    def optimize (self):  # 21
        """
        This method optimizes model according to the optimized options. It can be called:
        only without any arguments
        usage
        >>> optimize()
        """
        cName = self.modelName
        properties = ','.join("%s=%s" % (key, val) for (key, val) in list(self.optimizeOptions.items()))
        optimizeError = ''
        self.getconn.sendExpression("setCommandLineOptions(\"-g=Optimica\")")
        optimizeResult = self.requestApi('optimize', cName, properties)
        optimizeError = self.requestApi('getErrorString')
        if optimizeError:
            print(optimizeError)

        return optimizeResult

    # to linearize model
    def linearize (self):  # 22
        """
        This method linearizes model according to the linearized options. This will generate a linear model that consists of matrices A, B, C and D.  It can be called:
        only without any arguments
        usage
        >>> linearize()
        """
        try:
            self.getconn.sendExpression("setCommandLineOptions(\"+generateSymbolicLinearization\")")
            properties = ','.join("%s=%s" % (key, val) for (key, val) in list(self.linearOptions.items()))
            if (self.overridevariables):
                values = ','.join("%s=%s" % (key, val) for (key, val) in list(self.overridevariables.items()))
                override = "-override=" + values
            else:
                override = ""

            if self.inputFlag:
                nameVal = self.getInputs()
                for n in nameVal:
                    tupleList = nameVal.get(n)
                    for l in tupleList:
                        if l[0] < float(self.simulateOptions["startTime"]):
                            print('Input time value is less than simulation startTime')
                            return
                self.__simInput()
                csvinput = "-csvInput=" + self.csvFile
            else:
                csvinput = ""

            # linexpr="linearize(" + self.modelName + "," + properties + ", simflags=\" " + csvinput + " " + override + " \")"
            self.getconn.sendExpression(
                    "linearize(" + self.modelName + "," + properties + ", simflags=\" " + csvinput + " " + override + " \")")
            linearizeError = ''
            linearizeError = self.requestApi('getErrorString')
            if linearizeError:
                print(linearizeError)
                return

            # code to get the matrix and linear inputs, outputs and states
            getLinFile = '{}_{}.{}'.format('linear', self.modelName, 'mo')
            checkLinFile = os.path.exists(getLinFile)
            if checkLinFile:
                self.requestApi('loadFile', getLinFile)
                cNames = self.requestApi('getClassNames')
                linModelName = cNames[0]
                buildModelmsg = self.requestApi('buildModel', linModelName)
                self.xmlFile = os.path.join(os.path.dirname(buildModelmsg[0]), buildModelmsg[1]).replace("\\", "/")
                if (os.path.exists(self.xmlFile)):
                    self.linearizationFlag = True
                    self.linearparameters = {}
                    self.linearquantitiesList = []
                    self.linearinputs = []
                    self.linearoutputs = []
                    self.linearstates = []
                    self.xmlparse()
                    matrices = self.getlinearMatrix()
                    return matrices
                else:
                    return self.requestApi('getErrorString')
        except Exception as e:
            raise e

    def getLinearInputs (self):
        """
        function which returns the LinearInputs after Linearization is performed
        usage
        >>> getLinearInputs()
        """
        return self.linearinputs

    def getLinearOutputs (self):
        """
        function which returns the LinearInputs after Linearization is performed
        usage
        >>> getLinearOutputs()
        """
        return self.linearoutputs

    def getLinearStates (self):
        """
        function which returns the LinearInputs after Linearization is performed
        usage
        >>> getLinearStates()
        """
        return self.linearstates

    def getlinearMatrix (self):
        """
        Helper Function which generates the Linear Matrix A,B,C,D
        """
        matrix_A = OrderedDict()
        matrix_B = OrderedDict()
        matrix_C = OrderedDict()
        matrix_D = OrderedDict()
        for i in self.linearparameters:
            name = i
            if (name[0] == "A"):
                matrix_A[name] = self.linearparameters[i]
            if (name[0] == "B"):
                matrix_B[name] = self.linearparameters[i]
            if (name[0] == "C"):
                matrix_C[name] = self.linearparameters[i]
            if (name[0] == "D"):
                matrix_D[name] = self.linearparameters[i]

        tmpmatrix_A = self.getLinearMatrixValues(matrix_A)
        tmpmatrix_B = self.getLinearMatrixValues(matrix_B)
        tmpmatrix_C = self.getLinearMatrixValues(matrix_C)
        tmpmatrix_D = self.getLinearMatrixValues(matrix_D)

        return [tmpmatrix_A, tmpmatrix_B, tmpmatrix_C, tmpmatrix_D]

    def getLinearMatrixValues (self, matrix):
        """
        Helper Function which generates the Linear Matrix A,B,C,D
        """
        if (matrix):
            x = list(matrix.keys())
            name = x[-1]
            tmpmatrix = np.zeros((int(name[2]), int(name[4])))
            for i in x:
                rows = int(i[2]) - 1
                cols = int(i[4]) - 1
                tmpmatrix[rows][cols] = matrix[i]
            return tmpmatrix
        else:
            return np.zeros((0, 0))
