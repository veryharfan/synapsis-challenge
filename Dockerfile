FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go build -o myapp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .

EXPOSE 8080

CMD ["./myapp"]
