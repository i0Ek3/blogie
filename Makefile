.PHONY: build test clean

GO=go

build:
	@export GO111MODULE=on
	@export GOPROXY=https://goproxy.cn
	@$(GO) mod tidy
	@$(GO) mod vendor
	@docker-compose up -d
	@$(GO) build .

test:
	@$(GO) test -v .

clean:
	@rm blogie
