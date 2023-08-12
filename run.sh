#!/usr/bin/bash

#go get google.golang.org/grpc
#go get google.golang.org/grpc/credentials
#go get google.golang.org/grpc/credentials/insecure
#go get google.golang.org/grpc/examples/data
#go get iblog.pro/condor/protos
#go get github.com/digitalocean/go-libvirt
#go get github.com/digitalocean/go-libvirt/socket/dialers
#go get github.com/sirupsen/logrus
#go get github.com/sirupsen/logrus/hooks/syslog
#go get github.com/sirupsen/logrus/hooks/writer
#go get github.com/spf13/cobra
#go get gopkg.in/yaml.v2
#go get go.uber.org/zap
#go mod download golang.org/x/sync
#go mod download go.uber.org/zap
#go mod tidy
#go mod download github.com/grpc-ecosystem/grpc-gateway/v2

go build -o ./bin/client ./cmd/client.go
go build -o ./bin/server ./cmd/server.go
