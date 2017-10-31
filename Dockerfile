# vim: set ft=dockerfile:
FROM alpine:3.6
# Author with no obligation to maintain
MAINTAINER Paul Tötterman <paul.totterman@gmail.com>

RUN apk --no-cache add ca-certificates
ADD drone-plugin-matrix /
ENTRYPOINT /drone-plugin-matrix
