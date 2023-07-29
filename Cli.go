package main

import (
)


// CGO_ENABLED=0 go build -ldflags "-w -s -X 'main.Version=30.0.0'" .

// ./main LibVirt MachineState --id node190

// ./Cobra  Configure initVM --CORE 4 --EXT_DISK_SIZE 30 --MEMORY 4096 --Octet 180 --ROOTFS_SIZE 30 --USER_DATA_PATH user-data-debian.12.yaml  --VMNAME node180
// ./main Configure initVM --CORE 4 --EXT_DISK_SIZE 30 --MEMORY 4096 --Octet 180 --ROOTFS_SIZE 30 --USER_DATA_PATH user-data-jammy.yaml  --VMNAME node180
// ./main Configure initVMs --CORE 4 --EXT_DISK_SIZE 30 --MEMORY 4096 --Octet 180 --ROOTFS_SIZE 30 --USER_DATA_PATH user-data-jammy.yaml  --VMNAME node180 --NumVM 3
// ./main Configure destroyVM --VMNAME node180

// ./main Configure attachDiskVM --CORE 4 --EXT_DISK_SIZE 30 --MEMORY 4096 --Octet 180 --ROOTFS_SIZE 30 --USER_DATA_PATH user-data-jammy.yaml  --VMNAME node180
// ./main Configure detachDiskVM --VMNAME node180 --DetachDisk sdb --USER_DATA_PATH user-data-jammy.yaml

// ./main Configure attachInerfaceVM --CORE 4 --EXT_DISK_SIZE 30 --MEMORY 4096 --Octet 180 --ROOTFS_SIZE 30 --USER_DATA_PATH user-data-jammy.yaml  --VMNAME node180
// ./main Configure detachInerfaceVM  --USER_DATA_PATH user-data-jammy.yaml  --VMNAME node180 --NetworkSlot "0x04"

// ./main LibVirt MachineDestroy --id node180

// ./main LibVirt MachineCreate --xml  /cloud/KVM/kvm_examples.v20/node180/vm.xml
// ./main LibVirt MachineStart --id node180

// ./main LibVirt MachineShutoff --id node180
// ./main LibVirt MachineDelete --id node180

// https://stackoverflow.com/questions/10200178/call-a-method-from-a-go-template

//***************************************************************************************

// ./main.2 Configure sub1 -c 5 -m 4096


func main() {

        log = InitLogrus()
        defer PanicRecover()

	VersionBuild(Version)

	initMenu().Execute()

}
