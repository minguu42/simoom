FROM golang:1.23 AS dev
WORKDIR /go/src/myapp

RUN go install github.com/air-verse/air@latest

RUN --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download

CMD ["air", "-c", ".air.toml"]

FROM golang:1.23 AS build
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
      ./api

FROM gcr.io/distroless/static-debian12:nonroot AS prod
COPY --chown=nonroot:nonroot --from=build /go/bin/myapp /
ENTRYPOINT ["/myapp"]
