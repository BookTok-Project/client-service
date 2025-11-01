FROM golang:1.24.1 AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o run ./cmd/

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /build/run .
COPY --from=builder /build/.docker/remote/config.json .
EXPOSE 8000
ENTRYPOINT ["./run"]
