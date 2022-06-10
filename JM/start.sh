#!/bin/sh

if test "${JAVA_HOME}" = ""; then
  export JAVA_HOME="$(java -XshowSettings:properties -version 2>&1    | sed '/^[[:space:]]*java\.home/!d;s/^[[:space:]]*java\.home[[:space:]]*=[[:space:]]*//')"
fi

#Create path to file
CURRENT_DIR="$(pwd)"
cd "$(dirname "$0")"
FILE_DIR="$(pwd)"
#Go back to previous location
cd ${CURRENT_DIR}

#Set variables using relative paths
JMODELICA_HOME=${FILE_DIR}/.. \
IPOPT_HOME=/Ipopt \
SUNDIALS_HOME=${JMODELICA_HOME}/ThirdParty/Sundials \
PYTHONPATH=:${JMODELICA_HOME}/Python/::$PYTHONPATH \
LD_LIBRARY_PATH=:/Ipopt/lib/:${JMODELICA_HOME}/ThirdParty/Sundials/lib:${JMODELICA_HOME}/ThirdParty/CasADi/lib:$LD_LIBRARY_PATH \
SEPARATE_PROCESS_JVM=${JAVA_HOME} \
python /JM/main.py $@
