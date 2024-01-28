FROM golang:1.21 AS base
WORKDIR /go/src/myapp

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download

FROM base AS dev
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]

FROM base AS build
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod/ \
    CGO_ENABLED=0 go build \
      -ldflags "-s -w" \
      -trimpath \
      -o /go/bin/myapp \
      ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot AS prod
COPY --chown=nonroot:nonroot --from=build /go/bin/myapp /
ENTRYPOINT ["/myapp"]
