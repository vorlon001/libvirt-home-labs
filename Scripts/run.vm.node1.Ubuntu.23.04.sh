#!/usr/bin/bash

set -x
function initvm {
./Cobra Configure attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --Octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
./Cobra Configure attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --Octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
./Cobra Configure attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --Octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
}


#initvm node156 5 8 50 156 30 lunar
#exit

MEMORY=8192
CORE=5
ROOTFS_SIZE=30
UBUNNTU=lunar
VMNAME=node
NUMVM=3
DISKSIZE=30

case $(hostname) in

  node1)
    OCTET=150
    ;;

  node2)
    OCTET=160
    ;;

  node3)
    OCTET=170
    ;;

  node4)
    OCTET=180
    ;;
  *)
    echo -n "unknown"
    exit 1
    ;;
esac

./Cobra Configure initVMs --CORE $CORE --EXT_DISK_SIZE $DISKSIZE --MEMORY $MEMORY --Octet $OCTET --ROOTFS_SIZE $ROOTFS_SIZE --USER_DATA_PATH CONFIG/user-data-$UBUNNTU.yaml  --VMNAME $VMNAME$OCTET  --NumVM $NUMVM

case $(hostname) in
  node1)
	for i in {150..152};
	do
		initvm $VMNAME$i ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} ${UBUNNTU}
	done
    ;;
  node2)
        for i in {160..162};
        do
                initvm $VMNAME$i ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} ${UBUNNTU}
        done
    ;;
  node3)
        for i in {170..172};
        do
                initvm $VMNAME$i ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} ${UBUNNTU}
        done
    ;;
  node4)
        for i in {180..182};
        do
                initvm $VMNAME$i ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} ${UBUNNTU}
        done
    ;;
  *)
    echo -n "unknown"
    exit 1
    ;;
esac

