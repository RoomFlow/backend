PROTO_FILE_PATHS=`find internal/proto -name "*.proto"`

find $GOPATH/src -name "grpc-gateway*" | tail -1

for PROTO_FILE_PATH in $PROTO_FILE_PATHS; do
    protoc -I/usr/local/include -I. \
        -I$GOPATH/src \
        -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --go_out=plugins=grpc:. \
        $PROTO_FILE_PATH --swagger_out=json_names_for_fields=true:.
done
