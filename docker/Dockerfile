FROM golang:1.20-alpine AS base

# Install dependency based on package manager
FROM base AS deps
WORKDIR /app
COPY ../go.mod go.sum ./
RUN go mod download

# Rebuild the source code only when needed
FROM base AS go-builder
WORKDIR /app
COPY --from=deps /go/pkg/ /go/pkg/
COPY .. .

RUN go build -ldflags="-X github.com/zjc17/pprof-web/internal/version.Version=1.0" -o ./pprof-web

# Runner
FROM alpine AS runner
RUN apk add --update graphviz
RUN rm -rf /var/cache/apk/*
WORKDIR /
COPY --from=go-builder /app/pprof-web /bin/pprof-web

EXPOSE 8080
ENTRYPOINT ["pprof-web"]



