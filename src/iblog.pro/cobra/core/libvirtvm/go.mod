module virsh

go 1.20

replace iblog.pro/cobra/store => ./src/iblog.pro/cobra/store

replace iblog.pro/cobra/logs => ./src/iblog.pro/cobra/logs

replace iblog.pro/cobra/system => ./src/iblog.pro/cobra/system

replace iblog.pro/cobra/core/utils => ./src/iblog.pro/cobra/core/utils

replace iblog.pro/cobra/core/interfacenetwork => ./src/iblog.pro/cobra/core/interfacenetwork

replace iblog.pro/cobra/core/panicrecover => ./src/iblog.pro/cobra/core/panicrecover

replace iblog.pro/cobra/core/interfacedisk => ./src/iblog.pro/cobra/core/interfacedisk

replace iblog.pro/cobra/core => ./src/iblog.pro/cobra/core

replace iblog.pro/cobra/core/model => ./src/iblog.pro/cobra/core/model

replace iblog.pro/cobra/core/virtualmachine => ./src/iblog.pro/cobra/core/virtualmachine

replace iblog.pro/cobra/menu => ./src/iblog.pro/cobra/menu

require (
	github.com/sirupsen/logrus v1.9.3
	gopkg.in/yaml.v2 v2.4.0
	iblog.pro/cobra/core/interfacedisk v0.0.0-00010101000000-000000000000
	iblog.pro/cobra/core/model v0.0.0-00010101000000-000000000000
	iblog.pro/cobra/core/panicrecover v0.0.0-00010101000000-000000000000
	iblog.pro/cobra/core/utils v0.0.0-00010101000000-000000000000
	iblog.pro/cobra/logs v0.0.0-00010101000000-000000000000
	iblog.pro/cobra/menu v0.0.0-00010101000000-000000000000
	iblog.pro/cobra/store v0.0.0-00010101000000-000000000000
	iblog.pro/cobra/system v0.0.0-00010101000000-000000000000
)

require (
	github.com/digitalocean/go-libvirt v0.0.0-20221205150000-2939327a8519 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	iblog.pro/cobra/core/interfacenetwork v0.0.0-00010101000000-000000000000 // indirect
	iblog.pro/cobra/core/virtualmachine v0.0.0-00010101000000-000000000000 // indirect
)
