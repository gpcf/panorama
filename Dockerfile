FROM docker.io/golang:1.20-alpine AS backend_builder
WORKDIR /app
RUN apk add git
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY cmd ./cmd
COPY pkg ./pkg
RUN go build -v ./cmd/panorama

FROM docker.io/node:18-alpine AS ui_builder
WORKDIR /app/ui
COPY ui/package*.json ./
RUN npm install
COPY ui .
RUN npm run build

FROM docker.io/alpine:latest
WORKDIR /app
COPY --from=backend_builder /app/panorama ./
COPY --from=ui_builder /app/ui/build static
COPY config.example.toml /etc/panorama/config.toml

ENTRYPOINT ["./panorama", "--config", "/etc/panorama/config.toml"]
