all: test build

build:
	godep go buil

clean:
	git clean -Xdf -e '!.vagrant'

test: test-unit

test-unit:
	godep go test ./...

dev:
	docker-compose -f docker-dev.yml run --service-ports --rm dev /bin/bash

.PHONY: all clean test test-unit
