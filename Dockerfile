#syntax=docker/dockerfile:1
ARG GO_VERSION=1.22.3
FROM golang:${GO_VERSION}-bookworm as builder

# Update package lists and install ca-certificates
RUN apt-get update && apt-get install -y ca-certificates
RUN update-ca-certificates
#set current work directory inside the container
WORKDIR /app

#copy mod and sum files
COPY go.mod go.sum /

#download dependencies
RUN go mod download

#copy source code into container
COPY . .

#build go app
RUN CGO_ENABLE=0 GOOS=linux go build -o /app/main cmd/*.go

#use small base image for final build
FROM debian:bookworm

# Update package lists and install ca-certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

#copy the compiled binary from the builder stage
COPY --from=builder /app/main /main 

# Ensure CA certificates are available in the final image
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

#command to run executable
CMD ["/main"]

