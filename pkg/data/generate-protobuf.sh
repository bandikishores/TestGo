#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT=$(dirname $(dirname $SCRIPT_DIR))

protoc="protoc
    -I.
    -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.1/third_party/googleapis
    -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.1"

cd $ROOT/data/proto

echo "Generating gRPC server, gateway, swagger"
$protoc --go_out=plugins=grpc:$ROOT/pkg/data \
        --grpc-gateway_out=logtostderr=true,request_context=true:$ROOT/pkg/data \
        --swagger_out=logtostderr=true:$ROOT/data/swagger \
        ./*.proto