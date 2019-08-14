.PHONY: all
all: test bench

.PHONY: test
test:
	go test -v .

.PHONY: bench
bench:
	go test -bench=. -benchmem
