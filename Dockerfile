FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod download && \
    CGO_ENABLED=0 go build -ldflags='-w -s' -o sequential_thinking main.go

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /app/sequential_thinking /sequential_thinking
ENTRYPOINT ["/sequential_thinking"]