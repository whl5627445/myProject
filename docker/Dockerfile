FROM openmodelica/openmodelica:v1.21.0-minimal
#RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone
ENV TZ=Asia/Shanghai

RUN useradd --create-home --no-log-init --shell /bin/bash simtek \
    && adduser simtek sudo \
    && echo 'simtek:simtek' | chpasswd
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list \
    && apt-get update -y \
    && apt-get upgrade -y \
    && apt-get install sudo -y \
    && apt-get install supervisor -y \
#    && apt-get install libzmq3-dev -y \
    && apt-get install rar -y\
    && apt-get install unrar -y


RUN   ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone \
      && apt-get install tzdata \
      && apt-get clean \
      && apt-get autoclean \
      && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


WORKDIR /home/simtek
COPY ./start.sh start.sh
USER root

#CMD ["supervisord", "-n", "-c", "/home/simtek/code/docker/yssim/supervisor_OpenModelica.ini"]
