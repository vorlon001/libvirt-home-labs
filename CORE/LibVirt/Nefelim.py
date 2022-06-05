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
                                "CORE.Core.Warp",
                                "CORE.LibVirt.Connector"
				)
    boot.iglob(globals(),[
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "subprocess",  "method":"STDOUT"                      },
                            { "module": "subprocess",  "method":"PIPE"                        },
                            { "module": "jinja2",      "method":"Template"                    },
                            { "module": "pprint",      "method":"pprint",   "as": "dump"      },
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "CORE.LibVirt.Connector",   "method":"Connector","as": "Connector"   },
                            { "module": "CORE.Core.Logger",      "method":"log"                    },
                            { "module": "CORE.Core.Warp",   "method":"Decorator"                   },
                            { "module": "CORE.Core.Warp",   "method":"WARP_DRIVE"                  }
                          ]);

except Exception as e:
    print("type error:", str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)




@Decorator("decorated class Nefelim")
class Nefelim(Connector):

    def __init_subclass__(self):
        super(Connector,self).__init_subclass__()
        log(message = "Nefelim().__init__()")


    @WARP_DRIVE.decorator_void
    def initVM(self, VMNAME: str = "node100", CORE: int = 2, MEMORY:int = 8, octet: int = 200, ROOTFS_SIZE: int = 20, EXT_DISK_SIZE: int = 20, USER_DATA_PATH: str = "CONFIG/user-data.yaml" ):


        log(message = "Nefelim().Run()")

        self.initConfig( VMNAME, CORE, MEMORY, octet, ROOTFS_SIZE, EXT_DISK_SIZE, USER_DATA_PATH)

        self.setVarConfig()
        self.loadConfig()

        subprocess.call(f"mkdir -p {self.VMPATH}/{self.VMNAME}", shell=True)


        log(message = "Nefelim().createUserDataConfig()")
        self.createUserDataConfig()
        log(message = "Nefelim().createNetworkConfig()")
        self.createNetworkConfig()

        log(message = "Nefelim().createLibvirtConfig()")
        self.createLibvirtConfig()
        log(message = "Nefelim().LibvirtConfig()")
        self.initRunNewVm()
        log(message = "Nefelim().createRunNewVMStep2()")
        if self.STEP2 == True:
            e = self.createRunNewVMStep2()
        log(message = "Nefelim().createRunNewVMStep3()")
        if self.STEP3 == True:
            self.createRunNewVMStep3()
        log(message = "Nefelim().Run() END")


    @WARP_DRIVE.decorator_void
    def createUserDataConfig(self):

        _,self.create_image_vm = (tmp_vars:=self.network["block"][[ i for i in copy.copy(self.network['block'])][:1].pop() ]["network"]), \
                            [ i.format(VMNAME=self.VMNAME,octet=self.octet,network=tmp_vars,ROOTFS_SIZE=self.ROOTFS_SIZE,EXT_DISK_SIZE=self.EXT_DISK_SIZE,VMPATH=self.VMPATH,VMIMAGEPATH=self.VMIMAGEPATH ) for i in self.create_image_vm_tpl ]

        self.IP_ADDR_200="{network}{octet}".format(network=self.network["block"][[ i for i in copy.copy(self.network['block'])][:1].pop() ]["network"],octet=self.octet)

        log(message = f"POINT 10054 {self.IP_ADDR_200 }")
        log(message = f"POINT 10055 {self.create_image_vm}")

        _ = {    v.update({ "octet": self.octet, "ipaddress": str(ipaddress.IPv4Address( f"{v['network']}{self.octet}")) }) for k,v in self.node_ip.items() }

        Args = {
            "sshd_config_append": self.sshd_config_append,
            "pip_conf_append": self.pip_conf_append,
            "PKG": self.PKG,
            "CMD": self.CMD,
            "NEXUS_REPO": self.NEXUS_REPO,
            "NEXUS_REPO_SEC":  self.NEXUS_REPO_SEC,
            "SSHKEY": self.ssk_key,
            "VMNAME": self.VMNAME,
            "VMNAME_FQDN": self.VMNAME_FQDN,
            "VM_REPO": self.VM_REPO,
            "INTERFACE": self.INTERFACE,
            "node_ip": self.node_ip,
            "root_cert_append": self.root_cert_append
        }

        self.user_data = self.user_data_template.render( **Args )
        self.fileWrite( f"{self.VMPATH}/{self.VMNAME}/user-data", self.user_data)


    @WARP_DRIVE.decorator_void
    def createNetworkConfig(self):

        _ = {    v.update({ "ipaddress": str(ipaddress.IPv4Address( f"{v['network']}{self.octet}")) }) for k,v in self.node_ip.items() }

        self.network_config = self.Network_Config_Template.render( node_ip = self.node_ip, INTERFACE = self.INTERFACE )

        self.fileWrite( f"{self.VMPATH}/{self.VMNAME}/network-config", self.network_config )


    @WARP_DRIVE.decorator_void
    def createXML(self):

        self.doc['domain']['memory']['#text']=f"{self.MEMORY*1024*1024}"
        self.doc['domain']['vcpu']['#text']=f"{self.CORE}"
        self.doc['domain']['name']=self.VMNAME

        for count, path in enumerate(self.vm_disk):
            log(message = f"DEBUG 10105, {count}, {path}")
            self.doc['domain']['devices']['disk'][count]["source"]["@file"]=path

        for count, mac in enumerate(self.INTERFACE):
            log(message = f"DEBUG 10106, {count}, {mac['mac']}")
            self.doc['domain']['devices']['interface'][count]['mac']['@address']=mac["mac"]


    @WARP_DRIVE.decorator_void
    def createLibvirtConfig(self):

        _ = [ (log(message = f"POINT 12000 {i}"),(returned_value:=subprocess.call(i, shell=True)),log(message = f"returned value: {returned_value}")) for i in self.create_image_vm ]

        self.createXML()

        log(message = f"{self.doc}")

        log(message = "create xml config")
        self.xml_cfg = xmltodict.unparse(self.doc, pretty=True)
        log(message = "save xml config")

        self.fileWrite( f"{self.VMPATH}/{self.VMNAME}/{self.VMNAME}.xml", self.xml_cfg)


    @WARP_DRIVE.decorator_void
    def createRunNewVMStep2(self):

        # -----------------------------

        log(message = "waiting for the cloud init settings")
        self.getStatusVmgetStatusShutDown()

        log(message = "init start VM")
        #l(conn,"node100").create()
        self.vmCreate(self.conn,self.VMNAME)

        # -----------------------------

        self.timeSleep(sec=20, message="sleep 20sec - LibvirtRunStep2Vm")

        # -----------------------------

        self.dom = self.vmInfo(self.conn,self.VMNAME)
        if self.dom == None:
            log(message = f"Failed to get the domain object, file={sys.stderr}")
            exit(1)

        self.disk_seed="""
<disk type="file" device="disk">
        <driver name="qemu" type="raw"></driver>
        <source file="/cloud/TEST.1/KVM/{VMNAME}/{VMNAME}-seed.qcow2"></source>
        <target dev="vda" bus="virtio"></target>
</disk>
"""

        self.disk_seed = self.disk_seed.format(VMNAME = self.VMNAME)

        _,self.vm = (tmp:=self.vmObject(self.conn,self.VMNAME)), tmp["result"] if tmp["code"]==200 else sys.exit(1);

        self.vm.detachDeviceFlags( self.disk_seed, flags = libvirt.VIR_DOMAIN_AFFECT_CURRENT | libvirt.VIR_DOMAIN_AFFECT_CONFIG | libvirt.VIR_DOMAIN_AFFECT_LIVE)

        # -----------------------------


    @WARP_DRIVE.decorator_void
    def createRunNewVMStep3(self):
        # -----------------------------
        log(message = "init shutdown  VM - LibvirtRunStep3Vm")
        #l(conn,"node100").shutdown()
        self.vmShutdown(self.conn,self.VMNAME)
        # -----------------------------
        self.getStatusVmgetStatusShutDown()
        # -----------------------------
        log(message = "init start VM")
        #l(conn,"node100").create()
        self.vmCreate(self.conn,self.VMNAME)
        # -----------------------------
        log(message = "Nefelim().self.getStatusVmgetStatusRunning 1")
        self.getStatusVmgetStatusRunning()
        # -----------------------------
        self.ipaddress = self.IP_ADDR_200
        log(message = "Nefelim().self.getStatusVmgetStatusRunning 2")
        self.getStatusSsh(ipaddress=self.ipaddress)
        log(message = "Nefelim().self.getStatusVmgetStatusRunning 3")

        # -----------------------------
        self.timeSleep(sec=20, message="sleep 20sec")
        # -----------------------------

        log(message = f"{self.vmInfo(self.conn,self.VMNAME)}")
        self.cmd = [ i.format(self.ipaddress) for i in self.user_data_json["after-deploy"] ]
        dump( self.cmd)
        for i in self.cmd:
            returned_value = subprocess.call(i, shell=True)
            log(message = f"returned value: {returned_value}")
            log(message = f"{self.vmInfo(self.conn,self.VMNAME)}")
        # -----------------------------
        log(message = "init shutdown  VM")
        self.vmShutdown(self.conn,self.VMNAME)
        # -----------------------------
        self.getStatusVmgetStatusShutDown()
        # -----------------------------
        log(message = "init start VM")
        log(message = "Nefelim().self.getStatusVmgetStatusRunning 5")
        self.vmCreate(self.conn,self.VMNAME)
        # -----------------------------
        log(message = "Nefelim().self.getStatusVmgetStatusRunning 6")
        self.getStatusVmgetStatusRunning()
        # -----------------------------
        self.ipaddress = self.IP_ADDR_200
        log(message = "Nefelim().self.getStatusVmgetStatusRunning 7")
        self.getStatusSsh(ipaddress=self.ipaddress)
        # -----------------------------
        self.timeSleep(sec=20, message="sleep 20sec")
        # -----------------------------
        self.closeLibvirtD()


