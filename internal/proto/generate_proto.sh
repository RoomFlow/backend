PROTO_FILE_PATHS=`find internal/proto -name "*.proto"`

for proto_file_path in $PROTO_FILE_PATHS; do
    protoc -I ./ $proto_file_path --go_out=plugins=grpc:./
done
