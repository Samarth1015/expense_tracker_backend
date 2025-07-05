# # ---- Build Stage ----
# FROM golang:1.23 as builder
# WORKDIR /app
# COPY . .
# WORKDIR /app/backend
# RUN go mod download
# RUN go build -o main .

# # ---- Run Stage ----
# FROM gcr.io/distroless/base-debian11
# WORKDIR /app
# COPY --from=builder /app/backend/main .
# COPY --from=builder /app/backend/.env .
# # If you don't have a .env file, comment out the above line
# EXPOSE 8080
# ENV PORT=8080
# ENV PROJECT=production
# # Set these as needed, or use docker-compose to inject them
# # ENV DB_HOST=db
# # ENV DB_USER=postgres
# # ENV DB_PASSWORD=example
# # ENV DB_NAME=expense
# # ENV DB_PORT=5432
# CMD ["./main"] 


# syntax=docker/dockerfile:1.5   <-- enables build‑cache mounts

########################
# ---- Build stage ----
########################
FROM golang:1.23.3 AS builder

WORKDIR /src

# 1️⃣ Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# 2️⃣ Copy the rest of your code
COPY . .

# 3️⃣ Build the Go app from the current working directory
RUN --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/app

########################
# ---- Run stage ------
########################
FROM gcr.io/distroless/static-debian11

WORKDIR /app
COPY --from=builder /out/app .

# Default envs (can be overridden in Deployment)
ENV PORT=8080 \
    PROJECT=production

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/app/app"]

