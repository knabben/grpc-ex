FROM golang:1.9.2-alpine3.7

RUN apk update && apk add git tcpdump

WORKDIR /go/src/github.com/knabben/grpc-ex/ 
COPY . /go/src/github.com/knabben/grpc-ex/ 

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

RUN go build .

ENTRYPOINT ["./grpc-ex"]
CMD ["serve"]