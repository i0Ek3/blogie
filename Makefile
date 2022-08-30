.PHONY: build clean

GO=go

build:
	@$(GO) build .

clean:
	@rm blogie
