#!/usr/bin/bash


# go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
# export PATH="$PATH:$(go env GOPATH)/bin"
# wget https://github.com/protocolbuffers/protobuf/releases/download/v23.4/protoc-23.4-linux-x86_64.zip
# apt  install protobuf-compiler

#export PROTOC_VERSION="23.4"
#export PROTO_OS_VERSION="linux-x86_64"
#wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-$PROTO_OS_VERSION.zip
#unzip protoc-$PROTOC_VERSION-$PROTO_OS_VERSION.zip -d /usr/local/bin

go get  github.com/spf13/afero
go get github.com/envoyproxy/protoc-gen-validate \
    && go get google.golang.org/grpc/cmd/protoc-gen-go-grpc  \
    && go get google.golang.org/protobuf/cmd/protoc-gen-go

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

export PATH="$PATH:$(go env GOPATH)/bin"
/usr/local/bin/bin/protoc $(pwd)/*.proto     \
       --go_out=$(pwd)  \
       --go_opt=paths=source_relative  \
       --go-grpc_out=$(pwd) \
       --go-grpc_opt=paths=source_relative  \
       --proto_path=$(pwd) \
       --openapiv2_out $(pwd) \
       --openapiv2_opt logtostderr=true \
       --openapiv2_opt allow_merge=true \
       --openapiv2_opt merge_file_name=api \
       --openapiv2_opt disable_service_tags=true \
       --openapiv2_opt allow_delete_body=true \
       --openapiv2_opt disable_default_responses=true \
       --openapiv2_opt visibility_restriction_selectors=HIDDEN \
       --openapiv2_opt visibility_restriction_selectors=INTERNAL \
       --openapiv2_opt visibility_restriction_selectors=ADVANCED \
