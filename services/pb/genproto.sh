#!/bin/bash
protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gofast_out=plugins=grpc:. portal.proto