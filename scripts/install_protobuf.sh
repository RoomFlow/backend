PROTOBUF_VERSION=3.11.0
PROTOC_DOWNLOAD_FILENAME=protoc-$PROTOBUF_VERSION.zip
PROTOC_DOWNLOAD_LOCATION=/tmp/protocDownload/$PROTOC_DOWNLOAD_FILENAME

# UNAME_S=`uname -s`
# # Only automatically install protoc if on Linux
# if [ "Linux" == "$UNAME_S" ]; then
curl -L https://github.com/google/protobuf/releases/download/v{$PROTOBUF_VERSION}/protoc-{$PROTOBUF_VERSION}-linux-x86_64.zip -o $PROTOC_DOWNLOAD_LOCATION && \
cd $PROTOC_DOWNLOAD_LOCATION && \
unzip $PROTOC_DOWNLOAD_FILENAME && \
sudo cp ./bin/protoc /usr/local/bin/. && \
sudo cp -r ./include /usr/local/. && \
sudo chmod 755 /usr/local/bin/protoc && \
sudo chmod -R 755 /usr/local/include/google && \
protoc --version

rm -rf $PROTOC_DOWNLOAD_LOCATION
# else
#     echo "Please manually install protoc based on your OS from https://github.com/protocolbuffers/protobuf/releases"
# fi
