#!/usr/bin/bash

go get github.com/digitalocean/go-libvirt
go get github.com/digitalocean/go-libvirt/socket/dialers
go get github.com/sirupsen/logrus
go get github.com/sirupsen/logrus/hooks/syslog
go get github.com/sirupsen/logrus/hooks/writer
go get github.com/spf13/cobra
go get gopkg.in/yaml.v2
go get github.com/go-kit/log

go mod tidy -e

# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0

/root/go/bin/golangci-lint run

CGO_ENABLED=0 go build -ldflags "-w -s -X 'main.Version=35.0.2'"  -o ./server.grpc ./cmd.grpc/server.go
