 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

all: build

build:
	$(GOBUILD) service/usermanagement/server/main.go

clean:
	$(GOCLEAN)