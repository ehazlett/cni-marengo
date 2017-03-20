CGO_ENABLED=0
all: build

build:
	@go build -v .

.PHONY: build
