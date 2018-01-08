# vim: set ft=dockerfile:
FROM golang:1.9 as builder

ENV CGO_ENABLED=0

RUN go get github.com/ptman/drone-plugin-matrix

FROM scratch

COPY --from=builder /go/bin/drone-plugin-matrix /drone-plugin-matrix
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT [ "/drone-plugin-matrix" ]
