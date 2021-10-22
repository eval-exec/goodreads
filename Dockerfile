FROM golang:1.16 as builder
WORKDIR /app
COPY go.mod /app/
COPY go.sum /app/
RUN go mod download

RUN go build .

FROM alpine:3.16

COPY --from=builder /app/goodreads /goodreads

ENTRYPOINT ["/goodreads", "--tag", "hope"]

