FROM golang:1.20-alpine AS builder

ARG ARCH=amd64

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH
ENV GO_VERSION 1.20
ENV GO111MODULE on
ENV CGO_ENABLED=0

# Build dependencies
WORKDIR /go/src/
COPY . .
RUN apk update && apk add make git
RUN go get ./...
RUN mkdir /go/src/build 
RUN go build -a -gcflags=all="-l -B" -ldflags="-w -s" -o build/rubrik-exporter ./...

# Second stage
FROM alpine:3.17

COPY --from=builder /go/src/build/workspaceone-exporter /usr/local/bin/rubrik-exporter
CMD ["/usr/local/bin/rubrik-exporter"]
EXPOSE 9740
