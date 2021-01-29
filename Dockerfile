FROM golang:1.15.6-alpine3.12 AS test

ARG TOKEN
ARG PORT
RUN export TOKEN=${TOKEN} && export PORT=${PORT}
RUN apk add git && mkdir /home/go/ && cd /home/go/ && git clone https://github.com/terujun/dialog.git && cd dialog/ && go mod init github.com/terujun/dialog && go get -u github.com/labstack/echo/... && go get github.com/mattn/go-jsonpointer && go mod tidy && go build cmd/meal-dialog-bot/main.go && chmod 777 main && ls -ltrh

ENTRYPOINT [ "/home/go/dialog/main" ]