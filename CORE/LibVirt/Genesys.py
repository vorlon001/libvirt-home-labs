try:
    import CORE.Core.LOADER as boot
    boot.lm(globals(),
				"subprocess",
				"xmltodict",
				"libvirt",
				"sys",
				"time",
				"yaml",
				"ipaddress",
				"copy",
				"random",
				"socket",
				"jinja2",
				"pprint",
                                "CORE.Core.Logger",
                                "CORE.Base.Base",
				"CORE.LibVirt.VirtUtils",
                                "CORE.Core.Warp"
				)
    boot.iglob(globals(),[
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "subprocess",  "method":"STDOUT"                      },
                            { "module": "subprocess",  "method":"PIPE"                        },
                            { "module": "jinja2",      "method":"Template"                    },
                            { "module": "pprint",      "method":"pprint",   "as": "dump"      },
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "CORE.Core.Logger",      "method":"log"                    },
                            { "module": "CORE.LibVirt.VirtUtils",   "method":"VirtUtils","as": "VirtUtils" },
                            { "module": "CORE.Core.Logger",      "method":"log"                    },
                            { "module": "CORE.Core.Warp",   "method":"Decorator"                   },
                            { "module": "CORE.Base.Base",   "method":"LoadUtils","as": "LoadUtils" },
                            { "module": "CORE.Core.Warp",   "method":"WARP_DRIVE"             }
                          ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)

    def __init_subclass__(self):
        super(LoadUtils,self).__init_subclass__()


#@Decorator("decorated class Genesys")
class Genesys(VirtUtils):
    def __init_subclass__(self): #,VMNAME: str = "node100", CORE: int = 2, MEMORY:int = 8, octet: int = 200, ROOTFS_SIZE: int = 20, EXT_DISK_SIZE: int = 20, USER_DATA_PATH: str = "CONFIG/user-data.yaml" ):
        super(VirtUtils,self).__init_subclass__()

