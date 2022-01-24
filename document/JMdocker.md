# JModelica镜像使用
## 拉取镜像：
    镜像已经是打包好的运行环境，如果需要修改代码，修改容器启动时的脚本， 请查看容器内/start.sh文件， 守护进程是supervisor，
    docker pull mihoutao/jmimages:v2

## 创建容器：
### 创建容器时按照下面的对应修改一下即可
    /lib/omlibrary：OpenModelica的标准库路径，
    /home/simtek/dev/public/：项目代码当中的静态文件目录， 存放用户上传的.om、.fmu后缀等文件的目录，
    /home/simtek/JM_ubuntu：是JModelica的编译安装目录, 在本项目代码中已经有一份编译好的JM代码，但是需要docker运行环境
    2204：ssh登录端口
    56789：内部服务端口

    sudo docker run -itd  -p 56789:56789 -v /home/simtek/dev/omlibrary/:/omlibrary -v /home/simtek/dev/public/:/public -v /home/simtek/dev/JM:/JM --name jm-v2.2 --restart always mihoutao/jmodelica:v2.2 /start.sh




# openmodelica镜像使用
    sudo docker build -t omc-v1 .
## docker文件夹路径下的dockerfile可直接构建， 下面命令直接执行运行容器即可，--name是容器名称， -p是python服务对外端口
    sudo docker run -itd  --restart always -v /home/simtek/dev:/home/simtek/code -v /home/simtek/dev/omlibrary/:/usr/bin/../lib/omlibrary --name wanghailong --network yssim-net -p 4327:8084 omc-v1 /home/simtek/docker/start.sh
