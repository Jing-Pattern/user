FROM golang:1.20
LABEL authors="jingze"
# 设置工作目录
WORKDIR /app
COPY . .

RUN go env -w GOPROXY=https://proxy.golang.com.cn,direct

RUN  go build  -o app .

# 暴露Gateway服务的端口（假设Gateway服务监听在3000端口上）
EXPOSE 8080

# 定义启动命令
CMD ["./app","-f", "etc/user.yaml"]