#        self.CORE = CORE
#        self.MEMORY = MEMORY
#        self.octet = octet
#        self.ROOTFS_SIZE = ROOTFS_SIZE
#        self.EXT_DISK_SIZE = EXT_DISK_SIZE
#        self.VMNAME = VMNAME
#        self.USER_DATA_PATH = USER_DATA_PATH


    @WARP_DRIVE.decorator_void
    def initConfig(self,VMNAME: str = "node100", CORE: int = 2, MEMORY:int = 8, octet: int = 200, ROOTFS_SIZE: int = 20, EXT_DISK_SIZE: int = 20, USER_DATA_PATH: str = "CONFIG/user-data.yaml" ):

        self.CORE = CORE
        self.MEMORY = MEMORY
        self.octet = octet
        self.ROOTFS_SIZE = ROOTFS_SIZE
        self.EXT_DISK_SIZE = EXT_DISK_SIZE
        self.VMNAME = VMNAME
        self.USER_DATA_PATH = USER_DATA_PATH

        log(message = "Genesys().initConfig()")
        _,USER_DATA_PATH = (tmp := self.fileRead(self.USER_DATA_PATH)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        self.user_data_json = yaml.safe_load(USER_DATA_PATH)


    @WARP_DRIVE.decorator_void
    def setVarConfig(self):
        log(message = "Genesys().setVarConfig()")

        self.PREFIX_NETWORK = self.user_data_json["config"]["PREFIX"] if "PREFIX" in self.user_data_json["config"] else None
        self.VMNAME_FQDN = self.user_data_json["config"]["VMNAME_FQDN"].format(VMNAME = self.VMNAME)
        self.VM_REPO = self.user_data_json["config"]["VM_REPO"]
        self.NEXUS_REPO = self.user_data_json["config"]["NEXUS_REPO"]
        self.NEXUS_REPO_SEC = self.user_data_json["config"]["NEXUS_REPO_SEC"]
        self.ROOT_CERT_PATH = self.user_data_json["config"]["ROOT_CERT_PATH"]
        self.SSD_TPL_PATH = self.user_data_json["config"]["SSD_TPL_PATH"]
        self.PIP_TPL_PATH = self.user_data_json["config"]["PIP_TPL_PATH"]
        self.USER_DATA_TPL_PATH = self.user_data_json["config"]["USER_DATA_TPL_PATH"]
        self.SSH_AUTH_KEYS = self.user_data_json["config"]["SSH_AUTH_KEYS"]
        self.NETWORK_CONFIG_TPL = self.user_data_json["config"]["NETWORK_CONFIG_TPL"]
        self.VM_TEMPLATE = self.user_data_json["config"]["VM_TEMPLATE"]
        self.VMPATH = self.user_data_json["config"]["VMPATH"]
        self.VMIMAGEPATH = self.user_data_json["config"]["VMIMAGEPATH"]
        self.CMD = self.user_data_json["command"]
        self.PKG = self.user_data_json["pgk"]
        self.node_ssh_key = self.user_data_json["ssh-keys"]
        self.network = self.user_data_json["network"]
        self.INTERFACE_INIT = self.user_data_json["config"]["INTERFACE_INIT"]

        self.STEP2 = self.user_data_json["config"].get("STEP2",True)
        self.STEP3 = self.user_data_json["config"].get("STEP3",True)


    @WARP_DRIVE.decorator_void
    def loadRootCertConfig(self):
        log(message = "Genesys().loadRootCertConfig()")
        # 1
        self.root_cert_append = self.load_Root_Cert(self.ROOT_CERT_PATH)
#(tmp := self.fileRead(self.ROOT_CERT_PATH)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        _,self.root_ssh_authorized_keys = (tmp := self.fileRead(self.SSH_AUTH_KEYS)), tmp["result"] if tmp["code"]==200 else sys.exit(1);


    @WARP_DRIVE.decorator_void
    def loadSshConfig(self):
        log(message = "Genesys().loadSshConfig()")
        # 2
        _,self.sshd_config = (tmp := self.fileRead(self.SSD_TPL_PATH)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        self.sshd_config_append = self.sshd_config.split("\n")


    @WARP_DRIVE.decorator_void
    def loadPipConfig(self):
        log(message = "Genesys().loadPipConfig()")
        # 3
        _,self.pip_conf = (tmp := self.fileRead(self.PIP_TPL_PATH)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        self.pip_conf_append = self.pip_conf.split("\n")


    @WARP_DRIVE.decorator_void
    def loaduser_data_template_tpl_Config(self):
        log(message = "Genesys().loaduser_data_template_tpl_Config()")
        # 4
        _,self.user_data_template_tpl = (tmp := self.fileRead(self.USER_DATA_TPL_PATH)), tmp["result"] if tmp["code"]==200 else sys.exit(1);


    @WARP_DRIVE.decorator_void
    def load_Network_Config_Tpl_Config(self):
        log(message = "Genesys().load_Network_Config_Tpl_Config()")
        # 5
        _,self.Network_Config_Tpl = (tmp := self.fileRead(self.NETWORK_CONFIG_TPL)), tmp["result"] if tmp["code"]==200 else sys.exit(1);


    @WARP_DRIVE.decorator_void
    def load_Init_Template_Config(self):
        log(message = "Genesys().load_Init_Template_Config()")
        # 6
        self.Network_Config_Template = Template(self.Network_Config_Tpl)
        self.user_data_template = Template(self.user_data_template_tpl)


    @WARP_DRIVE.decorator_void
    def load_Xml_Template_Config(self):
        log(message = "Genesys().load_Xml_Template_Config()")
        # 7
        _,VM_TEMPLATE = (tmp := self.fileRead(self.VM_TEMPLATE)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
        self.doc = xmltodict.parse(VM_TEMPLATE)


    @WARP_DRIVE.decorator_void
    def postLoadConfig(self):
        log(message = "Genesys().postLoadConfig()")
        # 8
        self.create_image_vm_tpl = self.user_data_json['create-image-vm']
        # 9
        self.vm_disk_tpl = self.user_data_json["vm-disk"]
        self.vm_disk = [ i.format(VMPATH=self.VMPATH,VMNAME=self.VMNAME)for i in self.vm_disk_tpl ]
        # 10
        if self.PREFIX_NETWORK!=None:
            self.network_interface = [f"{self.PREFIX}0",f"{self.PREFIX}1"]
        else:
            self.network_interface = self.INTERFACE_INIT #["enp1s0","enp2s0"]
            print("="*40)
            print("="*40)
            print(self.network_interface)
            print("-"*40)
            print(self.INTERFACE_INIT)
            print("-"*40)
            print(self.user_data_json["config"]["INTERFACE_INIT"])
            print("="*40)
            print("="*40)

        self.INTERFACE = [ {'name': i, 'mac': self.random_mac(self.network)["result"]   } for i in self.network_interface ]

        print("="*40)
        print("="*40)
        print(self.INTERFACE)
        print("-"*40)
        print(self.INTERFACE_INIT)
        print("="*40)
        print("="*40)

        # 11
        self.node_ip = copy.deepcopy(self.network["block"])
        # 12
        _,_=(r:={}),[ r.update({ i[2]: {"key":i[1], "type":i[2], "full": i[3]} }) for i in [  i.split(" ")+[i] for i in self.root_ssh_authorized_keys.split("\n") if len(i)>0]]
        self.ssk_key = [ r[i]['full'] for i in self.node_ssh_key if i in r]


    @WARP_DRIVE.decorator_void
    def loadConfig(self):
        log(message = "Genesys().loadConfig()")
        self.loadRootCertConfig()
        self.loadSshConfig()
        self.loadPipConfig()
        self.loaduser_data_template_tpl_Config()
        self.load_Network_Config_Tpl_Config()
        self.load_Init_Template_Config()
        self.load_Xml_Template_Config()
        self.postLoadConfig()


