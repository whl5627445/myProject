感谢https://www.jetbrains.com/ 的赞助
# 介绍：
### 这是一个基于OpenModelica项目后端omc的Web API项目
# 项目目标与愿景：
### 本项目的目标是为有需要的用户提供一个基于omc的后端API，将Modelica模型的图形显示、模型操作、执行仿真与具体的Modelica代码解析分离，简单来说就是基于omc进行功能上的增强与简化，方便其他开发者更容易的使用

# 功能
#### 1.模型编译
#### 2.模型仿真
#### 3.模型源代码解析
#### 4.模型实例化
#### 5.更多功能还在开发中

# 依赖：
#### omc官方的安装，最低版本1.21
#### Golang最低版本1.21
#### 此项目并无图形显示组件，需要使用者自行开发图形渲染代码
#### MySQL数据库，用于存储用户信息，模型信息等
#### 启动项目非常简单，设置好config里的各项配置后进行编译，例如go build main.go

# 安装指南：
    1.安装omc： 官方地址：https://openmodelica.org/download/download-linux/
    2.安装Golang： 官方地址：https://go.dev/doc/install
    3.git clone本项目
    4.配置config文件中的数据库地址
    5.编译项目

# 部署指南：
    1.如果你已经安装安装指南配置安装完成，那么你只需要将编译好的二进制文件运行即可
    2.目前建议在linux服务器上运行