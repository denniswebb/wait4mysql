FROM golang:latest

WORKDIR /go/src/github.com/denniswebb/wait4mysql

COPY main.go .

RUN go get ./... && \
    go build -o wait4mysql

FROM alpine:latest
MAINTAINER "Dennis Webb <dhwebb@gmail.com>"

COPY --from=0 /go/src/github.com/denniswebb/wait4mysql/wait4mysql /usr/local/bin/wait4mysql

ENTRYPOINT ["wait4mysql"]
