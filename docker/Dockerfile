#FROM golang:1.12.0-alpine3.9
#RUN mkdir /app
#ADD . /app
#WORKDIR /app
### Add this go mod download command to pull in any dependencies
#RUN go mod download
### Our project will now successfully build with the necessary go libraries included.
#RUN go build -o main .
### Our start command which kicks off
### our newly created binary executable
#CMD ["/app/main"]

# Pull base image
#FROM golang:alpine
#
## Install git
## RUN apk update && apk add --no-cache git
#
## Where our file will be in the docker container
#WORKDIR /opt/go-app
#
## Copy the source from the current directory to the working directory inside the container
## Source also contains go.mod and go.sum which are dependency files
#COPY . .
#
## Get Dependency
#RUN go mod download
#
## Install Air for hot reload
#RUN go get -u github.com/cosmtrek/air
#
## The ENTRYPOINT defines the command that will be ran when the container starts up
## In this case air command for hot reload go apps on file changes
#ENTRYPOINT air



#FROM golang:1.10
#
## Set the Current Working Directory inside the container
#WORKDIR $GOPATH/src/github.com/codefresh-contrib/go-sample-app
#
## Copy everything from the current directory to the PWD (Present Working Directory) inside the container
#COPY . .
#
## Download all the dependencies
#RUN go get -d -v ./...
#
## Install the package
#RUN go install -v ./...
#
## This container exposes port 8080 to the outside world
#EXPOSE 8080
#
## Run the executable
#CMD ["go-sample-app"]


#FROM golang:latest
## Set the Current Working Directory inside the container
#WORKDIR /app
#
## Copy go mod and sum files
#COPY go.mod go.sum ./
#
## Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download
#
## Copy the source from the current directory to the Working Directory inside the container
#COPY . .
#
## Build the Go app
#RUN go build -o main .
#
## Expose port 8080 to the outside world
#EXPOSE 8080
#
## Command to run the executable
#CMD ["./main server"]

#FROM golang:1.13
#
#COPY . /gosrc
#WORKDIR /gosrc
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fillq
#
#
#FROM alpine:3.9
#
#WORKDIR /fillq
#COPY --from=0 /gosrc/fillq fillq
#
#CMD ["/gosrc/fillq server"]


# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ../go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY .. .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o github.com/barqus/fillq_backend .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main server"]