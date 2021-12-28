ARCH:=$(shell uname -m)
path?="test"
.PHONY: create_pb

all: generate_pb

generate_pb:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pb/*.proto


