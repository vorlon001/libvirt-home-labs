#!/usr/bin/bash


go get github.com/digitalocean/go-libvirt
go get github.com/digitalocean/go-libvirt/socket/dialers
go get github.com/sirupsen/logrus
go get github.com/sirupsen/logrus/hooks/syslog
go get github.com/sirupsen/logrus/hooks/writer
go get github.com/spf13/cobra
go get gopkg.in/yaml.v2


go mod tidy -e

CGO_ENABLED=0 go build -ldflags "-w -s -X 'main.Version=30.3.0'" .
cp  main Cobra

