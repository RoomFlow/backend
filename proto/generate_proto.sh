echo $PWD
echo $GOPATH

protoc -I. -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. \
proto/usermanagement/user_management.proto --swagger_out=json_names_for_fields=true:.