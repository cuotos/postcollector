FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .

RUN go build -o postCollector main.go

FROM alpine:3.22

COPY --from=builder /app/postCollector /postCollector

CMD ["/postCollector"]

