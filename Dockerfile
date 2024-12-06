FROM golang:1.21.4 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .
FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/app .
COPY .env .env
EXPOSE 80
CMD ["./app"]