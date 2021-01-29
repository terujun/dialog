FROM golang:1.15.6-alpine3.12 AS test

RUN  apk add git && mkdir /home/go/ && cd /home/go/ &&  git clone https://github.com/terujun/dialog.git  && go get -u github.com/labstack/echo/... && go get github.com/mattn/go-jsonpointer && ls -ltrh && go mod init github.com/terujun/dialog && cat go.mod && cd dialog/　&& go build cmd/meal-dialog-bot/main.go && chmod 777 main && ls -ltrh

ENTRYPOINT [ "/home/main" ]