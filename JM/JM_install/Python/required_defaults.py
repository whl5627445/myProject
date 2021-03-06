# -*- coding: utf-8 -*-

#    Copyright (C) 2016 Modelon AB
#
#    This program is free software: you can redistribute it and/or modify
#    it under the terms of the GNU General Public License as published by
#    the Free Software Foundation, version 3 of the License.
#
#    This program is distributed in the hope that it will be useful,
#    but WITHOUT ANY WARRANTY; without even the implied warranty of
#    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#    GNU General Public License for more details.
#
#    You should have received a copy of the GNU General Public License
#    along with this program.  If not, see <http://www.gnu.org/licenses/>.

"""Used by the default startup script for jmodelica.org

This module will keep trac on which environment variables
that are required to exist and point to a valid folder in
the distribution. This file is python file is generated 
during configuration.

This module and the user startup script are executed in the jmodelica
module global namespace. Therefore, please use a leading underscore on
all names that are used, unless it is explicitly desired that they are
available in this namespace (unlikely).
"""

def get_required_paths_dict():
    """
    Returns a dictionary with values that are booleans
    indicting if the environment varibles that are the keys
    are required. These are the keys in the dictionary:
    'IPOPT_HOME'
    'SUNDIALS_HOME'
    'MINGW_HOME'
    'COMPILER_JARS'
    'BEAVER_LIB'
    'CLASSPATH'
    'JYPE_JVM'
    'JVM_ARGS'
    'MINGW_HOME'
    """
    return {'IPOPT_HOME'    : True,
            'SUNDIALS_HOME' : True,
            'COMPILER_JARS' : True,
            'BEAVER_PATH'   : True,
            'MODELICAPATH'  : True,
            'JPYPE_JVM'     : True,
            'JVM_ARGS'      : False,
            'MINGW_HOME'    : True}
            
def optimica_compiler_included():
    return True

