# Build stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download
ENV CGO_ENABLED=1
RUN go build -o backend .

# Run stage
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/backend .
COPY --from=builder /app/shop.db ./shop.db
RUN apt-get update && apt-get install -y ca-certificates libsqlite3-0 && rm -rf /var/lib/apt/lists/*
EXPOSE 8080
CMD ["./backend"]