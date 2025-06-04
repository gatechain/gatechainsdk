#!/usr/bin/make -f
SRCPATH := $(shell pwd)
SODIUM_PATH := $(shell dirname $(SRCPATH))/libsodium
export GO111MODULE = on
export CGO_CFLAGS +=-I$(SODIUM_PATH)/include -w
export CGO_LDFLAGS += -L$(SODIUM_PATH)/lib
export CGO_CXXFLAGS= -g -O2  -std=c++11 -w

sodium:
	./scripts/build_sodium.sh

test:
	@echo "Running tests..."
	CGO_ENABLED=1 \
	CGO_CFLAGS="-I$(SODIUM_PATH)/include -w" \
	CGO_LDFLAGS="-L$(SODIUM_PATH)/lib" \
	go test ./... -v

.PHONY: test