# Use the official Golang image as the base image
FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./be-user

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/be-user ./be-user
CMD ["/app/be-user"]