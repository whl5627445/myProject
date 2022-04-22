# -*- coding: utf-8 -*-

from __future__ import absolute_import
from __future__ import division
from __future__ import print_function

import abc
import getpass
import json
import logging
import os
import shlex
import signal
import subprocess
import sys
import tempfile
import time
import uuid
from builtins import int, range
from distutils import spawn
import psutil
import pyparsing
from future.utils import with_metaclass

if sys.platform == 'darwin':
    # On Mac let's assume omc is installed here and there might be a broken omniORB installed in a bad place
    sys.path.append('/opt/local/lib/python2.7/site-packages/')
    sys.path.append('/opt/openmodelica/lib/python2.7/site-packages/')


# Logger Defined
logger = logging.getLogger('OMPython')
logger.setLevel(logging.DEBUG)
# create console handler with a higher log level
logger_console_handler = logging.StreamHandler()
logger_console_handler.setLevel(logging.INFO)

# create formatter and add it to the handlers
logger_formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
logger_console_handler.setFormatter(logger_formatter)

# add the handlers to the logger
logger.addHandler(logger_console_handler)
logger.setLevel(logging.WARNING)


class DummyPopen():
    def __init__ (self, pid):
        self.pid = pid
        self.process = psutil.Process(pid)
        self.returncode = 0

    def poll (self):
        return None if self.process.is_running() else True

    def kill (self):
        return os.kill(self.pid, signal.SIGKILL)

    def wait (self, timeout):
        return self.process.wait(timeout=timeout)


class OMCSessionHelper():
    def __init__ (self):
        # Get the path to the OMC executable, if not installed this will be None
        omc_env_home = os.environ.get('OPENMODELICAHOME')
        if omc_env_home:
            self.omhome = omc_env_home
        else:
            path_to_omc = spawn.find_executable("omc")
            if path_to_omc is None:
                raise ValueError("Cannot find OpenModelica executable, please install from openmodelica.org")
            self.omhome = os.path.split(os.path.split(os.path.realpath(path_to_omc))[0])[0]

    def _get_omc_path (self):
        try:
            return os.path.join(self.omhome, 'bin', 'omc')
        except BaseException:
            logger.error(
                    "The OpenModelica compiler is missing in the System path (%s), please install it" % os.path.join(
                            self.omhome, 'bin', 'omc'))
            raise


