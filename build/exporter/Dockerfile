FROM golang:1.16.3-alpine3.12 AS build

COPY ./ /go/src/github.com/lexfrei/gow/
WORKDIR /go/src/github.com/lexfrei/gow/cmd/exporter

RUN go get ./... && CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o ow-exporter

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/github.com/lexfrei/gow/cmd/exporter/ow-exporter /

EXPOSE 9420/tcp

ENTRYPOINT ["/ow-exporter"]