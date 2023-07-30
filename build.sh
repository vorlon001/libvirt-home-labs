#!/usr/bin/bash

export GOLANG_VERSION="1.20.6"
export GOPROXY='https://nexus3.iblog.pro/repository/go-proxy/'
export GONOSUMDB="https://gitlab.iblog.pro/*"
export GONOPROXY="https://gitlab.iblog.pro/*"
#export GOSUMDB='sum.golang.org https://nexus.iblog.pro/repository/golang-sum/'

go get github.com/digitalocean/go-libvirt
go get github.com/digitalocean/go-libvirt/socket/dialers
go get github.com/sirupsen/logrus
go get github.com/sirupsen/logrus/hooks/syslog
go get github.com/sirupsen/logrus/hooks/writer
go get github.com/spf13/cobra
go get gopkg.in/yaml.v2
go get iblog.pro/cobra/store
go get iblog.pro/cobra/logs
go get iblog.pro/cobra/system
go get iblog.pro/cobra/core/utils
go get iblog.pro/cobra/core/interfacedisk

go mod tidy -e

CGO_ENABLED=0 go build -ldflags "-w -s -X 'main.Version=32.8.0'" .
cp  cobra Cobra
cp Cobra ..