FROM golang

COPY ./ /go/src/github.com/lexfrei/gow/
WORKDIR /go/src/github.com/lexfrei/gow/cmd/exporter

RUN go get ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ow-exporter

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=0 /go/src/github.com/lexfrei/gow/cmd/exporter/ow-exporter /
ENTRYPOINT ["/ow-exporter"]