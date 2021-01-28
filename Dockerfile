FROM golang:1.15.6-alpine3.12 AS test

ENV GOPATH=""
RUN pwd && ls -ltrh
COPY ./dialog/ /home/
RUN go version && cd /home/dialog && ls -ltrh && go mod init github.com/terujun/dialog && go build cmd/meal-dialog-bot/main.go && chmod 777 main && ls -ltrh

ENTRYPOINT [ "/home/dialog/main" ]