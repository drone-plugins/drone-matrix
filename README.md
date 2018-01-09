# drone-matrix

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-matrix/status.svg)](http://beta.drone.io/drone-plugins/drone-matrix)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-matrix?status.svg)](http://godoc.org/github.com/drone-plugins/drone-matrix)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-matrix)](https://goreportcard.com/report/github.com/drone-plugins/drone-matrix)
[![](https://images.microbadger.com/badges/image/plugins/matrix.svg)](https://microbadger.com/images/plugins/matrix "Get your own image badge on microbadger.com")

Drone plugin for sending build notifications to [Matrix](https://matrix.org/). For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-matrix/).

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-matrix
docker build --rm -t plugins/matrix .
```

### Usage

```
docker run --rm \
  -e PLUGIN_ROOMID=0123456789abcdef:matrix.org \
  -e PLUGIN_USERNAME=yourbot \
  -e PLUGIN_PASSWORD=p455w0rd \
  plugins/matrix
```
