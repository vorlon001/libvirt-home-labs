try:
    import CORE.LOADER
    CORE.LOADER.lm(globals(),
				"subprocess",
				"xmltodict",
				"libvirt",
				"sys",
				"time",
				"json",
				"ipaddress",
				"copy",
				"random",
				"socket",
				"jinja2",
				"pprint",
				"CORE.Genesys",
                                "CORE.Warp"
				)
    CORE.LOADER.iglob(globals(),[
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "subprocess",  "method":"STDOUT"                      },
                            { "module": "subprocess",  "method":"PIPE"                        },
                            { "module": "jinja2",      "method":"Template"                    },
                            { "module": "pprint",      "method":"pprint",   "as": "dump"      },
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "CORE.Genesys","method":"Genesys","as": "Genesys"     },
                            { "module": "CORE.Warp",   "method":"Decorator"                   }
                          ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)



@Decorator("decorated class Connector")
class Connector(Genesys):

    def __init_subclass__(self,VMNAME: str = "node100", CORE: int = 2, MEMORY:int = 8, octet: int = 200, ROOTFS_SIZE: int = 20, EXT_DISK_SIZE: int = 20, USER_DATA_PATH: str = "CONFIG/user-data.yaml" ):
        super(Genesys,self).__init_subclass__(VMNAME,CORE,MEMORY,octet,ROOTFS_SIZE,EXT_DISK_SIZE,USER_DATA_PATH)

#    def __init__(self,VMNAME: str = "node100", CORE: int = 2, MEMORY:int = 8, octet: int = 200, ROOTFS_SIZE: int = 20, EXT_DISK_SIZE: int = 20, USER_DATA_PATH: str = "CONFIG/user-data.yaml" ):
#        super(Genesys,self).__init__(VMNAME,CORE,MEMORY,octet,ROOTFS_SIZE,EXT_DISK_SIZE,USER_DATA_PATH)


    def connectLibvirtD(self):
        self.conn = None
        try:
            self.conn = libvirt.open("qemu:///system")
        except libvirt.libvirtError as e:
            print(repr(e), file=sys.stderr)
            exit(1)

    def closeLibvirtD(self):
        self.conn.close()


    def loadXMLconfigLibvirtD(self):
        dom = None
        try:
            self.dom = self.conn.defineXMLFlags(self.xml_cfg, 0)
        except libvirt.libvirtError as e:
            print(repr(e), file=sys.stderr)
            exit(1)

    def startVmLibvirtD(self):
        if self.dom.create() < 0:
            print('Can not boot guest domain.', file=sys.stderr)
            exit(1)
        print('Guest '+self.dom.name()+' has booted', file=sys.stderr)
        print(self.vmInfo(self.conn,self.VMNAME))


    def LibvirtRunVm(self):

        self.connectLibvirtD()
        self.loadXMLconfigLibvirtD()

        print("init first start vm")

        self.startVmLibvirtD()


    def getStatusVmgetStatusShutDown(self):
        self.getStatusVm( self.conn, self.VMNAME, libvirt.VIR_DOMAIN_SHUTOFF, "VM is ShutDown" )


    def getStatusVmgetStatusRunning(self):
        self.getStatusVm( self.conn, self.VMNAME, libvirt.VIR_DOMAIN_RUNNING, "VM is init is DONE!" )




