#!/usr/bin/env bash

mkdir -p grpc_gen
#protoc --proto_path=devops/idl --go_out=plugins=grpc:grpc_gen devops/idl/comment/info.proto