#!/bin/bash

# Check if protoc is installed
if ! command -v protoc &> /dev/null; then
    echo "Protobuf is not installed, installing it now..."
    sudo snap install protobuf --classic
fi

# Run protoc commands
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./hello_grpc.protoprotoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./hello_grpc.proto