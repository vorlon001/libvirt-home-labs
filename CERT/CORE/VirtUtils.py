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
                                "CORE.Base",
                                "CORE.Warp"
                                )
    CORE.LOADER.iglob(globals(),[
                            { "module": "subprocess",  "method":"run"                    },
                            { "module": "subprocess",  "method":"STDOUT"                 },
                            { "module": "subprocess",  "method":"PIPE"                   },
                            { "module": "jinja2",      "method":"Template"               },
                            { "module": "pprint",      "method":"pprint",   "as": "dump" },
                            { "module": "subprocess",  "method":"run"                    },
                            { "module": "CORE.Base",   "method":"LoadUtils","as": "LoadUtils"},
                            { "module": "CORE.Warp",   "method":"Decorator"              }
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

    random_mac = lambda self,network: f"{network['magic_mac']}:{':'.join([f'{random.randint(0, 255):02x}' for _ in range(3)])}"

    vmInfo = lambda self,conn,VMNAME: [ { "ID": i.ID(), "name": i.name(), "UUID": i.UUIDString(), "state": i.state(), "object": i } for i in conn.listAllDomains(0) if VMNAME==i.name() ]
    vmObject = lambda self,conn,VMNAME: self.vmInfo(conn,VMNAME).pop()["object"]
    vmCreate = lambda self,conn,VMNAME: self.vmObject(conn,VMNAME).create()
    vmShutdown = lambda self,conn,VMNAME: self.vmObject(conn,VMNAME).shutdown()

    def __init_subclass__(self):
        super(LoadUtils,self).__init_subclass__()

#    def __init__(self):
#        super(VirtUtils).__init__()

    def getStatusVm(self, conn, VMNAME: str, status: int, message: str):
        while True:
            if self.vmInfo(conn,VMNAME)[0]["state"][0]==status:
                print( message )
                break
            self.timeSleep(sec=5,message="")
