PROTO_FILE_PATHS=`find internal/proto -name "*.proto"`

go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway

GRPC_GATEWAY_PATH=`find $GOPATH/pkg/mod/github.com/grpc-ecosystem -name "grpc-gateway*" | tail -1`
GOOGLE_APIS=$GRPC_GATEWAY_PATH/third_party/googleapis

for PROTO_FILE_PATH in $PROTO_FILE_PATHS; do
    protoc -I/usr/local/include -I. \
        -I$GOPATH/mod \
        -I$GOOGLE_APIS \
        --go_out=plugins=grpc:. \
        $PROTO_FILE_PATH --swagger_out=json_names_for_fields=true:.
done
