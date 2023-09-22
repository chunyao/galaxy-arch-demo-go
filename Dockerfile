# 使用官方的 Golang 镜像作为构建环境
FROM golang:1.19.0-alpine3.16  AS builder

ENV TZ=Asia/Shanghai \
    LANG=C.UTF-8 \
    LANGUAGE=C.UTF-8
# 设置工作目录
WORKDIR /app

# 将代码加入镜像中
COPY . .
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn
# 构建二进制文件
RUN go env -w GO111MODULE=on && \
    GOOS=linux GOARCH=amd64 go build -o main main.go

# 使用最小化的 alpine 镜像作为运行环境
FROM mabangerp-docker.pkg.coding.net/public/php-project/golang:1.1

# 设置工作目录
ARG SERVICENAME=mas-arch-demo-go
ENV profile=$ACTIVE_PROFILE
ENV NODE_OWN_IP=1
ENV APP_HOST_PORT=1

WORKDIR /data/web/website/$SERVICENAME
EXPOSE 8080
EXPOSE 8081
# 从之前构建的镜像中将二进制文件复制到 alpine 镜像中
COPY --from=builder /app/resources ./resources
COPY --from=builder /app/run.sh .
COPY --from=builder /app/main .
USER root
# 运行二进制文件
CMD ./run.sh $ACTIVE_PROFILE