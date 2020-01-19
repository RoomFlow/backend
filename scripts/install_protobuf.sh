PROTOBUF_VERSION=3.11.0

curl -L https://github.com/google/protobuf/releases/download/v{$PROTOBUF_VERSION}/protoc-{$PROTOBUF_VERSION}-linux-x86_64.zip -o /tmp/protoc.zip
unzip /tmp/protoc.zip
sudo cp ./bin/protoc /usr/local/bin/.
sudo cp -r ./include /usr/local/.
sudo chmod 755 /usr/local/bin/protoc
sudo chmod -R 755 /usr/local/include/google
which protoc
protoc --version

git clone https://github.com/grpc-ecosystem/grpc-gateway.git
cd grpc-gateway
sudo cp -R ./protoc-gen-swagger /usr/local/include/.
sudo chmod -R 755 /usr/local/include/protoc-gen-swagger

go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
