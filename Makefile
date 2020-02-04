# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=app
    
all: clean build run
clean: 
	$(GOCLEAN)
build: 
	$(GOBUILD) -o $(BINARY_NAME)
run:
	./$(BINARY_NAME)
