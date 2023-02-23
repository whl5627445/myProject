FROM openmodelica/openmodelica:v1.20.0-minimal
RUN useradd --create-home --no-log-init --shell /bin/bash simtek \
    && adduser simtek sudo \
    && echo 'simtek:simtek' | chpasswd

RUN apt-get update -y \
    && apt-get upgrade -y \
    && apt-get install sudo -y \
    && apt-get install supervisor -y \
    && apt-get install libzmq3-dev -y \
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


COPY requirements.txt /home/simtek/
COPY start.sh /home/simtek/
WORKDIR /home/simtek
RUN sudo chmod +x start.sh
RUN sudo apt-get update -y
RUN sudo apt-get install python3-pip -y
RUN pip3 config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple && \
	pip3 install -r requirements.txt
USER simtek

#上面的代码如果不适用
#可以用备选方案，先用yssim管理系统创建一个镜像，然后在ini文件中只开启omc的进程，让容器跑起来，然后再用docker commit
#生成新的镜像，修改新的镜像中的start.sh文件，和supervisor_GRPC.ini。