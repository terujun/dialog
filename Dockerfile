FROM golang:1.15.6-alpine3.12 AS test

ENV GOPATH=""
RUN mkdir /go/src/dialog
COPY ./* /go/src/dialog/
RUN cd echo $GOPATH && /go/src/dialog && ls -ltrh && go mod init github.com/terujun/dialog && go build cmd/meal-dialog-bot/main.go && chmod 777 main && ls -ltrh

ENTRYPOINT [ "/go/src/dialog/main" ]