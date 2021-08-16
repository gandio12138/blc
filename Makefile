BINARY := blockchain

all: build test

build: deps
	@echo "====> Go build"
	@go build -o $(BINARY)

deps:
	@go get -u github.com/boltdb/bolt

test:
	./$(BINARY) printChain
	./$(BINARY) addBlock -data "Send 1 BTC to Ivan"
	./$(BINARY) addBlock -data "Pay 0.31337 BTC for a coffee"
	./$(BINARY) printChain

.PHONY: build deps test