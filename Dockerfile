FROM golang:1.15.6-alpine3.12 AS test

ARG TOKEN
ARG PORT
ARG CONFIGDIRPATH
ENV TOKEN=${TOKEN} PORT=${PORT} CONFIGDIRPATH=${CONFIGDIRPATH}
RUN apk add git && mkdir /home/go/ && cd /home/go/ && git clone https://github.com/terujun/dialog.git && cd dialog/ && go mod init github.com/terujun/dialog && go get -u github.com/labstack/echo/... && go get github.com/mattn/go-jsonpointer && go get cloud.google.com/go/firestore && go mod tidy && go build cmd/meal-dialog-bot/main.go && chmod 777 main && ls -ltrh

ENTRYPOINT [ "/home/go/dialog/main" ]