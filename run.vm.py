try:
    import CORE.Core.LOADER as boot
    boot.lm(globals(),
				"CORE.LibVirt.Nefelim", "sys", "libvirt", "subprocess", "argparse"
				)

    boot.iglob(globals(),[
                            { "module": "CORE.LibVirt.Nefelim",   "method":"Nefelim",     "as": "Nefelim"    }
                          ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)

# python3 init-vm-8.py --COMMAND initVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND initVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND attachDiskVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND attachInerfaceVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20


# python3 init-vm-8.py --COMMAND initVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND attachInerfaceVM --VMNAME node101 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 101  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND attachInerfaceVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20
# python3 init-vm-8.py --COMMAND attachDiskVM --VMNAME node100 --CORE 4  --MEMORY 8 --ROOTFS_SIZE 30 --octet 100  --EXT_DISK_SIZE  20


def initVM(vm: dict):


    del vm['COMMAND']
    nefelim = Nefelim()
    nefelim.initVM( **vm )


def attachInerfaceVM(vm: dict):

    print(vm)

    del vm['COMMAND']
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
    nefelim = Nefelim()
    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()

    VMNAME=vm["VMNAME"]

    dom = nefelim.conn.lookupByName(VMNAME)
    docj = nefelim.vmGetJsonConfig(nefelim.conn,VMNAME)
    docj = docj['result']

    DISKID=nefelim.vmGetIdNewDisk(docj)['result']
    DISKSIZE=5
    cmd=f"qemu-img create -f qcow2 /cloud/TEST.1/KVM/{VMNAME}/{VMNAME}-disk-{DISKID}.qcow2 {DISKSIZE}G"
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


def getArgs() -> dict:
    parser = argparse.ArgumentParser()
    parser.add_argument('--COMMAND', help='COMMAND initVM/attachDiskVM/attachInerfaceVM  help', default='initVM', type=str)
    parser.add_argument('--VMNAME', help='VMNAME help', default='node100', type=str)
    parser.add_argument('--CORE', help='CORE help', default='4', type=int)
    parser.add_argument('--MEMORY', help='MEMORY help', default='8', type=int)
    parser.add_argument('--ROOTFS_SIZE', help='octet help', default='20', type=int)
    parser.add_argument('--octet', help='foo help', default='100', type=int)
    parser.add_argument('--EXT_DISK_SIZE', help='ext-disk-size help', default='20', type=int)
    parser.add_argument('--USER_DATA_PATH', help='USER_DATA_PATH', default="CONFIG/user-data.yaml", type=str)

    args = parser.parse_args()
    print(args)
    print(args.__dict__)
    return vars(args)


def main():


    cmd = { 'initVM': initVM, 'attachDiskVM': attachDiskVM, 'attachInerfaceVM': attachInerfaceVM }
    args = getArgs()
    cmd[args['COMMAND']](args)

    exit(1)

if __name__ == '__main__':
    main()

