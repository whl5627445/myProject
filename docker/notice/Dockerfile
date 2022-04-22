FROM ubuntu:20.04
RUN useradd --create-home --no-log-init --shell /bin/bash simtek \
    && adduser simtek sudo \
    && echo 'simtek:simtek' | chpasswd
RUN apt-get update -y \
    && apt-get upgrade -y \
    && apt-get install sudo -y
ENV TZ=Asia/Shanghai
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list \
          && apt-get update \
          && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone \
          && apt-get install tzdata \
          && apt-get clean \
          && apt-get autoclean \
          && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*



WORKDIR /code
#RUN  ./miniconda3/bin/pip install -r pipList.txt -i https://pypi.tuna.tsinghua.edu.cn/simple


USER simtek
#CMD ["/usr/bin/omc", "--interactive=zmq", "--locale=C", "-z=simtek_omc", "--interactivePort=23456"]
