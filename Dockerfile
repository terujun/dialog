FROM golang:1.15.6-alpine3.12 AS test

RUN  apk add git && mkdir /home/go/ && cd /home/go/ && export GOPATH="/home/go" &&  git clone https://github.com/terujun/dialog.git  && cd dialog/ && go mod init github.com/terujun/dialog && go env && cat go.mod && go get -u github.com/labstack/echo/... && go get github.com/mattn/go-jsonpointer && cat go.mod && go mod tidy && ls -ltrh && cat go.mod && go build cmd/meal-dialog-bot/main.go && chmod 777 main && ls -ltrh

ENTRYPOINT [ "/home/go/main" ]