.PHONY: build run

build:
	@docker build -t my-golang-app .

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