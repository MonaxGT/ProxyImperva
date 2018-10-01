FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -v github.com/elazarl/goproxy
RUN go build

ENTRYPOINT /go/src/app/ProxyImperva
EXPOSE 8080