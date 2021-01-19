FROM golang:1.15.6-alpine3.12 AS test

RUN pwd
RUN ls
RUN cd workspqce/dialog && go mod init meal.com/dialog
RUN go build && chmod 777 ./dialog

ENTRYPOINT [ "dialog" ]