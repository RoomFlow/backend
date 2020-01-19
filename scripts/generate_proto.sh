git clone https://github.com/grpc-ecosystem/grpc-gateway.git
cd grpc-gateway
sudo cp -R ./protoc-gen-swagger /usr/local/include/.
sudo chmod -R 755 /usr/local/include/protoc-gen-swagger

go get github.com/grpc-ecosystem/grpc-gateway
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

GRPC_GATEWAY_PATH=`go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway`
GOOGLE_APIS=$GRPC_GATEWAY_PATH/third_party/googleapis

PROTO_FILE_PATHS=`find internal/proto -name "*.proto"`

for PROTO_FILE_PATH in $PROTO_FILE_PATHS; do
    protoc -I/usr/local/include -I. \
        -I$GOPATH/mod \
        -I$GOOGLE_APIS \
        --go_out=plugins=grpc:. \
        $PROTO_FILE_PATH --swagger_out=json_names_for_fields=true:.
done
