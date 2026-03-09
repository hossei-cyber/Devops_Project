# -------- Builder stage --------

FROM golang:1.26.0-alpine AS builder

ARG SERVICE
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN test -n "$SERVICE"

# Build static binary for selected service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /out/app ./${SERVICE}/cmd/main.go


# -------- Runtime stage --------

FROM scratch
WORKDIR /app

COPY --from=builder /out/app ./app

EXPOSE 8080
ENTRYPOINT ["/app/app"]