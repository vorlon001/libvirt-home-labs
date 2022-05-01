try:
    import CORE.Core.LOADER as boot
    boot.lm(globals(),
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
                                "CORE.Core.Logger",
				"CORE.LibVirt.Genesys",
                                "CORE.Core.Warp"
				)
    boot.iglob(globals(),[
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "subprocess",  "method":"STDOUT"                      },
                            { "module": "subprocess",  "method":"PIPE"                        },
                            { "module": "jinja2",      "method":"Template"                    },
                            { "module": "pprint",      "method":"pprint",   "as": "dump"      },
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "CORE.LibVirt.Genesys","method":"Genesys","as": "Genesys"     },
                            { "module": "CORE.Core.Logger",      "method":"log"                    },
                            { "module": "CORE.Core.Warp",   "method":"Decorator"                   },
                            { "module": "CORE.Core.Warp",   "method":"WARP_DRIVE"                  }
                          ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)



#@Decorator("decorated class Connector")
class Connector(Genesys):

    def __init_subclass__(self): #,VMNAME: str = "node100", CORE: int = 2, MEMORY:int = 8, octet: int = 200, ROOTFS_SIZE: int = 20, EXT_DISK_SIZE: int = 20, USER_DATA_PATH: str = "CONFIG/user-data.yaml" ):
        super(Genesys,self).__init_subclass__() #VMNAME,CORE,MEMORY,octet,ROOTFS_SIZE,EXT_DISK_SIZE,USER_DATA_PATH)


    @WARP_DRIVE.decorator_void
    def connectLibvirtD(self):
        self.conn = None
        try:
            self.conn = libvirt.open("qemu:///system")
        except libvirt.libvirtError as e:
            log(message = f"{repr(e)}")
            log(message = f"{sys.stderr}")
            exit(1)


    @WARP_DRIVE.decorator_void
    def closeLibvirtD(self):
        self.conn.close()


    @WARP_DRIVE.decorator_void
    def loadXMLconfigLibvirtD(self):
        dom = None
        try:
            self.dom = self.conn.defineXMLFlags(self.xml_cfg, 0)
        except libvirt.libvirtError as e:
            log(message = f"{repr(e)}")
            log(message = f"{sys.stderr}")
            exit(1)


    @WARP_DRIVE.decorator_void
    def startVmLibvirtD(self):
        if self.dom.create() < 0:
            log(message = f"Can not boot guest domain.: {sys.stderr}")
            exit(1)
        log(message = f"Guest: {self.dom.name()} has booted, file={sys.stderr}")
        log(message = f"{self.vmInfo(self.conn,self.VMNAME)}")


    @WARP_DRIVE.decorator_void
    def initRunNewVm(self):

        self.connectLibvirtD()
        self.loadXMLconfigLibvirtD()

        log(message = f"init first start vm")

        self.startVmLibvirtD()


    @WARP_DRIVE.decorator_void
    def getStatusVmgetStatusShutDown(self):
        self.getStatusVm( self.conn, self.VMNAME, libvirt.VIR_DOMAIN_SHUTOFF, "VM is ShutDown" )


    @WARP_DRIVE.decorator_void
    def getStatusVmgetStatusRunning(self):
        self.getStatusVm( self.conn, self.VMNAME, libvirt.VIR_DOMAIN_RUNNING, "VM is init is DONE!" )




    @WARP_DRIVE.decorator
    def random_mac(self, network:str) -> str:
        log(message = "run VirtUtils.random_mac()")
        return f"{network['magic_mac']}:{':'.join([f'{random.randint(0, 255):02x}' for _ in range(3)])}"

    @WARP_DRIVE.decorator
    def vmInfo(self, conn: object, VMNAME:str) -> list:
        log(message = "run VirtUtils.vmInfo()")
        return [ { "ID": i.ID(), "name": i.name(), "UUID": i.UUIDString(), "state": i.state(), "object": i } for i in conn.listAllDomains(0) if VMNAME==i.name() ]

    @WARP_DRIVE.decorator
    def vmObject(self, conn:object, VMNAME:str) ->object:
        log(message = "run VirtUtils.vmObject()")
        _,vmObject = (tmp:=self.vmInfo(conn,VMNAME)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        return vmObject.pop()["object"]


    @WARP_DRIVE.decorator_void
    def vmCreate(self, conn:object, VMNAME:str):
        log(message = "run VirtUtils.vmCreate()")
        _,vmObject = (tmp:=self.vmObject(conn,VMNAME)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        vmObject.create()


    @WARP_DRIVE.decorator_void
    def vmShutdown(self, conn:object, VMNAME:str):
        log(message = "run VirtUtils.vmShutdown()")
        _,vmObject = (tmp:=self.vmObject(conn,VMNAME)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        vmObject.shutdown()



    @WARP_DRIVE.decorator_void
    def getStatusVm(self, conn: object, VMNAME: str, status: int, message: str):
        log(message = "run VirtUtils.getStatusVm()")
        while True:
            _,vmInfo = (tmp:=self.vmInfo(conn,VMNAME)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
            if vmInfo[0]["state"][0]==status:
                log(message = message )
                return
            self.timeSleep(sec=20,message="")

