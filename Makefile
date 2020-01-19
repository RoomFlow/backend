 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

build: build_usermanagement build_apigateway

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
	./scripts/generate_proto.sh

protoc_install:
	./scripts/install_protobuf.sh
