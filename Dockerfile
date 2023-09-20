FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
GOPROXY=https://goproxy.cn,direct \
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 拷贝项目文件
COPY . .

# 编译代码
RUN go mod tidy && \
    go build main.go

###################
# 接下来创建一个小镜像
###################
FROM alpine
WORKDIR /app

COPY . .
# 从builder镜像中把可执行文件拷贝到当前目录
COPY --from=builder /build/main /app

# 设置时区
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk --no-cache add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates


ENTRYPOINT ["./main"]