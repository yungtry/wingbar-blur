GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=wingbar-blur

build: 
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/wingbar-blur 