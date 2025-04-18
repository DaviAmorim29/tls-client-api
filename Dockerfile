# Build Stage
FROM golang:1.23 AS builder

# Disable CGO and set the target OS and architecture
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the binary
RUN go build -o app ./cmd/tls-client-api/main.go

# Final Stage
FROM gcr.io/distroless/static

COPY --from=builder /app/app /

EXPOSE 8080
EXPOSE 8081

CMD ["/app"]