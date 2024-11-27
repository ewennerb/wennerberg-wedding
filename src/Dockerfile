FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
COPY go.mod .
COPY go.sum .
RUN go mod download
ENV PORT=8080

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main .

FROM alpine:latest
COPY --from=builder /app/main /app/main
EXPOSE 8080
ENTRYPOINT ["/app/main"]
