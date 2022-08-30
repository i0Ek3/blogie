.PHONY: build clean tool help

GO=go

all: build

build:
	@$(GO) build -v .

tool:
	@$(GO) vet . | grep -v vendor; true
	gofmt -w .

clean:
	rm blogie
	@$(GO) clean -i .

help:
	@echo "[make] Build and clean Go code."
	@echo "make build: only for build"
	@echo "make tool: only for vet and format Go code"
	@echo "make clean: only for clean"
