ARCH:=$(shell uname -m)

.PHONY: all

all: test_cover

test:
	go test $@

test_cover:
	go test -cover -coverprofile=cover.out
	go tool cover -html=cover.out


