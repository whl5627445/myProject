FROM openmodelica/openmodelica:v1.20.0-minimal
RUN useradd --create-home --no-log-init --shell /bin/bash simtek \
    && adduser simtek sudo \
    && echo 'simtek:simtek' | chpasswd

RUN apt-get update -y \
    && apt-get upgrade -y \
    && apt-get install sudo -y \
    && apt-get install supervisor -y \
#    && apt-get install libzmq3-dev -y \
    && apt-get install rar -y\
    && apt-get install unrar -y

ENV TZ=Asia/Shanghai
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list \
          && apt-get update \
          && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone \
          && apt-get install tzdata \
          && apt-get clean \
          && apt-get autoclean \
          && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


COPY requirements.txt /yssim-go/grpc/PythonGrpc/
COPY start.sh /yssim-go/grpc/PythonGrpc/
WORKDIR /yssim-go/grpc/PythonGrpc
RUN sudo apt-get update -y
RUN sudo apt-get install python3-pip -y
RUN pip3 config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple && \
	pip3 install -r requirements.txt
USER simtek