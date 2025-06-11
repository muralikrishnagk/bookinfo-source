# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY src/ ./

# Define build arguments for versioning
ARG VERSION=dev
ARG COMMIT=none
ARG BUILD_DATE=unknown

# Build the Go app with ldflags to inject version information
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.buildDate=${BUILD_DATE}" -o bookinfo .

# Production stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/bookinfo .
EXPOSE 8080
CMD ["./bookinfo"]
