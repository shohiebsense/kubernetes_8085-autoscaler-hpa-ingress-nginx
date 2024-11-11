FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o networkTrafficMonitor .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/networkTrafficMonitor .
RUN chmod +x networkTrafficMonitor

EXPOSE 8085

CMD ["./networkTrafficMonitor"]
