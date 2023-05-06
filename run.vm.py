try:
    import CORE.Core.LOADER as boot
    boot.lm(globals(),
				"CORE.LibVirt.Nefelim", "sys", "libvirt", "subprocess", "argparse", "asyncio", "copy", "multiprocessing"
				)


#from multiprocessing import Process

    boot.iglob(globals(),[
                            { "module": "multiprocessing",   "method":"Process",     "as": "Process"    },
                            { "module": "CORE.LibVirt.Nefelim",   "method":"Nefelim",     "as": "Nefelim"    }
                          ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)

#python3 ./init-vm-8.py  --COMMAND createNetworkVlan --virtual_network_bridge_name=sw1 --virtual_network_vlan_id=700
#python3 ./init-vm-8.py  --COMMAND destroyNetworkVlan --virtual_network_bridge_name=sw1 --virtual_network_vlan_id=700
#python3 ./init-vm-8.py  --COMMAND createNetworkTrunk --virtual_network_bridge_name=sw1
#python3 ./init-vm-8.py  --COMMAND destroyNetworkTrunk --virtual_network_bridge_name=sw1



# python3 init-vm-8.py --COMMAND initVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND initVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND attachDiskVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND attachInerfaceVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20


# python3 init-vm-8.py --COMMAND initVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND attachInerfaceVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND attachInerfaceVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND attachDiskVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND destroyVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20

# python3 ./init-vm-8.py --COMMAND  listDiskVM --VMNAME node170 --CORE 4  --MEMORY 4 --ROOTFS_SIZE 30 --octet 170  --EXT_DISK_SIZE  20
##### [2022-05-02:20:08:44] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Nefelim.py:49, message: Nefelim().__init__()
##### Namespace(COMMAND='listDiskVM', VMNAME='node170', CORE=4, MEMORY=4, ROOTFS_SIZE=30, octet=170, EXT_DISK_SIZE=20, USER_DATA_PATH='CONFIG/user-data.yaml')
##### {'COMMAND': 'listDiskVM', 'VMNAME': 'node170', 'CORE': 4, 'MEMORY': 4, 'ROOTFS_SIZE': 30, 'octet': 170, 'EXT_DISK_SIZE': 20, 'USER_DATA_PATH': 'CONFIG/user-data.yaml'}
##### [2022-05-02:20:08:44] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Genesys.py:73, message: Genesys().initConfig()
##### [2022-05-02:20:08:44] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Genesys.py:80, message: Genesys().setVarConfig()
##### {'source': '/cloud/TEST.1/KVM/node170/node170.qcow2', 'DEVICE': 'sda', 'BUS': 'scsi'}

# python3 ./init-vm-8.py --COMMAND listInterfaceVM --VMNAME node170 --CORE 4  --MEMORY 4 --ROOTFS_SIZE 30 --octet 170  --EXT_DISK_SIZE  20
##### [2022-05-02:20:08:45] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Nefelim.py:49, message: Nefelim().__init__()
##### Namespace(COMMAND='listInterfaceVM', VMNAME='node170', CORE=4, MEMORY=4, ROOTFS_SIZE=30, octet=170, EXT_DISK_SIZE=20, USER_DATA_PATH='CONFIG/user-data.yaml')
##### {'COMMAND': 'listInterfaceVM', 'VMNAME': 'node170', 'CORE': 4, 'MEMORY': 4, 'ROOTFS_SIZE': 30, 'octet': 170, 'EXT_DISK_SIZE': 20, 'USER_DATA_PATH': 'CONFIG/user-data.yaml'}
##### [2022-05-02:20:08:45] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Genesys.py:73, message: Genesys().initConfig()
##### [2022-05-02:20:08:45] INFO [Logger.py.log:54] LOGGER: uuid: /KVM/init.kvm.v20/CORE/LibVirt/Genesys.py:80, message: Genesys().setVarConfig()
##### {'address': 'pci', 'network': 'cloud_sw1', 'interfaceid': '8d2b612f-017e-4f35-b6e1-75bbdcbce43b', 'target': 'vnet238', 'domain': '0x0000', 'bus': '0x01', 'slot': '0x00', 'function': '0x0', 'multifunction': 'on'}
##### {'address': 'pci', 'network': 'cloud_sw1', 'interfaceid': 'ab0fc13f-8ea0-4837-a813-0fb2a16a69c0', 'target': 'vnet239', 'domain': '0x0000', 'bus': '0x03', 'slot': '0x01', 'function': '0x0', 'multifunction': 'on'}

def initVM(vm: dict):


    del vm['COMMAND']
    del vm['virtual_network_vlan_id']
    del vm['virtual_network_bridge_name']
    del vm["numvm"]

    nefelim = Nefelim()
    nefelim.initVM( **vm )


def initVMs(vm: dict):

    procs = []
    proc = Process(target=initVM)  # instantiating without any argument
    procs.append(proc)
    proc.start()

    for i in range(vm["octet"],vm["octet"]+vm["numvm"]):
        vmTmp = copy.deepcopy(vm)
        vmTmp["VMNAME"] = f"{vm['VMNAME']}{i}"
        vmTmp["octet"] = i
        print(vmTmp)

        # print(name)
        proc = Process(target=initVM, args=(vmTmp,))
        procs.append(proc)
        proc.start()

    # complete the processes
    for proc in procs:
        proc.join()


def attachInerfaceVM(vm: dict):

    print(vm)

    del vm['COMMAND']
    del vm['virtual_network_vlan_id']
    del vm['virtual_network_bridge_name']
    del vm["numvm"]

    nefelim = Nefelim()
    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()


    VMNAME=vm["VMNAME"]
    dom = nefelim.conn.lookupByName(VMNAME)
    docj = nefelim.vmGetJsonConfig(nefelim.conn,VMNAME)
    docj = docj['result']


    slot_id_new_interface_vm=nefelim.vmGetIdNewInterface(docj)['result']
    print(slot_id_new_interface_vm)

    print( ">>>>>", nefelim.network)
    print( "<<<<<", nefelim.random_mac(nefelim.network))

    newMac = nefelim.random_mac(nefelim.network)["result"]

    disk_seed=f"""
    <interface type='network'>
      <source network='cloud_sw1'/>
      <model type='e1000'/>
      <mac address='{newMac}'/>
      <address type='pci' domain='0x0000' bus='0x03' slot='{slot_id_new_interface_vm}' function='0x0' multifunction='on'/>
    </interface>
"""
    print(disk_seed)

    dom.attachDeviceFlags( disk_seed, flags = libvirt.VIR_DOMAIN_AFFECT_CURRENT | libvirt.VIR_DOMAIN_AFFECT_CONFIG | libvirt.VIR_DOMAIN_AFFECT_LIVE)


def attachDiskVM(vm: dict):

    del vm['COMMAND']
    del vm['virtual_network_vlan_id']
    del vm['virtual_network_bridge_name']
    del vm["numvm"]

    nefelim = Nefelim()
    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()

    VMNAME=vm["VMNAME"]

    dom = nefelim.conn.lookupByName(VMNAME)
    docj = nefelim.vmGetJsonConfig(nefelim.conn,VMNAME)
    docj = docj['result']

    DISKID=nefelim.vmGetIdNewDisk(docj)['result']
    DISKSIZE=vm["EXT_DISK_SIZE"]
    cmd=f"qemu-img create -f qcow2 {nefelim.VMPATH}/{VMNAME}/{VMNAME}-disk-{DISKID}.qcow2 {DISKSIZE}G"
    value=subprocess.call(cmd, shell=True)
    print(value)

    disk_seed=f"""
<disk type="file" device="disk">
        <driver name='qemu' type='qcow2'/>
        <source file="{nefelim.VMPATH}/{nefelim.VMNAME}/{nefelim.VMNAME}-disk-{DISKID}.qcow2"></source>
        <target dev='{DISKID}' bus='scsi'/>
</disk>
"""
    print(disk_seed)

    dom.attachDeviceFlags( disk_seed, flags = libvirt.VIR_DOMAIN_AFFECT_CURRENT | libvirt.VIR_DOMAIN_AFFECT_CONFIG | libvirt.VIR_DOMAIN_AFFECT_LIVE)



def destroyVM(vm: dict):

    del vm['COMMAND']
    del vm['virtual_network_vlan_id']
    del vm['virtual_network_bridge_name']
    del vm["numvm"]

    nefelim = Nefelim()
    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()

    VMNAME=vm["VMNAME"]

    dom = nefelim.conn.lookupByName(VMNAME)
    docj = nefelim.vmGetJsonConfig(nefelim.conn,VMNAME)
    docj = docj['result']


    dom.destroy()
    dom.undefine()

    cmd=f"rm -R {nefelim.VMPATH}/{VMNAME}"
    value=subprocess.call(cmd, shell=True)
    print(value)


def listDiskVM(vm: dict):

    del vm['COMMAND']
    del vm['virtual_network_vlan_id']
    del vm['virtual_network_bridge_name']
    del vm["numvm"]

    nefelim = Nefelim()
    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()

    VMNAME=vm["VMNAME"]

    dom = nefelim.conn.lookupByName(VMNAME)
    docj = nefelim.vmGetJsonConfig(nefelim.conn,VMNAME)
    docj = docj['result']

    for i in nefelim.vmGetDiskJsonConfig(docj)['result']:
        print(i)


def listInterfaceVM(vm: dict):

    del vm['COMMAND']
    del vm['virtual_network_vlan_id']
    del vm['virtual_network_bridge_name']
    del vm["numvm"]

    nefelim = Nefelim()
    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()

    VMNAME=vm["VMNAME"]

    dom = nefelim.conn.lookupByName(VMNAME)
    docj = nefelim.vmGetJsonConfig(nefelim.conn,VMNAME)
    docj = docj['result']

    for i in nefelim.vmGetInterface(docj)['result']:
        print(i)


def createNetworkVlan(vm: dict):
    del vm['COMMAND']
    del vm["numvm"]

    nefelim = Nefelim()
    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()

    conn = nefelim.conn

    virtual_network_vlan_id = vm['virtual_network_vlan_id']
    virtual_network_name = f"cloud_vlan{virtual_network_vlan_id}"
    virtual_network_bridge_name = vm['virtual_network_bridge_name']

    virtual_network_xml = f"""
<network>
  <name>{virtual_network_name}</name>
  <forward mode='bridge'/>
  <bridge name='{virtual_network_bridge_name}'/>
  <virtualport type='openvswitch'/>
  <portgroup name='vlan{virtual_network_vlan_id}' default='yes'>
    <vlan>
      <tag id='{virtual_network_vlan_id}'/>
    </vlan>
  </portgroup>
</network>
"""
    print(virtual_network_xml)

    conn.networkDefineXML(virtual_network_xml)
    conn.networkLookupByName(virtual_network_name).autostart()
    conn.networkLookupByName(virtual_network_name).create()
    conn.networkLookupByName(virtual_network_name).setAutostart(True)

    print(conn.networkLookupByName(virtual_network_name).autostart())
    print(conn.networkLookupByName(virtual_network_name).UUIDString())
    print(conn.networkLookupByName(virtual_network_name).bridgeName())
    print(conn.networkLookupByName(virtual_network_name).isActive())
    print(conn.listNetworks())



def destroyNetworkVlan(vm: dict):

    del vm['COMMAND']
    del vm["numvm"]

    nefelim = Nefelim()
    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()

    conn = nefelim.conn

    virtual_network_vlan_id = vm['virtual_network_vlan_id']
    virtual_network_name = f"cloud_vlan{virtual_network_vlan_id}"
    virtual_network_bridge_name = vm['virtual_network_bridge_name']

    virtual_network_xml = f"""
<network>
  <name>{virtual_network_name}</name>
  <forward mode='bridge'/>
  <bridge name='{virtual_network_bridge_name}'/>
  <virtualport type='openvswitch'/>
  <portgroup name='vlan{virtual_network_vlan_id}' default='yes'>
    <vlan>
      <tag id='{virtual_network_vlan_id}'/>
    </vlan>
  </portgroup>
</network>
"""
    print(virtual_network_xml)

    conn.networkLookupByName(virtual_network_name).destroy()
    conn.networkLookupByName(virtual_network_name).undefine()

    print(conn.listNetworks())



def createNetworkTrunk(vm: dict):
    del vm['COMMAND']
    del vm["numvm"]

    nefelim = Nefelim()
    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()

    conn = nefelim.conn

    virtual_network_bridge_name = vm['virtual_network_bridge_name']
    virtual_network_name = f"cloud_trunk_{virtual_network_bridge_name}"

    virtual_network_xml = f"""
<network>
 <name>{virtual_network_name}</name>
 <forward mode='bridge'/>
 <bridge name='{virtual_network_bridge_name}'/>
<virtualport type='openvswitch'/>
</network>
"""
    print(virtual_network_xml)


    conn.networkDefineXML(virtual_network_xml)
    conn.networkLookupByName(virtual_network_name).autostart()
    conn.networkLookupByName(virtual_network_name).create()
    conn.networkLookupByName(virtual_network_name).setAutostart(True)

    print(conn.networkLookupByName(virtual_network_name).autostart())
    print(conn.networkLookupByName(virtual_network_name).UUIDString())
    print(conn.networkLookupByName(virtual_network_name).bridgeName())
    print(conn.networkLookupByName(virtual_network_name).isActive())
    print(conn.listNetworks())


def destroyNetworkTrunk(vm: dict):

    del vm['COMMAND']
    del vm["numvm"]

    nefelim = Nefelim()
    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()

    conn = nefelim.conn

    virtual_network_bridge_name = vm['virtual_network_bridge_name']
    virtual_network_name = f"cloud_trunk_{virtual_network_bridge_name}"

    virtual_network_xml = f"""
<network>
 <name>{virtual_network_name}</name>
 <forward mode='bridge'/>
 <bridge name='{virtual_network_bridge_name}'/>
<virtualport type='openvswitch'/>
</network>
"""
    print(virtual_network_xml)

    conn.networkLookupByName(virtual_network_name).destroy()
    conn.networkLookupByName(virtual_network_name).undefine()

    print(conn.listNetworks())


def getArgs() -> dict:
    parser = argparse.ArgumentParser()
    parser.add_argument('--COMMAND', help='COMMAND initVM/attachDiskVM/attachInerfaceVM/destroyVM/listInterfaceVM/listDiskVM/initVMs', default='initVM', type=str)
    parser.add_argument('--VMNAME', help='VMNAME for COMMAND initVM/attachDiskVM/attachInerfaceVM/destroyVM/listInterfaceVM/listDiskVM', default='node100', type=str)
    parser.add_argument('--CORE', help='CORE for COMMAND initVM/attachDiskVM/attachInerfaceVM/destroyVM', default='4', type=int)
    parser.add_argument('--MEMORY', help='MEMORY for COMMAND initVM/attachDiskVM/attachInerfaceVM/destroyVM', default='8', type=int)
    parser.add_argument('--ROOTFS_SIZE', help='ROOTFS_SIZE for COMMAND initVM/attachDiskVM/attachInerfaceVM/destroyVM', default='20', type=int)
    parser.add_argument('--octet', help='octet for COMMAND initVM/attachDiskVM/attachInerfaceVM/destroyVM', default='100', type=int)
    parser.add_argument('--EXT_DISK_SIZE', help='ext-disk-size for COMMAND initVM/attachDiskVM/attachInerfaceVM/destroyVM', default='20', type=int)
    parser.add_argument('--USER_DATA_PATH', help='USER_DATA_PATH for COMMAND initVM/attachDiskVM/attachInerfaceVM/destroyVM', default="CONFIG/user-data.yaml", type=str)
    parser.add_argument('--numvm', help='num vm, default 1', default='1', type=int)
    parser.add_argument('--virtual_network_vlan_id', help='virtual_network_vlan_id for COMMAND createNetworkVlan/destroyNetworkVlan/createNetworkTrunk/destroyNetworkTrunk', default=0, type=int)
    parser.add_argument('--virtual_network_bridge_name', help='virtual_network_bridge_name for COMMAND createNetworkVlan/destroyNetworkVlan/createNetworkTrunk/destroyNetworkTrunk', default="sw1", type=str)

    args = parser.parse_args()
    print(args)
    print(args.__dict__)
    return vars(args)


def main():

#    globals()["logger"] = initLog(loggingLevel = logging.DEBUG, exp = __name__)
#    print(sys.modules['__main__'].__dict__["logger"])

    cmd = {
              'initVMs':initVMs, 'initVM': initVM, "destroyVM": destroyVM,
              'attachDiskVM': attachDiskVM,
              'attachInerfaceVM': attachInerfaceVM,
              "listDiskVM": listDiskVM, "listInterfaceVM": listInterfaceVM,
              "createNetworkVlan": createNetworkVlan,   "destroyNetworkVlan": destroyNetworkVlan,
              "createNetworkTrunk": createNetworkTrunk, "destroyNetworkTrunk": destroyNetworkTrunk,
              "createNetworkTrunk": createNetworkTrunk, "destroyNetworkTrunk": destroyNetworkTrunk,
              "createNetworkVlan": createNetworkVlan, "destroyNetworkVlan": destroyNetworkVlan
             }
    args = getArgs()
    cmd[args['COMMAND']](args)
#    attachInerfaceVM()
#    attachDiskVM()
    exit(1)

if __name__ == '__main__':
    main()

