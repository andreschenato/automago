FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod init automago || true
RUN go mod tidy
RUN go build -o server .

FROM scratch

WORKDIR /app

COPY --from=builder /app .
COPY --from=builder /app/templates ./templates

EXPOSE 8080

CMD ["./server"]
