#!/usr/bin/bash

export GOLANG_VERSION="1.21.4"

go get github.com/digitalocean/go-libvirt
go get github.com/digitalocean/go-libvirt/socket/dialers
go get github.com/sirupsen/logrus
go get github.com/sirupsen/logrus/hooks/syslog
go get github.com/sirupsen/logrus/hooks/writer
go get github.com/spf13/cobra
go get gopkg.in/yaml.v2

go mod tidy -e

CGO_ENABLED=0 go build -ldflags "-w -s -X 'main.Version=34.2.0'" -o ../Cobra ./cli
CGO_ENABLED=0 go build -ldflags "-w -s -X 'main.Version=34.2.0'" -o ./client.grpc ./cmd.grpc/client.go
CGO_ENABLED=0 go build -ldflags "-w -s -X 'main.Version=34.2.0'"  -o ./server.grpc ./cmd.grpc/server.go
