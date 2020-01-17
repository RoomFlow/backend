 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

all: 
	build-usermanagement
	build-apigateway
	

build-usermanagement:
	$(GOBUILD) services/usermanagement/server/main.go

build-apigateway:
	$(GOBUILD) services/apigateway/cmd/main.go

clean:
	$(GOCLEAN)
