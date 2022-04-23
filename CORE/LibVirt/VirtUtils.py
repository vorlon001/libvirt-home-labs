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
                                "CORE.Base.Base",
                                "CORE.Core.Warp"
                                )
    boot.iglob(globals(),[
                            { "module": "subprocess",  "method":"run"                    },
                            { "module": "subprocess",  "method":"STDOUT"                 },
                            { "module": "subprocess",  "method":"PIPE"                   },
                            { "module": "jinja2",      "method":"Template"               },
                            { "module": "pprint",      "method":"pprint",   "as": "dump" },
                            { "module": "subprocess",  "method":"run"                    },
                            { "module": "CORE.Base.Base",   "method":"LoadUtils","as": "LoadUtils"},
                            { "module": "CORE.Core.Warp",   "method":"Decorator"              },
                            { "module": "CORE.Core.Warp",   "method":"WARP_DRIVE"             }
                          ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)


@Decorator("decorated class VirtUtils")
class VirtUtils(LoadUtils):

    @WARP_DRIVE.decorator
    def random_mac(self, network:str) -> str:
        return f"{network['magic_mac']}:{':'.join([f'{random.randint(0, 255):02x}' for _ in range(3)])}"

    @WARP_DRIVE.decorator
    def vmInfo(self, conn: object, VMNAME:str) -> list:
        return [ { "ID": i.ID(), "name": i.name(), "UUID": i.UUIDString(), "state": i.state(), "object": i } for i in conn.listAllDomains(0) if VMNAME==i.name() ]

    @WARP_DRIVE.decorator
    def vmObject(self, conn:object, VMNAME:str) ->object:
        _,vmObject = (tmp:=self.vmInfo(conn,VMNAME)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        return vmObject.pop()["object"]


    @WARP_DRIVE.decorator_void
    def vmCreate(self, conn:object, VMNAME:str):
        _,vmObject = (tmp:=self.vmObject(conn,VMNAME)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        vmObject.create()


    @WARP_DRIVE.decorator_void
    def vmShutdown(self, conn:object, VMNAME:str):
        _,vmObject = (tmp:=self.vmObject(conn,VMNAME)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        vmObject.shutdown()


    def __init_subclass__(self):
        super(LoadUtils,self).__init_subclass__()



    @WARP_DRIVE.decorator_void
    def getStatusVm(self, conn: object, VMNAME: str, status: int, message: str):
        while True:
            _,vmInfo = (tmp:=self.vmInfo(conn,VMNAME)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
            if vmInfo[0]["state"][0]==status:
                print( message )
                return
            self.timeSleep(sec=5,message="")
