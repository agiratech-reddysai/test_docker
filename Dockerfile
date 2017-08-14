FROM golang

ADD . /go/src/github.com/golang/test_docker
RUN go install github.com/golang/test_docker
ENTRYPOINT /go/bin/test_docker

EXPOSE 8080