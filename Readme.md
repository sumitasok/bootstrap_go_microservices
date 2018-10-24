# Readme

Add this to you `.git/config` file

```
[remote "bootstrap"]
        url = git@github.com:sumitasok/bootstrap_go_microservices.git
        fetch = +refs/heads/*:refs/remotes/bootstrap/*
```

Then do git merge commits from that remote.

How to generate go files from proto files:
We use proto compiler from this unofficial docker image from this [github repo](https://github.com/znly/docker-protobuf) forked over [here](https://github.com/go-get/docker-protobuf)

to generate the proto go file:
```
make protofile path=procedure/product/product.proto
```
Where path is path to your protofile, and the *.pb.go file will be colocated to the input file.

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

In order to run migrations or dep:
If you want more tools for development, add those to `gotools.Dockerfile`

```
make gotools arg="sql-migrate --help"
```

pass your command as string.

In order to do dep ensure, run
```
make gotools arg="dep ensure"
```

This will update your vendor folder. Vendor folder is git ignored.