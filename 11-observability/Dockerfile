FROM golang:alpine AS builder

WORKDIR /usr/app

COPY . .
RUN go build -v -o /usr/app/bin/todos cmd/todos/main.go

FROM alpine:latest

WORKDIR /usr/app

COPY --from=builder /usr/app/bin/todos /usr/app/todos

CMD ["/usr/app/todos"]
