
help:
	@echo
	@echo "    make empty -- creates and registers the minimal base docker image"
	@echo "    make build -- builds a little docker image with a single app atop the empty image"
	@echo "    make run   -- runs the app"
	@echo
	@echo "    see also, https://docs.docker.com/engine/userguide/eng-image/baseimages/"
	@echo

all: empty build run

empty:
	@echo
	@echo "making an empty tar file and registering it as a docker image named 'empty'"
	@echo
	tar c --files-from /dev/null | docker import - empty

a-tiny-web-app: a-tiny-web-app.go
	go build -ldflags "-linkmode external -extldflags -static" a-tiny-web-app.go

build: a-tiny-web-app
	docker build -t tiny-web-app .

run:
	docker run --rm -p 8181:8080 -it tiny-web-app 
