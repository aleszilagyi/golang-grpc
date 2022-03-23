FROM golang:1.17
WORKDIR /go/src

ENV PATH="go/bin:${PATH}"

RUN apt-get update && apt-get install protobuf-compiler -y \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

CMD ["tail", "-f", "/dev/null"]