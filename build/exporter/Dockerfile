FROM golang:1.15.7-alpine3.12 AS build

COPY ./ /go/src/github.com/lexfrei/gow/
WORKDIR /go/src/github.com/lexfrei/gow/cmd/exporter

RUN go get ./... && CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o ow-exporter

FROM alpine:3.13.1

RUN apk add -u --no-cache ca-certificates=20191127-r5

COPY --from=build /go/src/github.com/lexfrei/gow/cmd/exporter/ow-exporter /

EXPOSE 9420/tcp

ENTRYPOINT ["/ow-exporter"]