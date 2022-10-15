FROM golang:alpine3.16 AS builder
WORKDIR /usr/src/app
ENV GOPROXY=https://goproxy.cn

#获取时区配置、证书配置
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
  apk add --no-cache upx ca-certificates tzdata

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY ./main.go ./

# -ldflags "-s -w" ：去除构建包的调式信息
# CGO_ENABLED=0 ：关闭CGO ,go build 默认使用CGO
# 压缩编译后的httpserver
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o server &&\
  upx --best server -o _upx_server && \
  mv -f _upx_server server

FROM alpine AS runner
#时区
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
#证书
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#httpserver
COPY --from=builder /usr/src/app/server /opt/app/
EXPOSE 8999
CMD ["/opt/app/server"]