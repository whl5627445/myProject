#FROM golang:alpine AS golang-build
#WORKDIR /go-service
#COPY ../../ /go-service
#ENV GOPROXY https://goproxy.cn,direct
#RUN  CGO_ENABLED=0 GOOS=linux go build main.go


#FROM openmodelica/openmodelica:v1.18.0-minimal AS app
FROM openmodelica/openmodelica:v1.19.2-minimal
#FROM openmodelica/openmodelica:v1.18.0-minimal
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


WORKDIR /home/simtek
COPY ./start.sh start.sh
USER simtek

#CMD ["supervisord", "-n", "-c", "/home/simtek/code/docker/yssim/supervisor_OpenModelica.ini"]
