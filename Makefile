GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u
BINARY_NAME=pprofit
WEBSERVER_ADDR=http://localhost:8080

all: test build
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
test: 
		$(GOTEST) -v ./...
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
stop: 
		killall $(BINARY_NAME)
deps:
		$(GOGET) github.com/tsliwowicz/go-wrk
perf: 
		go-wrk -d 60 $(WEBSERVER_ADDR)/dennis@golang.org
	
