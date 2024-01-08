# 使用官方的 Go 镜像作为基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 复制本地代码到容器内的工作目录
COPY . .

# 编译 Go 应用程序
RUN go build -o /app/bin/echo-app /app/src/main.go

# 设置容器启动命令
CMD ["/app/bin/echo-app"]