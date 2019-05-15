GONAME=music-grpc

default: build

init:
	@git submodule init && git submodule update && sh ./devops/grpc_gen.sh

update:
	@git submodule foreach git pull && sh ./devops/grpc_gen.sh

start:


dev:
	@export GO111MODULE=on && go run server.go
