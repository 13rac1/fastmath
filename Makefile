.PHONY: all

all: test bench

test:
	go test -v .

bench:
	go test -v -bench=.
