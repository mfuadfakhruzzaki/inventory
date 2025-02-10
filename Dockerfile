# Stage 1: Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o inventory-api ./cmd/app

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/inventory-api .
EXPOSE 8080
CMD ["./inventory-api"]
