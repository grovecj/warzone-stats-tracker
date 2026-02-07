# Stage 1: Build frontend
FROM node:22-alpine AS frontend
WORKDIR /app/web
COPY web/package*.json ./
RUN npm ci
COPY web/ .
RUN npm run build

# Stage 2: Build backend
FROM golang:1.25-alpine AS backend
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=frontend /app/web/dist ./web/dist
RUN CGO_ENABLED=0 go build -o server ./cmd/server

# Stage 3: Runtime
FROM alpine:3.21
RUN apk add --no-cache ca-certificates
COPY --from=backend /app/server /server
COPY --from=backend /app/migrations /migrations
EXPOSE 8080
CMD ["/server"]
