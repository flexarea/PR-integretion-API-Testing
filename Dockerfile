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