class OMCSessionBase(with_metaclass(abc.ABCMeta, object)):
    def __init__ (self, readonly=False, interactivePort = None, random_string=None, name="simtek"):
        self.readonly = readonly
        self.omc_cache = {}
        self._omc_process = None
        self._omc_command = None
        self._omc = None
        self._dockerCid = None
        self._serverIPAddress = "127.0.0.1"
        self._interactivePort = interactivePort
        self._name = name
        # FIXME: this code is not well written... need to be refactored
        self._temp_dir = tempfile.gettempdir()
        # generate a random string for this session
        self._random_string = random_string
        if not random_string:
            self._random_string = uuid.uuid4().hex
        # omc log file
        self._omc_log_file = None
        try:
            self._currentUser = getpass.getuser()
            if not self._currentUser:
                self._currentUser = "nobody"
        except KeyError:
            # We are running as a uid not existing in the password database... Pretend we are nobody
            self._currentUser = "nobody"

    def __del__ (self):
        try:
            self.sendExpression("quit()")
        except:
            pass
        self._omc_log_file.close()
        if sys.version_info.major >= 3:
            try:
                self._omc_process.wait(timeout=2.0)
            except:
                if self._omc_process:
                    self._omc_process.kill()
        else:
            for i in range(0, 100):
                time.sleep(0.02)
                if self._omc_process and (self._omc_process.poll() is not None):
                    break
        # kill self._omc_process process if it is still running/exists
        if self._omc_process is not None and self._omc_process.returncode is None:
            print("OMC did not exit after being sent the quit() command; killing the process with pid=%s" % str(
                    self._omc_process.pid))
            if sys.platform == "win32":
                self._omc_process.kill()
                self._omc_process.wait()
            else:
                os.killpg(os.getpgid(self._omc_process.pid), signal.SIGTERM)
                self._omc_process.kill()
                self._omc_process.wait()

    def _create_omc_log_file (self, suffix):
        if sys.platform == 'win32':
            self._omc_log_file = open(
                    os.path.join(self._temp_dir, "openmodelica.{0}.{1}.log".format(suffix, self._random_string)), 'w')
        else:
            # this file must be closed in the destructor
            self._omc_log_file = open(os.path.join(self._temp_dir,
                                                   "openmodelica.{0}.{1}.{2}.log".format(self._currentUser, suffix,
                                                                                         self._random_string)), 'w')

    def _start_omc_process (self, timeout, sys_start=True):
        if sys.platform == 'win32':
            omhome_bin = os.path.join(self.omhome, 'bin').replace("\\", "/")
            my_env = os.environ.copy()
            my_env["PATH"] = omhome_bin + os.pathsep + my_env["PATH"]
            self._omc_process = subprocess.Popen(self._omc_command, stdout=self._omc_log_file,
                                                 stderr=self._omc_log_file, env=my_env)
        else:
            if sys_start:
                for i in psutil.process_iter():
                    if i.name() == "omc":
                        i.kill()
                        time.sleep(0.5)
                        break
            # Because we spawned a shell, and we need to be able to kill OMC, create a new process group for this
            self._omc_process = subprocess.Popen(self._omc_command, shell=True, stdout=self._omc_log_file,
                                                 stderr=self._omc_log_file, preexec_fn=os.setsid)
        if self._docker:
            for i in range(0, 40):
                try:
                    with open(self._dockerCidFile, "r") as fin:
                        self._dockerCid = fin.read().strip()
                except:
                    pass
                if self._dockerCid:
                    break
                time.sleep(timeout / 40.0)
            try:
                os.remove(self._dockerCidFile)
            except:
                pass
            if self._dockerCid is None:
                logger.error("Docker did not start. Log-file says:\n%s" % (open(self._omc_log_file.name).read()))
                raise Exception(
                        "Docker did not start (timeout=%f might be too short especially if you did not docker pull the image before this command)." % timeout)
        if self._docker or self._dockerContainer:
            if self._dockerNetwork == "separate":
                self._serverIPAddress = \
                    json.loads(subprocess.check_output(["docker", "inspect", self._dockerCid]).decode().strip())[0][
                        "NetworkSettings"]["IPAddress"]
            for i in range(0, 40):
                if sys.platform == 'win32':
                    break
                dockerTop = subprocess.check_output(["sudo", "docker", "top", self._dockerCid]).decode().strip()
                self._omc_process = None
                for line in dockerTop.split("\n"):
                    columns = line.split()
                    if self._random_string in line:
                        try:
                            self._omc_process = DummyPopen(int(columns[1]))
                        except psutil.NoSuchProcess:
                            raise Exception(
                                    "Could not find PID %d - is this a docker instance spawned without --pid=host?\nLog-file says:\n%s" % (
                                        self._random_string, dockerTop, open(self._omc_log_file.name).read()))
                        break
                if self._omc_process is not None:
                    break
                time.sleep(timeout / 40.0)
            if self._omc_process is None:
                raise Exception("Docker top did not contain omc process %s:\n%s\nLog-file says:\n%s" % (
                    self._random_string, dockerTop, open(self._omc_log_file.name).read()))
        return self._omc_process

    def _getuid (self):
        """
        The uid to give to docker.
        On Windows, volumes are mapped with all files are chmod ugo+rwx,
        so uid does not matter as long as it is not the root user.
        """
        return 1000 if sys.platform == 'win32' else os.getuid()

    def _set_omc_command (self, omc_path_and_args_list):
        """Define the command that will be called by the subprocess module.

        On Windows, use the list input style of the subprocess module to
        avoid problems resulting from spaces in the path string.
        Linux, however, only works with the string version.
        """
        if (self._docker or self._dockerContainer) and sys.platform == "win32":
            extraFlags = ["-d=zmqDangerousAcceptConnectionsFromAnywhere"]
            if not self._interactivePort:
                raise Exception(
                        "docker on Windows requires knowing which port to connect to. For dockerContainer=..., the container needs to have already manually exposed this port when it was started (-p 127.0.0.1:n:n) or you get an error later.")
        else:
            extraFlags = []
        if self._docker:
            if sys.platform == "win32":
                p = int(self._interactivePort)
                dockerNetworkStr = ["-p", "127.0.0.1:%d:%d" % (p, p)]
            elif self._dockerNetwork == "host" or self._dockerNetwork is None:
                dockerNetworkStr = ["--network=host"]
            elif self._dockerNetwork == "separate":
                dockerNetworkStr = []
                extraFlags = ["-d=zmqDangerousAcceptConnectionsFromAnywhere"]
            else:
                raise Exception('dockerNetwork was set to %s, but only \"host\" or \"separate\" is allowed')
            self._dockerCidFile = self._omc_log_file.name + ".docker.cid"
            omcCommand = ["sudo", "docker", "run", "--cidfile", self._dockerCidFile, "--rm", "--env",
                          "USER=%s" % self._currentUser, "--user",
                          str(self._getuid())] + self._dockerExtraArgs + dockerNetworkStr + [self._docker,
                                                                                             self._dockerOpenModelicaPath]
        elif self._dockerContainer:
            omcCommand = ["sudo", "docker", "exec", "--env", "USER=%s" % self._currentUser, "--user",
                          str(self._getuid())] + self._dockerExtraArgs + [self._dockerContainer,
                                                                          self._dockerOpenModelicaPath]
            self._dockerCid = self._dockerContainer
        else:
            omcCommand = [self._get_omc_path()]
        if self._interactivePort:
            extraFlags = extraFlags + ["--interactivePort=%d" % int(self._interactivePort)]

        omc_path_and_args_list = omcCommand + omc_path_and_args_list + extraFlags

        if sys.platform == 'win32':
            self._omc_command = omc_path_and_args_list
        else:
            self._omc_command = ' '.join(
                    [shlex.quote(a) if (sys.version_info > (3, 0)) else a for a in omc_path_and_args_list])

        return self._omc_command

    @abc.abstractmethod
    def _connect_to_omc (self, timeout):
        pass

    @abc.abstractmethod
    def execute (self, command):
        pass

    def clearOMParserResult (self):
        OMParser.result = {}

    @abc.abstractmethod
    def sendExpression (self, command, parsed=True):
        """
        Sends an expression to the OpenModelica. The return type is parsed as if the
        expression was part of the typed OpenModelica API (see ModelicaBuiltin.mo).
        * Integer and Real are returned as Python numbers
        * Strings, enumerations, and typenames are returned as Python strings
        * Arrays, tuples, and MetaModelica lists are returned as tuples
        * Records are returned as dicts (the name of the record is lost)
        * Booleans are returned as True or False
        * NONE() is returned as None
        * SOME(value) is returned as value
        """
        pass

    def ask (self, question, opt=None, parsed=True):
        p = (question, opt, parsed)

        if self.readonly and question != 'getErrorString':
            # can use cache if readonly
            if p in self.omc_cache:
                return self.omc_cache[p]

        if opt:
            expression = '{0}({1})'.format(question, opt)
        else:
            expression = question

        logger.debug('OMC ask: {0}  - parsed: {1}'.format(expression, parsed))

        try:
            if not parsed:
                res = self.execute(expression)
            else:
                res = self.sendExpression(expression)

        except Exception as e:
            logger.error("OMC failed: {0}, {1}, parsed={2}".format(question, opt, parsed))
            raise e

        # save response
        self.omc_cache[p] = res

        return res

    def loadFile (self, filename):
        cmd = 'loadFile(\"{0}\", \"UTF-8\",true,true,false)'.format(filename)
        return self.ask(cmd)

    def simulate (self, className, fileNamePrefix, simulate_parameters_data):
        cmd = className + ', fileNamePrefix = "' + fileNamePrefix  + 'result\"'
        if simulate_parameters_data:
            simulate_parameters_list = []
            for k, v in simulate_parameters_data.items():
                if v:
                    simulate_parameters_list.append(str(k) + "=" + str(v))
            cmd = cmd + ", " + ", ".join(simulate_parameters_list)
        self.directoryExists(fileNamePrefix)
        simulate_result =  self.ask('simulate', '{0}'.format(cmd))
        return simulate_result

    def buildModel (self, className, fileNamePrefix, simulate_parameters_data):
        cmd = className + ', fileNamePrefix = "' + fileNamePrefix  + 'result\"'
        if simulate_parameters_data:
            simulate_parameters_list = []
            for k, v in simulate_parameters_data.items():
                if v:
                    simulate_parameters_list.append(str(k) + "=" + str(v))
            cmd = cmd + ", " + ", ".join(simulate_parameters_list)
        self.directoryExists(fileNamePrefix)
        simulate_result =  self.ask('buildModel', '{0}'.format(cmd))
        return simulate_result

    def cd (self, newWorkingDirectory):
        return self.ask('cd', '"{0}"'.format(newWorkingDirectory))

    def getSimulationOptions (self, className):
        return self.ask('getSimulationOptions', className)

    @staticmethod
    def directoryExists (directory: str):
        path = './' + '/'.join(directory.split('/')[:-1])
        if not os.path.exists(path):
            os.makedirs(path)
        return

    def loadModel (self, className):
        return self.ask('loadModel', className)

    def isModel (self, className):
        return self.ask('isModel', className)

    def getErrorString (self):
        return self.ask('getErrorString')

    def isPackage (self, className):
        return self.ask('isPackage', className)

    def isPrimitive (self, className):
        return self.ask('isPrimitive', className)

    def isConnector (self, className):
        return self.ask('isConnector', className)

    def isRecord (self, className):
        return self.ask('isRecord', className)

    def isBlock (self, className):
        return self.ask('isBlock', className)

    def isType (self, className):
        return self.ask('isType', className)

    def isFunction (self, className):
        return self.ask('isFunction', className)

    def isClass (self, className):
        return self.ask('isClass', className)

    def isParameter (self, className):
        return self.ask('isParameter', className)

    def isConstant (self, className):
        return self.ask('isConstant', className)

    def isProtected (self, className):
        return self.ask('isProtected', className)

    def getPackages (self, className="AllLoadedClasses"):
        return self.ask('getPackages', className)

    def getClassRestriction (self, className):
        return self.ask('getClassRestriction', className)

    def typeNameStrings (self, className):
        return self.ask('typeNameStrings', className)

    def getNthComponent (self, className, comp_id):
        """ returns with (type, name, description) """
        return self.ask('getNthComponent', '{0}, {1}'.format(className, comp_id))

    def getNthComponentAnnotation (self, className, comp_id):
        return self.ask('getNthComponentAnnotation', '{0}, {1}'.format(className, comp_id))

    def getImportCount (self, className):
        return self.ask('getImportCount', className)

    def getNthImport (self, className, importNumber):
        # [Path, id, kind]
        return self.ask('getNthImport', '{0}, {1}'.format(className, importNumber))

    def getInheritanceCount (self, className):
        return self.ask('getInheritanceCount', className)

    def getNthInheritedClass (self, className, inheritanceDepth):
        return self.ask('getNthInheritedClass', '{0}, {1}'.format(className, inheritanceDepth))

    def getExtendsModifierNames (self, className, componentName):
        return self.ask('getExtendsModifierNames', '{0}, {1}'.format(className, componentName))

    def getExtendsModifierValue (self, className, extendsName, modifierName):
        try:
            return self.ask('getExtendsModifierValue', '{0}, {1}, {2}'.format(className, extendsName, modifierName))
        except pyparsing.ParseException as ex:
            logger.warning('OMTypedParser error: {0}'.format(ex.message))
            result = self.ask('getExtendsModifierValue', '{0}, {1}, {2}'.format(className, extendsName, modifierName),
                              parsed=False)
            try:
                return result[2:]
            except (TypeError, UnboundLocalError) as ex:
                logger.warning('OMParser error: {0}'.format(ex))
                return result

    def getNthComponentModification (self, className, comp_id):
        # get {$Code(....)} field
        # \{\$Code\((\S*\s*)*\)\}
        value = self.ask('getNthComponentModification', '{0}, {1}'.format(className, comp_id), parsed=False)
        value = value.replace("{$Code(", "")
        return value[:-3]
        # return self.re_Code.findall(value)

    # function getClassNames
    #   input TypeName class_ = $Code(AllLoadedClasses);
    #   input Boolean recursive = false;
    #   input Boolean qualified = false;
    #   input Boolean sort = false;
    #   input Boolean builtin = false "List also builtin classes if true";
    #   input Boolean showProtected = false "List also protected classes if true";
    #   output TypeName classNames[:];
    # end getClassNames;
    def getClassNames (self, className=None, recursive=False, qualified=False, sort=False, builtin=False,
                       showProtected=False):
        if className:
            value = self.ask('getClassNames',
                             '{0}, recursive={1}, qualified={2}, sort={3}, builtin={4}, showProtected={5}'.format(
                                     className, str(recursive).lower(), str(qualified).lower(), str(sort).lower(),
                                     str(builtin).lower(), str(showProtected).lower()))
        else:
            value = self.ask('getClassNames',
                             'recursive={1}, qualified={2}, sort={3}, builtin={4}, showProtected={5}'.format(className,
                                                                                                             str(recursive).lower(),
                                                                                                             str(qualified).lower(),
                                                                                                             str(sort).lower(),
                                                                                                             str(builtin).lower(),
                                                                                                             str(showProtected).lower()))
        return value






