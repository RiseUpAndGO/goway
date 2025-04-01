FROM golang:1.24

WORKDIR /app

COPY /app /app

RUN go install github.com/air-verse/air@latest
RUN go mod download

CMD ["air", "-c", ".air.toml"]