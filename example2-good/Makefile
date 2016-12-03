
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
	@echo making an empty tar file and registering it as a docker image named 'empty'
	@echo
	tar c --files-from /dev/null | docker import - empty

a-tiny-app: a-tiny-app.go
	go build a-tiny-app.go

build: a-tiny-app
	docker build -t tiny-app .

run:
	docker run --rm -it tiny-app 
