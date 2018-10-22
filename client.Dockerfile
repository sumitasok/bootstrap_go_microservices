FROM golang:1.9

WORKDIR /go/src/app
COPY . .

# requires cgo to work
# export CGO_ENABLED=1

RUN cd proto_client && go get -d -v ./...
RUN cd proto_client && go install -v ./...

# run from the installed binary
CMD ["go", "run", "./proto_client/main.go"]