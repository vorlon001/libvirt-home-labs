#!/usr/bin/bash

function initvm {

python3 ./run.vm.py --COMMAND initVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
python3 run.vm.py --COMMAND attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
python3 run.vm.py --COMMAND attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
python3 run.vm.py --COMMAND attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
python3 ./run.vm.py --COMMAND listDiskVM --VMNAME $1
python3 ./run.vm.py --COMMAND listInterfaceVM --VMNAME $1

}


#initvm node156 5 8 50 156 30 jammy
#exit

MEMORY=8
CORE=5
DISKSIZE=30 # 150..155
for i in {150..152};
do
initvm node${i} ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} focal
done
