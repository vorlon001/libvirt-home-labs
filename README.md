# Create VM libvirt on python 

V20


need create libvirt network
need wget jammy cloud image

python3 run.sh

del.vm.sh


```shell
python3 init-vm-8.py --COMMAND initVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
python3 init-vm-8.py --COMMAND attachInerfaceVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20
python3 init-vm-8.py --COMMAND attachInerfaceVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
python3 init-vm-8.py --COMMAND attachDiskVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
```
