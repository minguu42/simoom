FROM golang:1.22 AS base
WORKDIR /go/src/myapp

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download

FROM base AS dev
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]

FROM base AS build
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=.,target=. \
    GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build \
      -ldflags "-s -w" \
      -trimpath \
      -o /go/bin/myapp \
      ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot AS prod
COPY --chown=nonroot:nonroot --from=build /go/bin/myapp /
ENTRYPOINT ["/myapp"]
