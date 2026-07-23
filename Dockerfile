FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /client ./cmd

FROM alpine:3.14 AS stage

WORKDIR /app

COPY --from=builder /client .

CMD ["./client"]