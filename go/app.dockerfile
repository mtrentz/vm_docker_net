FROM golang:1.17.0 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go get && go build -o net_test main.go
CMD ["./net_test"]