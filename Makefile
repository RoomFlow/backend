 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

build: build_apigateway build_usermanagement build_search build_schedule

buildLinux:
	cd services/apigateway && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./main-alpine
	cd services/search && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./main-alpine
	cd services/usermanagement && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./main-alpine

build_apigateway:
	@echo "Building apigateway binary..."
	$(GOBUILD) services/apigateway/main.go

build_usermanagement:
	@echo "Building usermanagement binary..."
	$(GOBUILD) services/usermanagement/main.go

build_search:
	@echo "Building search binary..."
	$(GOBUILD) services/search/main.go

build_schedule:
	@echo "Building schedule binary..."
	$(GOBUILD) services/schedule/main.go

clean:
	@echo "Cleaning..."
	$(GOCLEAN)

protoc:
	./scripts/generate_proto.sh

protoc_install:
	./scripts/install_protobuf.sh
