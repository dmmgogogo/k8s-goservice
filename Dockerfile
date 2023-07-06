# 基于Golang的基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 将Go模块依赖复制到镜像中
COPY go.mod ./

# 下载并安装依赖包
RUN go mod download

# 复制项目代码到镜像中
COPY . .

# 编译Go应用程序
RUN go build -o main .

# 指定启动容器时执行的命令
CMD ["./main"]

