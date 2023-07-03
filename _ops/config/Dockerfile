FROM golang:1.20 as builder

COPY _ops/config /opt/xi-config
WORKDIR /opt/xi-config

RUN go build .

FROM ubuntu:22.04

COPY --from=builder /opt/xi-config/xi-config /usr/local/bin/xi-config
ENTRYPOINT xi-config
