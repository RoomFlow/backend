# backend

## Prerequisites

### MacOS
Install Golang: `https://golang.org/dl/`

Install Docker: `https://docs.docker.com/docker-for-mac/install/`

Install Protobuf:
```
PROTOBUF_VERSION=3.11.2
PROTOC_ZIP=protoc-$PROTOBUF_VERSION-osx-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP
```

## Develop locally
1. Run `make protoc` to generate proto files
2. Place firebase auth file in `internal/secrets/firebase-credentials.json`
3. `make buildLinux`
4. `docker-compose up --build`
5. Apigateway is now running on `https://localhost:443`
