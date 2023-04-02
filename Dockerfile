FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o NyaBot .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/NyaBot .
ENV RUN_IN_DOCKER=1
EXPOSE 3000
CMD ["./NyaBot"]