FROM golang:1.17-alpine as builder

WORKDIR /workspace

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o /out/httpserver .

FROM alpine:3.12

COPY --from=builder /out/httpserver /app/httpserver

EXPOSE 3000
ENTRYPOINT ["/app/httpserver"]