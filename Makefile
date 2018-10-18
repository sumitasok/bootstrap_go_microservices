.PHONY: build run

build:
	@docker build -t my-golang-app .

run:
	@docker run -it --rm -p 9000:9000 --name my-running-app my-golang-app

shell:
	@docker run -it --rm -p 9000:9000 --name my-running-app my-golang-app bash
