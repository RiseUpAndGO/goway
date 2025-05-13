FROM golang:1.24

WORKDIR /app

COPY /app /app

RUN apt-get update && apt-get install -y protobuf-compiler

RUN go install github.com/air-verse/air@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go mod download

CMD ["air", "-c", ".air.toml"]