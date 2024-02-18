FROM golang:1.22 AS dev
WORKDIR /go/src/myapp

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download

RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]

FROM golang:1.22 AS build
WORKDIR /go/src/myapp

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=.,target=. \
#    GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build \
    CGO_ENABLED=0 go build \
      -ldflags "-s -w" \
      -trimpath \
      -o /go/bin/myapp \
      ./cmd/server

FROM --platform=linux/amd64 gcr.io/distroless/static-debian12:nonroot AS prod
COPY --chown=nonroot:nonroot --from=build /go/bin/myapp /
ENTRYPOINT ["/myapp"]
