#!/usr/bin/bash

function initvm {

python3 ./run.vm.py --COMMAND initVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data.debian.11.yaml
python3 run.vm.py --COMMAND attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data.debian.11.yaml
python3 run.vm.py --COMMAND attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data.debian.11.yaml
python3 run.vm.py --COMMAND attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data.debian.11.yaml
python3 run.vm.py --COMMAND attachInerfaceVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data.debian.11.yaml
python3 run.vm.py --COMMAND attachInerfaceVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data.debian.11.yaml

python3 ./run.vm.py --COMMAND listDiskVM --VMNAME $1
python3 ./run.vm.py --COMMAND listInterfaceVM --VMNAME $1

}

#initvm node81 5 16 30 81 30

initvm node150 5 8 30 150 30
initvm node151 5 8 30 151 30
initvm node152 5 8 30 152 30
