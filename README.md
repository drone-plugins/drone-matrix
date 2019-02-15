# drone-matrix

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-matrix/status.svg)](http://cloud.drone.io/drone-plugins/drone-matrix)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/matrix.svg)](https://microbadger.com/images/plugins/matrix "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-matrix?status.svg)](http://godoc.org/github.com/drone-plugins/drone-matrix)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-matrix)](https://goreportcard.com/report/github.com/drone-plugins/drone-matrix)

Drone plugin for sending build notifications to [Matrix](https://matrix.org/). For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-matrix/).

## Build

Build the binary with the following command:

```console
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/linux/amd64/drone-matrix
```

## Docker

Build the Docker image with the following command:

```console
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag plugins/matrix .
```

## Usage

```console
docker run --rm \
  -e PLUGIN_ROOMID=0123456789abcdef:matrix.org \
  -e PLUGIN_USERNAME=yourbot \
  -e PLUGIN_PASSWORD=p455w0rd \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/matrix
```
