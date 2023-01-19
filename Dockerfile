#1.build image 打包成镜像  docker build -f Dockerfile -t jonnymark/jonny-gin:v1.3 .
# 2. start: docker run --rm -it -p 801:801 jonnymark/jonny-gin:v1.3
# 3. test: curl -i http://localhost:801/health

#多阶段构建,实现镜像瘦身 原理：我不需要一整个编译的工具链，我只需要一个编译好的可执行文件
# 参考链接 https://www.orchome.com/8191

# 第一阶段 生成可执行文件阶段
# Compile stage
FROM golang:1.19-alpine AS builder

# 安装grpc_health工具
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.8 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

# 镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai

# 工作目录
WORKDIR /Users/jonny/project/go/src/gin-micro-mine

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY . .

# 编译成二进制文件
RUN go build main.go

## 安装 一些基础包
RUN apk add --no-cache  git  make  bash  ca-certificates  \
    # wget  tcpdump  iputils iproute2 libc6-compat \
    && apk add -U tzdata \
    && rm -rf /var/cache/apk/*

# 第二阶段 执行可执行文件的阶段，用到了第一阶段的执行文件和基础配置，执行过程中不再依赖golang基础镜像
# 创建一个小镜像
# Final stage
#FROM debian:stretch-slim
FROM alpine:latest

# 工作目录
WORKDIR /data/app

# 从builder镜像中把 /build 拷贝到当前目录
COPY --from=builder /Users/jonny/project/go/src/gin-micro-mine/main .
COPY --from=builder /Users/jonny/project/go/src/gin-micro-mine/config  ./config
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe

EXPOSE 801
EXPOSE 901

# 需要运行的命令
CMD [ "./main", "-c", "config", "-e" ,"local"]