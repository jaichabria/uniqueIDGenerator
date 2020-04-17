FROM golang

ADD . $GOPATH/src/UniqueIDGenerator

WORKDIR $GOPATH/src/UniqueIDGenerator

RUN go get -d -v

RUN go build UniqueIDGenerator

ENTRYPOINT ["./UniqueIDGenerator"]

EXPOSE 8080 8043

