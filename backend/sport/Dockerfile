FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY /backend/sport/go.mod /backend/sport/go.sum ./

COPY /backend/sport .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/sususport ./cmd/main.go

FROM alpine:latest

COPY --from=builder /app/sususport /app/
COPY backend/sport/.env /app/

WORKDIR /app

CMD ["./sususport"]