#!/bin/bash

protoc --go_out=../api/ --go-grpc_out=../api/ main.proto