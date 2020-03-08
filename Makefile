GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=wingbar-blur

all: test build
build: $(GOBUILD) ./cmd/wingbar-blur -o $(BINARY_NAME) -v
