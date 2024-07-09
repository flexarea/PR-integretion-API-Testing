#syntax=docker/dockerfile:1

FROM golang:1.22.3

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

#copy the compiled binary from the builder stage
COPY --from=builder /app/main /main 

EXPOSE 8080

#command to run executable
CMD ["/main"]

