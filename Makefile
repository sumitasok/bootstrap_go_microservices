# https://stackoverflow.com/questions/10121182/multiline-bash-commands-in-makefile
.PHONY: build run

build:
	{ \
		docker build -t gotools -f gotools.Dockerfile . ;\
		docker pull znly/protoc ;\
	}

run:
	@docker run -it --rm -p 9000:9000 --name my-running-app my-golang-app

shell:
	@docker run -it --rm -p 9000:9000 --name my-running-app my-golang-app bash

# stop the running containers asociated with this app
stop:
	@docker ps -a | grep data_access_layer | awk ' {print $1 } ' | xargs docker stop

# remove the running containers asociated with this app
# run `make stop` before this.
rm:
	@docker ps -a | grep data_access_layer | awk ' {print $1 } ' | xargs docker rm

status:
	@docker ps -a

force:
	@docker ps -a | grep Exited | awk '{print $1}' | xargs docker rm

gotools:
	docker run -it --rm -v `pwd`:/go/src/app gotools $(arg)

protofile:
	docker run --rm -v `pwd`:`pwd` -w `pwd` znly/protoc --go_out=plugins=grpc:. -I. $(path)

gitlog:
	git log --graph --oneline --decorate --all