all: test

clean:
	git clean -Xdf -e '!.vagrant'

test: test-unit

test-unit:
	godep go test ./...

.PHONY: all clean test test-unit
