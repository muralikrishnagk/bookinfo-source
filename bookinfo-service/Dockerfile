# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY src/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o bookinfo-service .

# Production stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/bookinfo-service .
EXPOSE 8080
CMD ["./bookinfo-service"]
