.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: lint
lint:
	golangci-lint run ./...
	go fmt ./...

.PHONY: clear
clear:
	rm -f logs/*
