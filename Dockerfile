FROM golang:alpine

RUN mkdir /app

ADD server_* /app/

WORKDIR /app

CMD ["/app/server_linux_amd64"]