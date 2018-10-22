FROM golang:1.9

WORKDIR /go/src/app
COPY . .

# requires cgo to work
# export CGO_ENABLED=1

RUN cd proto_server && go get -d -v ./...
RUN cd proto_server && go install -v ./...

CMD ["go", "run", "./proto_server/main.go"]