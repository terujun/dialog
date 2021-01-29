FROM golang:1.15.6-alpine3.12 AS test

COPY ./* /go/src/
RUN  apk add git && go get -u github.com/labstack/echo/... && go get github.com/mattn/go-jsonpointer && cd /go/src/ && ls -ltrh && go mod init github.com/terujun/dialog && go build cmd/meal-dialog-bot/main.go && chmod 777 main && ls -ltrh

ENTRYPOINT [ "/home/main" ]