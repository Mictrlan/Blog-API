.PHONY: build clean tool help

build:
	go build main.go 

clear:
	rm -rf blog-api
	go clean -i

tool:
	go tool vet . |& grep -v vendor; true
	gofmt -w .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make clean: remove object files and cached files"
