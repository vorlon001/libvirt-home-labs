#!/usr/bin/bash

#!/usr/bin/bash

function initvm {
#python3 ./init-vm-8.py --COMMAND initVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
python3 init-vm-8.py --COMMAND attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
python3 init-vm-8.py --COMMAND attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
python3 init-vm-8.py --COMMAND attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
python3 ./init-vm-8.py --COMMAND listDiskVM --VMNAME $1
python3 ./init-vm-8.py --COMMAND listInterfaceVM --VMNAME $1

}


#initvm node156 5 8 50 156 30 lunar
#exit

MEMORY=8
CORE=5
ROOTFS_SIZE=30
UBUNNTU=focal
VMNAME=node
OCTET=150
NUMVM=3
DISKSIZE=30 # 150..155

python3 ./init-vm-8.py --COMMAND initVMs --numvm ${NUMVM} --VMNAME ${VMNAME} --CORE ${CORE}  --MEMORY ${MEMORY} --ROOTFS_SIZE ${ROOTFS_SIZE} --octet ${OCTET}  --EXT_DISK_SIZE ${DISKSIZE} --USER_DATA_PATH CONFIG/user-data-${UBUNNTU}.yaml

for i in {150..152};
do
initvm node${i} ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} ${UBUNNTU}
done
