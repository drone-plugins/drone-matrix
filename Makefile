.PHONY: build
build: drone-plugin-matrix

drone-plugin-matrix: main.go
	CGO_ENABLED=0 go build -ldflags '-s -w'

.PHONY: docker
docker: drone-plugin-matrix
	docker build -t drone-plugin-matrix .

.PHONY: clean
clean:
	rm -f drone-plugin-matrix
