#FROM --platform=linux/amd64 golang:1.23-alpine AS builder
# This is for local dev
FROM --platform=linux/amd64 golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
COPY go.mod .
COPY go.sum .
RUN go mod download
ENV PORT=8080

RUN GOOS=linux GOARCH=amd64 go build -o main
# This is for local dev.
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main .

FROM --platform=linux/amd64 alpine:latest
COPY --from=builder /app/main /app/main
EXPOSE 8080
ENTRYPOINT ["/app/main"]
