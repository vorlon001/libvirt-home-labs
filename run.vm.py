try:
    import CORE.Core.LOADER as boot
    boot.lm(globals(),
				"CORE.LibVirt.Nefelim", "sys", "libvirt", "subprocess"
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


def initVM():


    vm = { "VMNAME":"node100", "CORE":4, "MEMORY":8, "octet":100, "ROOTFS_SIZE":20, "EXT_DISK_SIZE":20, "USER_DATA_PATH": "CONFIG/user-data.yaml" }

    nefelim = Nefelim()
    nefelim.initVM( **vm )


def attachInerfaceVM():

    vm = { "VMNAME":"node100", "CORE":4, "MEMORY":8, "octet":100, "ROOTFS_SIZE":20, "EXT_DISK_SIZE":20, "USER_DATA_PATH": "CONFIG/user-data.yaml" }

    nefelim = Nefelim()

    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()

    dom = nefelim.conn.lookupByName("node100")
    docj = nefelim.vmGetJsonConfig(nefelim.conn,"node100")
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


def attachDiskVM():

    vm = { "VMNAME":"node100", "CORE":4, "MEMORY":8, "octet":100, "ROOTFS_SIZE":20, "EXT_DISK_SIZE":20, "USER_DATA_PATH": "CONFIG/user-data.yaml" }

    nefelim = Nefelim()

    nefelim.initConfig( **vm )
    nefelim.setVarConfig()

    nefelim.connectLibvirtD()

    VMNAME="node100"

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



def main():

#    globals()["logger"] = initLog(loggingLevel = logging.DEBUG, exp = __name__)
#    print(sys.modules['__main__'].__dict__["logger"])

    initVM()
#    attachInerfaceVM()
#    attachDiskVM()
    exit(1)

if __name__ == '__main__':
    main()

