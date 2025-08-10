FROM golang:1.24-alpine AS builder

WORKDIR /app1

COPY /backend/auth/go.mod /backend/auth/go.sum ./

COPY /backend/auth .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app1/sususport ./cmd/main.go

FROM alpine:latest

COPY --from=builder /app1/sususport /app1/
COPY backend/auth/.env /app1/

WORKDIR /app1

CMD ["./sususport"]