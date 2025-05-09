FROM golang:1.23.1-bullseye as builder
RUN apt-get update && apt-get install -y nocache git ca-certificates && update-ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/bin/lighthouse pkg/controller/main.go


FROM debian:buster-slim
RUN useradd -ms /bin/bash --uid 1000 lighthouse
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY --from=builder /app/bin /app
EXPOSE 8080
EXPOSE 50051
USER lighthouse

CMD ["./lighthouse"]