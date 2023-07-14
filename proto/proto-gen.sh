#!/bin/bash

protoc --go_out=../api/ --go-grpc_out=../api/  main.proto ||
npx ../src/proto-loader-gen-types --grpcLib=../src/@grpc/grpc-js --outDir=../src/pb/ *.proto