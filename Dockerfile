FROM golang:1.15.6-alpine3.12 AS test


COPY ./* /home/
RUN GOPATH="" && go env && cd /home/ && ls -ltrh && go mod init github.com/terujun/dialog && go build cmd/meal-dialog-bot/main.go && chmod 777 main && ls -ltrh

ENTRYPOINT [ "/home/main" ]