FROM golang:alpine
MAINTAINER "Dennis Webb <dhwebb@gmail.com>"

RUN apk add --update --no-cache git

RUN go get github.com/denniswebb/wait4mysql && \
    go install github.com/denniswebb/wait4mysql

ENTRYPOINT ["wait4mysql"]
