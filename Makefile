.PHONY: all
all:
	make clear
	make run
	clear
	make getStat

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

.PHONY: getStat
getStat:
	grep "Added" logs/buffer* | wc -l
	grep "full" logs/buffer* | wc -l
	grep "level" logs/buffer* | wc -l

	grep "broken" logs/cpu1* | wc -l
	grep "level" logs/cpu1* | wc -l

	grep "broken" logs/cpu2* | wc -l
	grep "level" logs/cpu2* | wc -l

	grep "One of CPU" logs/server* | wc -l
	grep "Both" logs/server* | wc -l
	grep "handled" logs/server* | wc -l
	grep "level" logs/server* | wc -l
