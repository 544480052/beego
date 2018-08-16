#声明是基于centos7镜像搭建的golang开发环境
FROM centos:7
#声明镜像的作者
MAINTAINER lucky pingzi <chengx@weipaitang.com>
#安装go可能依赖的编译器
RUN yum install -y gcc
#安装go
RUN yum install -y go
#配置go环境变量
ENV GOROOT /usr/local/go
ENV PATH=$PATH:/usr/local/go/bin

RUN mkdir -p /data/www/go
RUN mkdir -p /data/www/go/src
RUN mkdir -p /data/www/go/pkg
RUN mkdir -p /data/www/go/bin
ENV GOPATH /data/www/go
#拷贝golang源码
RUN mkdir -p /data/www/go/src/server
COPY /data/www/go/beego/* /data/www/go/src/server/
#编译web服务器
WORKDIR /data/www/go/src/server
RUN go build -o server.bin main.go
#指定容器的入口
CMD /data/www/go/src/server/server.bin


//////////////////////////////////////////////////////////////
#生成镜像
docker build -t server:v0.1 .
#启动容器
docker run -d -p 9090:9090 server:v0.1







