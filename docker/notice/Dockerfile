FROM python:3.9.12
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



WORKDIR /home/simtek/code/docker/notice
COPY ./ /home/simtek/code/docker/notice
USER simtek
RUN pip install -r pipList.txt -i https://pypi.tuna.tsinghua.edu.cn/simple
USER root
RUN ln -f /home/simtek/.local/bin/uvicorn /usr/bin/uvicorn

USER simtek

