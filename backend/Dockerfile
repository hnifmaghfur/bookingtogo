# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app main.go

# Run stage
FROM alpine:latest
RUN apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /app/app ./app
COPY --from=builder /app/.env .env
COPY --from=builder /app/database ./database
COPY --from=builder /app/config ./config
COPY --from=builder /app/internal ./internal
COPY --from=builder /app/routes ./routes
EXPOSE 8080
CMD ["./app"]
