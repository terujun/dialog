FROM golang:1.15.6-alpine3.12 AS test

RUN go mod init meal.com/dialog
RUN go build && chmod 777 ./dialog

ENTRYPOINT [ "dialog" ]