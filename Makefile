ARCH:=$(shell uname -m)
path?="test"
.PHONY: all dir

all: dir

test:
	go test $@

dir:
	echo $(path)
	mkdir $(path)
	cd $(path) && go mod init && touch main.go


