#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"

protobase=proto

#protoc/win64/bin/protoc.exe --proto_path=${protobase} --go_out=${protobase} --go_opt=paths=source_relative proto/v3/string/string.proto
#echo "Generated protos"

find ${protobase} -name *.pb.go \
-exec rm {} +
echo "Removed generated code"

find ${protobase} -name *.proto \
-exec protoc/win64/bin/protoc.exe --proto_path=${protobase} --go_out=${protobase} --go_opt=paths=source_relative {} +
echo "Regenated go code"


read -p "Press [Enter] to exit."