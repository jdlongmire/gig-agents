# Stage 1: Frontend build
FROM oven/bun:latest AS frontend
WORKDIR /app
COPY web/ ./web/
WORKDIR /app/web
RUN bun install && bun run build

# Stage 2: Go build (CGO_ENABLED=1 required for go-sqlite3)
FROM golang:1.24-bookworm AS backend
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY sqlc.yaml ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /crewport ./cmd/crewport

# Stage 3: Runtime
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    sqlite3 \
    && rm -rf /var/lib/apt/lists/*

# Copy Go binary
COPY --from=backend /crewport /crewport

# Copy built frontend to web/dist/ relative to binary working directory
COPY --from=frontend /app/web/dist /web/dist

EXPOSE 8080
CMD ["/crewport"]
