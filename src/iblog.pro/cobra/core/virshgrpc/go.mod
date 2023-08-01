module main

go 1.20

replace iblog.pro/condor/protos => ../protoc

replace iblog.pro/cobra/store => ../iblog.pro/cobra/store

replace iblog.pro/cobra/logs => ../iblog.pro/cobra/logs

replace iblog.pro/cobra/system => ../iblog.pro/cobra/system

replace iblog.pro/cobra/core/utils => ../iblog.pro/cobra/core/utils

replace iblog.pro/cobra/core/interfacenetwork => ../iblog.pro/cobra/core/interfacenetwork

replace iblog.pro/cobra/core/panicrecover => ../iblog.pro/cobra/core/panicrecover

replace iblog.pro/cobra/core/interfacedisk => ../iblog.pro/cobra/core/interfacedisk

replace iblog.pro/cobra/core => ../iblog.pro/cobra/core

replace iblog.pro/cobra/core/model => ../iblog.pro/cobra/core/model

replace iblog.pro/cobra/core/virtualmachine => ../iblog.pro/cobra/core/virtualmachine

replace iblog.pro/cobra/menu => ../iblog.pro/cobra/menu

replace iblog.pro/cobra/core/libvirtvm => ../iblog.pro/cobra/core/libvirtvm

require (
	iblog.pro/cobra/core/model v0.0.0-00010101000000-000000000000
	iblog.pro/cobra/core/panicrecover v0.0.0-00010101000000-000000000000
	iblog.pro/cobra/menu v0.0.0-00010101000000-000000000000
	iblog.pro/cobra/system v0.0.0-00010101000000-000000000000
)

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.2
	google.golang.org/grpc v1.57.0
	iblog.pro/condor/protos v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20230711160842-782d3b101e98 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230726155614-23370e0ffb3e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230726155614-23370e0ffb3e // indirect
	google.golang.org/grpc/examples v0.0.0-20230801011711-02463732635a // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

require (
	github.com/digitalocean/go-libvirt v0.0.0-20221205150000-2939327a8519 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/cobra v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	iblog.pro/cobra/core/interfacedisk v0.0.0-00010101000000-000000000000 // indirect
	iblog.pro/cobra/core/interfacenetwork v0.0.0-00010101000000-000000000000 // indirect
	iblog.pro/cobra/core/libvirtvm v0.0.0-00010101000000-000000000000 // indirect
	iblog.pro/cobra/core/utils v0.0.0-00010101000000-000000000000 // indirect
	iblog.pro/cobra/core/virtualmachine v0.0.0-00010101000000-000000000000 // indirect
	iblog.pro/cobra/logs v0.0.0-00010101000000-000000000000 // indirect
	iblog.pro/cobra/store v0.0.0-00010101000000-000000000000 // indirect
)
