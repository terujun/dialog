FROM golang:1.15.6-alpine3.12 AS test

ENV GOPATH=""
RUN mkdir /go/src/dialog
COPY ./* /go/src/dialog/
RUN cd /go/src/dialog && ls -ltrh && go mod init meal.com/dialog && go build && chmod 777 dialog && ls -ltrh

ENTRYPOINT [ "/go/src/dialog/dialog" ]