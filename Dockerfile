FROM golang:1.14.6-alpine3.12 as builder

COPY go.mod go.sum /go/src/github.com/barqus/fillq_backend/

WORKDIR /go/src/github.com/barqus/fillq_backend
RUN go mod download
COPY . /go/src/github.com/barqus/fillq_backend

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/fillq_backend github.com/barqus/fillq_backend

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/barqus/fillq_backend/build/fillq_backend /usr/bin/barqus
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/barqus","server"]