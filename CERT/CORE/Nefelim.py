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
                                "CORE.Logger",
                                "CORE.Warp",
                                "CORE.Connector"
				)
    CORE.LOADER.iglob(globals(),[
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "subprocess",  "method":"STDOUT"                      },
                            { "module": "subprocess",  "method":"PIPE"                        },
                            { "module": "jinja2",      "method":"Template"                    },
                            { "module": "pprint",      "method":"pprint",   "as": "dump"      },
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "CORE.Connector",   "method":"Connector","as": "Connector"   },
                            { "module": "CORE.Logger",      "method":"log"                    },
                            { "module": "CORE.Warp",   "method":"Decorator"                   }
                          ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)




@Decorator("decorated class Nefelim")
class Nefelim(Connector):

    def __init_subclass__(self,VMNAME: str = "node100", CORE: int = 2, MEMORY:int = 8, octet: int = 200, ROOTFS_SIZE: int = 20, EXT_DISK_SIZE: int = 20, USER_DATA_PATH: str = "CONFIG/user-data.yaml" ):
        super(Connector,self).__init_subclass__(VMNAME,CORE,MEMORY,octet,ROOTFS_SIZE,EXT_DISK_SIZE,USER_DATA_PATH)

        log(message = "Nefelim().__init__()")


    def Run(self):

        log(message = "Nefelim().Run()")

        self.initConfig()

        self.setVarConfig()
        self.loadConfig()
        self.UserDataConfig()
        self.NetworkConfig()

        log(message = "Nefelim().LibvirtConfig()")
        self.LibvirtConfig()
        log(message = "Nefelim().LibvirtConfig()")
        self.LibvirtRunVm()
        log(message = "Nefelim().LibvirtConfig()")
        self.LibvirtRunStep2Vm()
        log(message = "Nefelim().LibvirtConfig()")
        self.LibvirtRunStep3Vm()
        log(message = "Nefelim().Run() END")


    def UserDataConfig(self):

        _,self.create_image_vm = (tmp_vars:=self.network["block"][[ i for i in copy.copy(self.network['block'])][:1].pop() ]["network"]), \
                            [ i.format(VMNAME=self.VMNAME,octet=self.octet,network=tmp_vars,ROOTFS_SIZE=self.ROOTFS_SIZE,EXT_DISK_SIZE=self.EXT_DISK_SIZE ) for i in self.create_image_vm_tpl ]

        self.IP_ADDR_200="{network}{octet}".format(network=self.network["block"][[ i for i in copy.copy(self.network['block'])][:1].pop() ]["network"],octet=self.octet)

        print("POINT 10054", self.IP_ADDR_200 )
        print("POINT 10055", self.create_image_vm )

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
            "root_cert_append": self.root_cert_append
        }

        self.user_data = self.user_data_template.render( **Args )
        self.fileWrite( "user-data.node100", self.user_data)



    def NetworkConfig(self):

        _ = {    v.update({ "ipaddress": str(ipaddress.IPv4Address( f"{v['network']}{self.octet}")) }) for k,v in self.node_ip.items() }

        self.network_config = self.Network_Config_Template.render( node_ip = self.node_ip, INTERFACE = self.INTERFACE )

        self.fileWrite( "network-config.node100", self.network_config )

    def createXML(self):

        self.doc['domain']['memory']['#text']=f"{self.MEMORY*1024*1024}"
        self.doc['domain']['vcpu']['#text']=f"{self.CORE}"
        self.doc['domain']['name']=self.VMNAME

        for count, path in enumerate(self.vm_disk):
            print("DEBUG 10105", count, path )
            self.doc['domain']['devices']['disk'][count]["source"]["@file"]=path

        for count, mac in enumerate(self.INTERFACE):
            print("DEBUG 10106", count, mac["mac"] )
            self.doc['domain']['devices']['interface'][count]['mac']['@address']=mac["mac"]


    def LibvirtConfig(self):

        _ = [ (print("POINT 12000",i),(returned_value:=subprocess.call(i, shell=True)),print('returned value:', returned_value)) for i in self.create_image_vm ]

        self.createXML()

        print(self.doc)
        dump(self.doc)

        print("create xml config")
        self.xml_cfg = xmltodict.unparse(self.doc, pretty=True)
        print("save xml config")

        self.fileWrite( "node100.xml", self.xml_cfg)


    def LibvirtRunStep2Vm(self):

        # -----------------------------

        print("waiting for the cloud init settings")
        self.getStatusVmgetStatusShutDown()

        print("init start VM")
        #l(conn,"node100").create()
        self.vmCreate(self.conn,self.VMNAME)

        # -----------------------------

        self.timeSleep(sec=20, message="sleep 20sec - LibvirtRunStep2Vm")

        # -----------------------------

        self.dom = self.vmInfo(self.conn,self.VMNAME)
        if self.dom == None:
            print('Failed to get the domain object', file=sys.stderr)
            exit(1)

        self.disk_seed="""
<disk type="file" device="disk">
        <driver name="qemu" type="raw"></driver>
        <source file="/cloud/TEST.1/KVM/{VMNAME}-seed.qcow2"></source>
        <target dev="vda" bus="virtio"></target>
</disk>
"""

        self.disk_seed = self.disk_seed.format(VMNAME = self.VMNAME)
        self.vm = self.vmObject(self.conn,self.VMNAME)
        self.vm.detachDeviceFlags( self.disk_seed, flags = libvirt.VIR_DOMAIN_AFFECT_CURRENT | libvirt.VIR_DOMAIN_AFFECT_CONFIG | libvirt.VIR_DOMAIN_AFFECT_LIVE)

        # -----------------------------


    def LibvirtRunStep3Vm(self):
        # -----------------------------
        print("init shutdown  VM - LibvirtRunStep3Vm")
        #l(conn,"node100").shutdown()
        self.vmShutdown(self.conn,self.VMNAME)
        # -----------------------------
        self.getStatusVmgetStatusShutDown()
        # -----------------------------
        print("init start VM")
        #l(conn,"node100").create()
        self.vmCreate(self.conn,self.VMNAME)
        # -----------------------------
        self.getStatusVmgetStatusRunning()
        # -----------------------------
        self.ipaddress = self.IP_ADDR_200
        self.getStatusSsh(ipaddress=self.ipaddress)
        # -----------------------------
        self.timeSleep(sec=20, message="sleep 20sec")
        # -----------------------------
        print(self.vmInfo(self.conn,self.VMNAME))
        self.cmd = [ i.format(self.ipaddress) for i in self.user_data_json["after-deploy"] ]
        dump( self.cmd)
        for i in self.cmd:
            returned_value = subprocess.call(i, shell=True)
            print('returned value:', returned_value)
            print( self.vmInfo(self.conn,self.VMNAME))
        # -----------------------------
        print("init shutdown  VM")
        self.vmShutdown(self.conn,self.VMNAME)
        # -----------------------------
        self.getStatusVmgetStatusShutDown()
        # -----------------------------
        print("init start VM")
        self.vmCreate(self.conn,self.VMNAME)
        # -----------------------------
        self.getStatusVmgetStatusRunning()
        # -----------------------------
        self.ipaddress = self.IP_ADDR_200
        self.getStatusSsh(ipaddress=self.ipaddress)
        # -----------------------------
        self.timeSleep(sec=20, message="sleep 20sec")
        # -----------------------------
        self.closeLibvirtD()


