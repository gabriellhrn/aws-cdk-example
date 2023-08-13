.PHONY: build test

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

docker-build:
	@docker build --rm -t cdk .

test:
	@docker run --entrypoint go \
		-v $(ROOT_DIR)/app:/app \
		cdk -- test -v ./...

synth:
	@docker run \
		-v $(ROOT_DIR)/app:/app \
		cdk -- synth
