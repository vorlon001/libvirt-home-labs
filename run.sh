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
#go get iblog.pro/cobra/store
#go get iblog.pro/cobra/logs
#go get iblog.pro/cobra/system
#go get iblog.pro/cobra/core/utils
#go get iblog.pro/cobra/core/interfacedisk
#go get iblog.pro/cobra/core/libvirtvm
#go mod download github.com/grpc-ecosystem/grpc-gateway/v2

go build ./cmd/client.go
go build ./cmd/server.go
