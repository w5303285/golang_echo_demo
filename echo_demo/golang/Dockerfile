FROM golang
RUN apt-get update && apt-get install -y vim
WORKDIR $GOPATH/src/http_demo
 
EXPOSE 8088

#初始化
RUN go mod init myapp
#框架包下载安装
RUN go get github.com/labstack/echo/v4

#执行Go Run 
CMD go run /go/src/http_demo/main.go 