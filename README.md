# Create VM libvirt on python 

V26



pip3 freeze > requirements.txt
pip3 install -r requirements.txt





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





```
apt-get install openvswitch-switch
```


# n1
```
ovs-vsctl del-br sw1
systemctl restart ovs-vswitchd.service
systemctl restart ovsdb-server.service
ovs-vsctl add-br sw1
```

# ovs-vsctl add-port sw1 internalPort1 -- set interface internalPort1 type=internal
# ifconfig internalPort1 12.101.0.1/24  up

```
ovs-vsctl del-port sw1 tun8
ovs-vsctl del-port sw1 tun10

ovs-vsctl add-port sw1 tun8 -- set interface tun8 type=vxlan options:remote_ip=192.168.1.20   options:key=234234 mtu_request=4000
ovs-vsctl add-port sw1 tun10 -- set interface tun10 type=vxlan options:remote_ip=192.168.1.30   options:key=234234 mtu_request=4000

sudo ovs-vsctl del-port sw1 tun8
sudo ovs-vsctl del-port sw1 tun10
ovs-vsctl add-port sw1 tun8 -- set interface tun8 type=geneve options:remote_ip=192.168.1.20   options:key=234234 mtu_request=4000
ovs-vsctl add-port sw1 tun10 -- set interface tun10 type=geneve options:remote_ip=192.168.1.30   options:key=234234 mtu_request=4000


ovs-vsctl add-port sw1 vlan200 tag=200 --\
                set interface vlan200 type=internal

ip addr add 192.168.93.10/24 dev vlan200
ip link set vlan200 up

ovs-vsctl add-port sw1 vlan400 tag=400 --\
                set interface vlan400 type=internal

ip addr add 192.168.94.10/24 dev vlan400
ip link set vlan400 up

ovs-vsctl set Bridge sw1 rstp_enable=true
ovs-vsctl set Bridge sw1 stp_enable=false

ovs-vsctl del-port sw1 vlan600
ovs-vsctl add-port sw1 vlan600 tag=600 --\
                set interface vlan600 type=internal mtu_request=4000


ovs-vsctl del-port sw1 vlan800
ovs-vsctl add-port sw1 vlan800 tag=800 --\
                set interface vlan800 type=internal mtu_request=4000

```





# n2
```
ovs-vsctl del-br sw1
systemctl restart ovs-vswitchd.service
systemctl restart ovsdb-server.service
ovs-vsctl add-br sw1
```
# ovs-vsctl add-port sw1 internalPort1 -- set interface internalPort1 type=internal
# ifconfig internalPort1 12.101.0.2/24  up
```
sudo ovs-vsctl add-port sw1 tun8 -- set interface tun8 type=vxlan options:remote_ip=192.168.1.10  options:key=234234 mtu_request=4000
sudo ovs-vsctl add-port sw1 tun9 -- set interface tun9 type=vxlan options:remote_ip=192.168.1.30  options:key=234234 mtu_request=4000

sudo ovs-vsctl del-port sw1 tun8
sudo ovs-vsctl del-port sw1 tun9
sudo ovs-vsctl add-port sw1 tun8 -- set interface tun8 type=geneve options:remote_ip=192.168.1.10  options:key=234234 mtu_request=4000
sudo ovs-vsctl add-port sw1 tun9 -- set interface tun9 type=geneve options:remote_ip=192.168.1.30  options:key=234234 mtu_request=4000

ovs-vsctl del-port sw1 vlan200
ovs-vsctl add-port sw1 vlan200 tag=200 --\
                set interface vlan200 type=internal mtu_request=4000


ip addr add 192.168.93.20/24 dev vlan200
ip link set vlan200 up

ovs-vsctl del-port sw1 vlan400
ovs-vsctl add-port sw1 vlan400 tag=400 --\
                set interface vlan400 type=internal mtu_request=4000
ovs-vsctl set int sw1 mtu_request=4000

ip addr add 192.168.94.20/24 dev vlan400
ip link set vlan400 up

ping  -M do -s 3900 192.168.200.2


ovs-vsctl set Bridge sw1 rstp_enable=true
ovs-vsctl set Bridge sw1 stp_enable=false

ovs-appctl rstp/show


ovs-vsctl del-port sw1 vlan600
ovs-vsctl add-port sw1 vlan600 tag=600 --\
                set interface vlan600 type=internal mtu_request=4000


ovs-vsctl del-port sw1 vlan800
ovs-vsctl add-port sw1 vlan800 tag=800 --\
                set interface vlan800 type=internal mtu_request=4000

```
