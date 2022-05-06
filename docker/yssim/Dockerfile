FROM openmodelica/openmodelica:v1.18.0-minimal
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




WORKDIR /home/simtek/code/docker/yssim
COPY Miniconda3.sh ./
COPY pipList.txt ./
COPY start.sh/ /home/simtek/docker/start.sh

RUN ./Miniconda3.sh -p ../../miniconda3 -f -b
RUN ln -f ../../miniconda3/bin/python3.9 /usr/bin/python \
    && ln -f ../../miniconda3/bin/pip /usr/bin/pip
RUN  pip install -r pipList.txt -i https://pypi.tuna.tsinghua.edu.cn/simple
RUN ln -f ../../miniconda3/bin/uvicorn /usr/bin/uvicorn

USER simtek
#CMD ["/home/simtek/code/miniconda3/bin/supervisord", "-n", "-c", "/home/simtek/code/docker/yssim/supervisor_OpenModelica.ini"]
