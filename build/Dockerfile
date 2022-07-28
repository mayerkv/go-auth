FROM golang:1.18-alpine3.14 as builder
WORKDIR /auth
COPY go.* .
RUN go mod tidy
COPY . .
RUN go build -o bin/server cmd/auth/*.go

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /auth/bin/server .
CMD ["/app/server"]