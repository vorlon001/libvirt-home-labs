#!/usr/bin/bash

set -x
function initvm {
./Cobra Configure attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --Octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
./Cobra Configure attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --Octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
./Cobra Configure attachDiskVM --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --Octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
# ./Cobra Configure attachInerfaceVM  --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --Octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
# ./Cobra Configure attachInerfaceVM  --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --Octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml
# ./Cobra Configure attachInerfaceVM  --VMNAME $1 --CORE $2  --MEMORY $3 --ROOTFS_SIZE $4 --Octet $5  --EXT_DISK_SIZE  $6 --USER_DATA_PATH CONFIG/user-data-$7.yaml

}


#initvm node156 5 8 50 156 30 lunar
#exit

MEMORY=16384
CORE=8
ROOTFS_SIZE=60
UBUNNTU=jammy
VMNAME=node
NUMVM=1
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
  node5)
    OCTET=140
    ;;
  node6)
    OCTET=130
    ;;

  *)
    echo -n "unknown"
    exit 1
    ;;
esac

./Cobra Configure initVMs --CORE $CORE --EXT_DISK_SIZE $DISKSIZE --MEMORY $MEMORY --Octet $OCTET --ROOTFS_SIZE $ROOTFS_SIZE --USER_DATA_PATH CONFIG/user-data-$UBUNNTU.yaml  --VMNAME $VMNAME$OCTET  --NumVM $NUMVM

case $(hostname) in
  node1)
	for i in {150..150};
	do
		initvm $VMNAME$i ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} ${UBUNNTU}
	done
    ;;
  node2)
        for i in {160..160};
        do
                initvm $VMNAME$i ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} ${UBUNNTU}
        done
    ;;
  node3)
        for i in {170..170};
        do
                initvm $VMNAME$i ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} ${UBUNNTU}
        done
    ;;
  node4)
        for i in {180..180};
        do
                initvm $VMNAME$i ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} ${UBUNNTU}
        done
    ;;
  node5)
        for i in {140..140};
        do
                initvm $VMNAME$i ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} ${UBUNNTU}
        done
    ;;
  node6)
        for i in {130..130};
        do
                initvm $VMNAME$i ${CORE} ${MEMORY} ${DISKSIZE} ${i} ${DISKSIZE} ${UBUNNTU}
        done
    ;;
  *)

    echo -n "unknown"
    exit 1
    ;;
esac

