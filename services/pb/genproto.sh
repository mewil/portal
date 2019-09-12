#!/bin/bash
protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gofast_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc:. *.proto