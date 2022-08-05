.PHONY: build clean

GO=go

build:
	@$(GO) build .
	@./scripts/setup.sh

clean:
	@rm blogie
