FROM golang:1.9.2-alpine3.7

WORKDIR /usr/local/go/src/github.com/knabben/grpc/
COPY . /usr/local/go/src/github.com/knabben/grpc/ 

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

RUN go build .
CMD sh