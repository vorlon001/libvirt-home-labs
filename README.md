
Version: 31.6.0


# example MachineMigrate

```shell

root@node1:/KVM/init.kvm.v30# ./Cobra help
INFO[2023-07-29T18:39:05+05:00]/KVM/init.kvm.v30/src/VersionBuild.go:8 main.VersionBuild() https://github.com/vorlon001, (C) Vorlon001   Version=31.5.0
INFO[2023-07-29T18:39:05+05:00]/KVM/init.kvm.v30/src/VersionBuild.go:9 main.VersionBuild() HomeLabs LibVirt Connector (Golang version)
Usage:
  HomeLabs [command]

Available Commands:
  Configure   Configure VM command
  LibVirt     LibVirt commands
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for HomeLabs

Use "HomeLabs [command] --help" for more information about a command.
root@node1:/KVM/init.kvm.v30# ./Cobra LibVirt -h
INFO[2023-07-29T18:39:06+05:00]/KVM/init.kvm.v30/src/VersionBuild.go:8 main.VersionBuild() https://github.com/vorlon001, (C) Vorlon001   Version=31.5.0
INFO[2023-07-29T18:39:06+05:00]/KVM/init.kvm.v30/src/VersionBuild.go:9 main.VersionBuild() HomeLabs LibVirt Connector (Golang version)
LibVirt commands

Usage:
  HomeLabs LibVirt [sub] [flags]
  HomeLabs LibVirt [command]

Available Commands:
  MachineCreate     creates a new machine. Requires --xml parameter. Returns result with a current machine state
  MachineDelete     deletes an existing machine.
  MachineDestroy    Destroy an existing machine.
  MachineHardReboot sends a VM into hard-reset mode. This is damaging to all ongoing file operations.
  MachineMigrate    Migrate up a VM. Returns result with a current machine state
  MachinePause      stops the execution of the VM. CPU is not used, but memory is still occupied.
  MachineResume     called after Pause, to resume the invocation of the VM. Returns result with a current machine state
  MachineShutdown   gracefully shuts down the VM. Returns result with a current machine state
  MachineShutoff    kills running VM. Equivalent to pulling a plug out of a computer. Returns result with a current machine state
  MachineSoftReboot reboots a machine gracefully, as chosen by hypervisor. Returns result with a current machine state
  MachineStart      starts up a VM. Returns result with a current machine state
  MachineState      Returns result with a current machine state

Flags:
  -h, --help   help for LibVirt

Use "HomeLabs LibVirt [command] --help" for more information about a command.


root@node1:/KVM/init.kvm.v30# ./Cobra LibVirt MachineMigrate -h
INFO[2023-07-29T18:39:37+05:00]/KVM/init.kvm.v30/src/VersionBuild.go:8 main.VersionBuild() https://github.com/vorlon001, (C) Vorlon001   Version=31.5.0
INFO[2023-07-29T18:39:37+05:00]/KVM/init.kvm.v30/src/VersionBuild.go:9 main.VersionBuild() HomeLabs LibVirt Connector (Golang version)
Migrate up a VM. Returns result with a current machine state

Usage:
  HomeLabs LibVirt MachineMigrate [--id vmname!] [flags]

Flags:
  -h, --help        help for MachineMigrate
      --id string   Libvirt VMname
      --to string   Libvirt Move to HyperVisor
root@node1:/KVM/init.kvm.v30#


root@node3:/KVM/init.kvm.v30# ./Cobra LibVirt MachineMigrate  --id node189 --to 192.168.1.40
INFO[2023-07-29T13:47:39Z]/KVM/init.kvm.v30/src/VersionBuild.go:8 main.VersionBuild() https://github.com/vorlon001, (C) Vorlon001   Version=31.5.0
INFO[2023-07-29T13:47:39Z]/KVM/init.kvm.v30/src/VersionBuild.go:9 main.VersionBuild() HomeLabs LibVirt Connector (Golang version)
INFO[2023-07-29T13:47:39Z]/KVM/init.kvm.v30/src/CobraMenu.go:109 main.(*CobraMenu).RootSubCmdvirtualMachineMigrate.func1() Inside RootSubCmdvirtualMachineStart Run with args  args="[]" core.VMid=node189
INFO[2023-07-29T13:47:39Z]/KVM/init.kvm.v30/src/CobraMenu.go:110 main.(*CobraMenu).RootSubCmdvirtualMachineMigrate.func1() Inside RootSubCmdvirtualMachineStart Run with args  args="[]" core.toMove=192.168.1.40
ID      Name            UUID                            Status
---------------------------------------------------------------------------------------------------------------------------------
libvirt.Domain{Name:"node189", UUID:libvirt.UUID{0x76, 0x34, 0xe2, 0xf4, 0xf0, 0x33, 0x4d, 0x96, 0xb5, 0xf9, 0xe7, 0xa, 0x6e, 0xd6, 0x29, 0xe2}, ID:1}
1       node189 DomainRunning
[]byte(nil)
ID      Name            UUID                            Status
---------------------------------------------------------------------------------------------------------------------------------
libvirt.Domain{Name:"node189", UUID:libvirt.UUID{0x76, 0x34, 0xe2, 0xf4, 0xf0, 0x33, 0x4d, 0x96, 0xb5, 0xf9, 0xe7, 0xa, 0x6e, 0xd6, 0x29, 0xe2}, ID:-1}
-1      node189 7634e2f4f0334d96b5f9e70a6ed629e2        DomainShutoff



root@node4:/KVM/init.kvm.v30# ./Cobra LibVirt MachineMigrate  --id node189 --to 192.168.1.30
INFO[2023-07-29T13:48:55Z]/KVM/init.kvm.v30/src/VersionBuild.go:8 main.VersionBuild() https://github.com/vorlon001, (C) Vorlon001   Version=31.5.0
INFO[2023-07-29T13:48:55Z]/KVM/init.kvm.v30/src/VersionBuild.go:9 main.VersionBuild() HomeLabs LibVirt Connector (Golang version)
INFO[2023-07-29T13:48:55Z]/KVM/init.kvm.v30/src/CobraMenu.go:109 main.(*CobraMenu).RootSubCmdvirtualMachineMigrate.func1() Inside RootSubCmdvirtualMachineStart Run with args  args="[]" core.VMid=node189
INFO[2023-07-29T13:48:55Z]/KVM/init.kvm.v30/src/CobraMenu.go:110 main.(*CobraMenu).RootSubCmdvirtualMachineMigrate.func1() Inside RootSubCmdvirtualMachineStart Run with args  args="[]" core.toMove=192.168.1.30
ID      Name            UUID                            Status
---------------------------------------------------------------------------------------------------------------------------------
libvirt.Domain{Name:"node189", UUID:libvirt.UUID{0x76, 0x34, 0xe2, 0xf4, 0xf0, 0x33, 0x4d, 0x96, 0xb5, 0xf9, 0xe7, 0xa, 0x6e, 0xd6, 0x29, 0xe2}, ID:1}
1       node189 DomainRunning
[]byte(nil)
ID      Name            UUID                            Status
---------------------------------------------------------------------------------------------------------------------------------
libvirt.Domain{Name:"node189", UUID:libvirt.UUID{0x76, 0x34, 0xe2, 0xf4, 0xf0, 0x33, 0x4d, 0x96, 0xb5, 0xf9, 0xe7, 0xa, 0x6e, 0xd6, 0x29, 0xe2}, ID:-1}
-1      node189 7634e2f4f0334d96b5f9e70a6ed629e2        DomainShutoff
root@node4:/KVM/init.kvm.v30#

```

