# Create VM libvirt on python 

V20



need create libvirt network
need wget jammy cloud image

python3 run.sh

del.vm.sh


```shell
python3 run.vm.py --COMMAND initVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
python3 run.vm.py --COMMAND attachInerfaceVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20
python3 run.vm.py --COMMAND attachInerfaceVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
python3 run.vm.py --COMMAND attachDiskVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
```

```shell

## create network on openvswitch

python3 ./run.vm.py  --COMMAND createNetworkVlan --virtual_network_bridge_name=sw1 --virtual_network_vlan_id=700
python3 ./run.vm.py  --COMMAND destroyNetworkVlan --virtual_network_bridge_name=sw1 --virtual_network_vlan_id=700
python3 ./run.vm.py  --COMMAND createNetworkTrunk --virtual_network_bridge_name=sw1
python3 ./run.vm.py  --COMMAND destroyNetworkTrunk --virtual_network_bridge_name=sw1


## init vm and attach disk and network interface
python3 run.vm.py --COMMAND initVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
python3 run.vm.py --COMMAND attachDiskVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20
python3 run.vm.py --COMMAND attachInerfaceVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20


python3 run.vm.py --COMMAND initVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
python3 run.vm.py --COMMAND attachInerfaceVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20
python3 run.vm.py --COMMAND attachInerfaceVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
python3 run.vm.py --COMMAND attachDiskVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
python3 run.vm.py --COMMAND destroyVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20

## view disk in vm
python3 ./run.vm.py --COMMAND  listDiskVM --VMNAME node170 --CORE 4  --MEMORY 4 --ROOTFS_SIZE 30 --octet 170  --EXT_DISK_SIZE  20
##### [2022-05-02:20:08:44] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Nefelim.py:49, message: Nefelim().__init__()
##### Namespace(COMMAND='listDiskVM', VMNAME='node170', CORE=4, MEMORY=4, ROOTFS_SIZE=30, octet=170, EXT_DISK_SIZE=20, USER_DATA_PATH='CONFIG/user-data.yaml')
##### {'COMMAND': 'listDiskVM', 'VMNAME': 'node170', 'CORE': 4, 'MEMORY': 4, 'ROOTFS_SIZE': 30, 'octet': 170, 'EXT_DISK_SIZE': 20, 'USER_DATA_PATH': 'CONFIG/user-data.yaml'}
##### [2022-05-02:20:08:44] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Genesys.py:73, message: Genesys().initConfig()
##### [2022-05-02:20:08:44] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Genesys.py:80, message: Genesys().setVarConfig()
##### {'source': '/cloud/TEST.1/KVM/node170/node170.qcow2', 'DEVICE': 'sda', 'BUS': 'scsi'}

## view network interface in vm
python3 ./run.vm.py --COMMAND listInterfaceVM --VMNAME node170 --CORE 4  --MEMORY 4 --ROOTFS_SIZE 30 --octet 170  --EXT_DISK_SIZE  20
##### [2022-05-02:20:08:45] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Nefelim.py:49, message: Nefelim().__init__()
##### Namespace(COMMAND='listInterfaceVM', VMNAME='node170', CORE=4, MEMORY=4, ROOTFS_SIZE=30, octet=170, EXT_DISK_SIZE=20, USER_DATA_PATH='CONFIG/user-data.yaml')
##### {'COMMAND': 'listInterfaceVM', 'VMNAME': 'node170', 'CORE': 4, 'MEMORY': 4, 'ROOTFS_SIZE': 30, 'octet': 170, 'EXT_DISK_SIZE': 20, 'USER_DATA_PATH': 'CONFIG/user-data.yaml'}
##### [2022-05-02:20:08:45] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Genesys.py:73, message: Genesys().initConfig()
##### [2022-05-02:20:08:45] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Genesys.py:80, message: Genesys().setVarConfig()
##### {'address': 'pci', 'network': 'cloud_sw1', 'interfaceid': '8d2b612f-017e-4f35-b6e1-75bbdcbce43b', 'target': 'vnet238', 'domain': '0x0000', 'bus': '0x01', 'slot': '0x00', 'function': '0x0', 'multifunction': 'on
'}
##### {'address': 'pci', 'network': 'cloud_sw1', 'interfaceid': 'ab0fc13f-8ea0-4837-a813-0fb2a16a69c0', 'target': 'vnet239', 'domain': '0x0000', 'bus': '0x03', 'slot': '0x01', 'function': '0x0', 'multifunction': 'on
'}

```

