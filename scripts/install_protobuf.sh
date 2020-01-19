PROTOBUF_VERSION=3.11.0

curl -L https://github.com/google/protobuf/releases/download/v{$PROTOBUF_VERSION}/protoc-{$PROTOBUF_VERSION}-linux-x86_64.zip -o /tmp/protoc.zip
unzip /tmp/protoc.zip
sudo cp ./bin/protoc /usr/local/bin/.
sudo cp -r ./include /usr/local/.
sudo chmod 755 /usr/local/bin/protoc
sudo chmod -R 755 /usr/local/include/google
which protoc
protoc --version
