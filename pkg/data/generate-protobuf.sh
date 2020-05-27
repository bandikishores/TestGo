#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT=$(dirname $(dirname $SCRIPT_DIR))
GOGOPATH=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1

protoc="protoc
    -I.
    -I$GOPATH/pkg/mod/cache/download/github.com/gogo/googleapis
    -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.5/third_party/googleapis
    -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.5
    -I$GOPATH/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20200303215514-541b1ded4aba/
    -I$GOGOPATH/protobuf
    -I$GOGOPATH
    -I$GOPATH/pkg/mod"

cd $ROOT/data/proto

echo "Generating gRPC Server, gateway"
$protoc --gogo_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:$ROOT/pkg/data \
        --grpc-gateway_out=logtostderr=true,allow_patch_feature=false,request_context=true:$ROOT/pkg/data \
        ./*.proto

echo "Injecting Custom Tags in Fields"
ls $ROOT/pkg/data/*.pb.go | while read file; do protoc-go-inject-tag -input=$file; done


echo "Generating swagger"
$protoc --swagger_out=logtostderr=true:$ROOT/data/swagger \
        --swagger_out=logtostderr=true,allow_merge=true:$ROOT/data/swagger \
        ./*.proto