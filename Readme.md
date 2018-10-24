# Readme

How to generate go files from proto files:
We use proto compiler from this unofficial docker image from this [github repo](https://github.com/znly/docker-protobuf) forked over [here](https://github.com/go-get/docker-protobuf)

Install Go Dep:

```
go get -u github.com/golang/dep/cmd/dep
```

To download the docker image:
```
docker run --rm znly/protoc --help
```

to generate the go file:
```
docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --go_out=. -I. file-to-convert.proto
```
```
docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --go_out=plugins=grpc:. -I. data/product.proto
```

Docker Compose

`.env.samples` file holds the configs for your docker compose/deployment.
copy this file to `.env` and edit it to satisfy your setup.
any new ENV VAR+VAL should be updated in the .env.samples file, and copied back to .env
.env is git ignored. Do not check in .env file or your local configs into repo.

run:

```
docker-compose up
```

After updating .env file with your configs.