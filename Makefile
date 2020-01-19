 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

build: build_usermanagement build_apigateway build_search build_schedule

build_apigateway:
	@echo "Building apigateway server..."
	$(GOBUILD) services/apigateway/cmd/main.go

build_usermanagement:
	@echo "Building usermanagement server..."
	$(GOBUILD) services/usermanagement/cmd/main.go

build_search:
	@echo "Building search server..."
	$(GOBUILD) services/search/cmd/main.go

build_schedule:
	@echo "Building schedule server..."
	$(GOBUILD) services/schedule/cmd/main.go

clean:
	@echo "Cleaning..."
	$(GOCLEAN)

protoc:
	./scripts/generate_proto.sh

protoc_install:
	./scripts/install_protobuf.sh
