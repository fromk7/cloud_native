# 模块三作业说明

## 目录说明
```
module3 
 ---Dockerfile 即构建容器的文件 
 ---main.go httpserver程序代码
 ---go.mod、go.sum 依赖
```
## 构建httpserver容器镜像
```
在Dockerfile文件所在的目录，执行命令 "docker build -t myhttpserver:1.0 ."
```
## 运行httpserver容器镜像
```
docker run myhttpserver:1.0
```
## 访问httpserver
```
访问地址：http://localhost:8999/healthz
正常页面响应结果：页面显示 200
```
