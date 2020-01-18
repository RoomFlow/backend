PROTO_FILE_PATHS=`find internal/proto -name "*.proto"`

for PROTO_FILE_PATH in $PROTO_FILE_PATHS; do
    protoc -I/usr/local/include -I. \
        -I$GOPATH/src \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --go_out=plugins=grpc:. \
        $PROTO_FILE_PATH --swagger_out=json_names_for_fields=true:.
done
