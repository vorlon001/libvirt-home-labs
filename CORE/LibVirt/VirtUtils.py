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
                            { "module": "CORE.Core.Logger",      "method":"log"                    },
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


#@Decorator("decorated class VirtUtils")
class VirtUtils(LoadUtils):

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
            self.timeSleep(sec=5,message="")


    @WARP_DRIVE.decorator
    def vmLookupByName(self, conn:object, VMNAME:str) -> object:
        return conn.lookupByName(VMNAME)


    @WARP_DRIVE.decorator
    def vmGetXmlConfig(self, conn: object, VMNAME:str) -> object:
        _,vmObject = (tmp:=self.vmLookupByName(conn,VMNAME)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        return xmltodict.parse( vmObject.XMLDesc())


    @WARP_DRIVE.decorator
    def vmGetJsonConfig(self, conn:object, VMNAME:str) -> object:
        _,vmObject = (tmp:=self.vmGetXmlConfig(conn,VMNAME)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        return json.loads(json.dumps( vmObject ))


    @WARP_DRIVE.decorator
    def vmGetDiskJsonConfig(self, docj: dict) -> dict:
        i = docj['domain']['devices']['disk']
        disk_vm = []
        if type(i) is dict:
            disk_vm.append({"source":i["source"]["@file"], "DEVICE":i["target"]["@dev"], "BUS":i["target"]["@bus"]})
        else:
            for j in  i:
                 disk_vm.append({"source":j["source"]["@file"], "DEVICE":j["target"]["@dev"], "BUS":j["target"]["@bus"]})
        return disk_vm


    @WARP_DRIVE.decorator
    def vmGetIdNewDisk(self, docj: dict) -> str:
         _,disk_vm = (tmp:=self.vmGetDiskJsonConfig(docj)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
         max_vm_disk_id=max([ ord(i['DEVICE'][2]) for i in disk_vm ])
         new_disk_vm = f"sd{chr(max_vm_disk_id+1)}"
         return new_disk_vm


    @WARP_DRIVE.decorator
    def vmGetInterface(self, docj: dict) -> dict:
        i = docj['domain']['devices']['interface']
        interface_vm = [] 
        if type(i) is dict:
            j =i
            z = {'address':j['mac']['@address'],'network':j['source']['@network'],'interfaceid':j['virtualport']['parameters']['@interfaceid'],'target':j['target']['@dev'],'address':j['address']['@type'],'domain':j['address']['@domain'],'bus':j['address']['@bus'],'slot':j['address']['@slot'],'function':j['address']['@function'],'multifunction':j['address']['@multifunction']}
            interface_vm.append(z)
        else:
            for k in  i:
                j = k
                z = {'address':j['mac']['@address'],'network':j['source']['@network'],'interfaceid':j['virtualport']['parameters']['@interfaceid'],'target':j['target']['@dev'],'address':j['address']['@type'],'domain':j['address']['@domain'],'bus':j['address']['@bus'],'slot':j['address']['@slot'],'function':j['address']['@function'],'multifunction':j['address']['@multifunction']}
                interface_vm.append(z)
        return interface_vm


    @WARP_DRIVE.decorator
    def vmGetIdNewInterface(self, docj: dict) -> str:
        _,interface_vm = (tmp:=self.vmGetInterface(docj)), tmp["result"] if tmp["code"]==200 else sys.exit(1)
        slot_id_interface_vm = [ int(i['slot'],16) for i in interface_vm if (i['address']=='pci' and i['bus']=='0x03')]
        slot_id_new_interface_vm = hex(max(slot_id_interface_vm) + 1)
        return slot_id_new_interface_vm


    def __init_subclass__(self):
        super(LoadUtils,self).__init_subclass__()
