FROM golang:1.17 as builder

# Enable Go modules
ENV GO111MODULE=on

WORKDIR /app
# COPY go.mod .
# RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:latest
# mailcap adds mime detection and ca-certificates help with TLS (basic stuff)
RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app
USER app
WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT ["./app"]
