
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
    docker run -itd -p 2204:22 -p 56789:56789 -v /lib/omlibrary/:/omlibrary -v /home/simtek/dev/public/:/public -v /home/simtek/JM_ubuntu:/JM --name jm-v2 --restart always jmimages:v2 /start.sh

