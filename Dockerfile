FROM golang:alpine as builder
WORKDIR /app
COPY . .

RUN go build -o postCollector main.go

FROM alpine

COPY --from=builder /app/postCollector /postCollector

CMD /postCollector

