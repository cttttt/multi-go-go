all: test binaries

test:
	go test ./...

binaries:
	go build -o bin/ ./cmd/...

clean:
	rm -rf bin/

.PHONY: test binaries
