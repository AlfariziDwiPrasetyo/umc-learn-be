# Image golang
FROM golang:1.25.3-alpine AS builder

# Set working directory
WORKDIR /app

# Copy file to container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build Binary
RUN go build -o umc-learn-be ./cmd/main.go

# Run
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/umc-learn-be .
COPY --from=builder /app/internal/configs ./internal/configs

EXPOSE 9090

CMD [ "./umc-learn-be" ]
