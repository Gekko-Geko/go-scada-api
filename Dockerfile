FROM golang:1.24.4-alpine3.22 AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o scada-api .

FROM scratch
ENV GIN_MODE=release
WORKDIR /
COPY --from=builder /app/scada-api .
EXPOSE 8080
ENTRYPOINT ["/scada-api"]

