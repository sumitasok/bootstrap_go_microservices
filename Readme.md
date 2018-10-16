# Readme

How to generate go files from proto files:
We use proto compiler from this unofficial docker image from this [github repo](https://github.com/znly/docker-protobuf) forked over [here](https://github.com/go-get/docker-protobuf)

To download the docker image:
```
docker run --rm znly/protoc --help
```

to generate the go file:
```
docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --go_out=. -I. file-to-convert.proto
```

