# Base container image
FROM golang:1.14-alpine AS builder

# Create app folder
WORKDIR /go/src/app

# Install and cache dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy app files into container
COPY . .

# Build the app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o main


# Smallest container image
FROM scratch

# Copy built executable from base image to here
COPY --from=builder /go/src/app/main /

# Run the app
ENTRYPOINT [ "/main" ]
