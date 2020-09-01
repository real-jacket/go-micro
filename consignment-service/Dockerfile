# 我们使用Golang的官方镜像，它包含了所有用于构建Golang应用的依赖和工具
# 请注意`AS builder`，这命名了我们这个镜像，以后可以用`builder`来指向它
FROM golang:1.13 AS builder

# 设置 go module 代理
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

# 将工作目录设置为当前微服务在gopath中的完整路径
WORKDIR /microke/consignment-service


# 将代码复制到工作目录中
COPY . .

# 下载相关包
RUN go mod download

# 编译我们的二进制文件
RUN GOOS=linux GOARCH=amd64 go build


# 在这里，我们第二次使用了FROM，它告诉Docker，我们从这里
# 要开始第二个构建流程
FROM alpine:latest
# 确保安全性的包
# RUN apk --no-cache add ca-certificates
RUN mkdir /app

WORKDIR /app
# 在这里，我们不再直接从我们的宿主机器中直接拷贝二进制文件,
# 而是从一个叫做`builder`的容器中获取。它会从我们起先构建的镜像中
# 获得已经编译好的文件并引入到这个容器里。
COPY --from=builder /microke/consignment-service/consignment-service .
# 运行二进制文件。不同的是，这个二进制文件是从另一个容器中获取的！

CMD [ "./consignment-service" ]