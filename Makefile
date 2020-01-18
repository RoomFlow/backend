 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

all: build_usermanagement build_apigateway
	

build_usermanagement:
	@echo "Building usermanagement server..."
	$(GOBUILD) services/usermanagement/server/main.go

build_apigateway:
	@echo "Building apigateway server..."
	$(GOBUILD) services/apigateway/cmd/main.go

clean:
	@echo "Cleaning..."
	$(GOCLEAN)

protoc:
	./internal/proto/generate_proto.sh
