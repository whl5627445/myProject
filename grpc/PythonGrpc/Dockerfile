FROM openmodelica/openmodelica:v1.20.0-minimal
RUN apt-get update -y
RUN apt-get install python3-pip -y
COPY requirements.txt /yssim-go/grpc/PythonGrpc/
WORKDIR /yssim-go/grpc/PythonGrpc
RUN pip3 config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple && \
	pip3 install -r requirements.txt \
RUN useradd --create-home --no-log-init --shell /bin/bash simtek \
    && adduser simtek sudo \
    && echo 'simtek:simtek' | chpasswd \
RUN su simtek
RUN /usr/bin/omc --interactive=zmq --locale=C -z=simtek_omc --interactivePort=23456
CMD python3 app.py

