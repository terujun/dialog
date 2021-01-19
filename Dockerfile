FROM golang:1.15.6-alpine3.12 AS test

COPY /workspace/dialog /go/src/
RUN ls
RUN cd /go/src/dialog && go mod init meal.com/dialog
RUN go build && chmod 777 ./dialog

ENTRYPOINT [ "dialog" ]