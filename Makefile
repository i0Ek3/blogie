.PHONY: build test clean

GO=go

build:
	@$(GO) build .

linux:
	CGO_ENABLED=0 GOOS=linux go build -a -o blogie .
windows:
	CGO_ENABLED=0 GOOS=windows go build -a -o blogie.exe .

test:
	@$(GO) test -v .

clean:
	@rm blogie
