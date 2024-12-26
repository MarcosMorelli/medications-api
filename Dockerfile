FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o medication-api ./cmd

FROM scratch
WORKDIR /
COPY --from=builder /app/medication-api .

EXPOSE 8080

CMD ["./medication-api"]
