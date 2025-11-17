FROM golang:1.25 AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

COPY backend/ .

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/api

FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/server .

USER nonroot:nonroot

ENV PORT=8080

EXPOSE 8080

CMD ["./server"]
