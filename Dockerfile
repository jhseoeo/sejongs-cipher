FROM golang:1.20 AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o ./ ./...

FROM alpine:3.9
RUN apk add ca-certificates

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/app", "/build/.env", "/"]

# Expose necessary port.
EXPOSE 8000

# Command to run when starting the container.
ENTRYPOINT ["/app"]