FROM golang:1.22-bullseye AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

FROM base AS dev
RUN go install github.com/cosmtrek/air@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest
COPY . .
CMD ["air"]

FROM base AS production
RUN useradd -u 1001 dockremap
COPY . .
RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o explorer-api

FROM scratch
ENV APP_ENV=production
WORKDIR /
COPY --from=production /etc/passwd /etc/passwd
COPY --from=production /app/explorer-api explorer-api
USER dockremap
EXPOSE 8080
CMD ["/explorer-api"]
