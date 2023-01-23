FROM golang:1.17

RUN go env -w GO111MODULE=on

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app

ADD . /app

CMD go mod init game

CMD go mod tidy

CMD go get -u github.com/gin-gonic/gin

EXPOSE 8080