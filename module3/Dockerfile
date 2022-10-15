FROM golang:latest
MAINTAINER "fromk7@hotmail.com"
LABEL descriptio="myhttpserver"
WORKDIR /myapp
RUN apt-get update && apt-get install  && apt-get -y curl
RUN curl  https://github.com/fromk7/cloud_native/blob/main/module2/main.go >> myhttpserver.go
RUN go build myhttpserver.go
EXPOSE 80
ENTRYPOINT ["./myhttpserver"]