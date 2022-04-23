try:
    import CORE.LOADER
    CORE.LOADER.lm(globals(),
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
                                "CORE.Logger",
				"CORE.VirtUtils",
                                "CORE.Warp"
				)
    CORE.LOADER.iglob(globals(),[
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "subprocess",  "method":"STDOUT"                      },
                            { "module": "subprocess",  "method":"PIPE"                        },
                            { "module": "jinja2",      "method":"Template"                    },
                            { "module": "pprint",      "method":"pprint",   "as": "dump"      },
                            { "module": "subprocess",  "method":"run"                         },
                            { "module": "CORE.Logger",      "method":"log"                    },
                            { "module": "CORE.VirtUtils",   "method":"VirtUtils","as": "VirtUtils" },
                            { "module": "CORE.Warp",   "method":"Decorator"                   }
                          ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)



@Decorator("decorated class Genesys")
class Genesys(VirtUtils):
    def __init_subclass__(self,VMNAME: str = "node100", CORE: int = 2, MEMORY:int = 8, octet: int = 200, ROOTFS_SIZE: int = 20, EXT_DISK_SIZE: int = 20, USER_DATA_PATH: str = "CONFIG/user-data.yaml" ):
        super(VirtUtils,self).__init_subclass__() #VMNAME,CORE,MEMORY,octet,ROOTFS_SIZE,EXT_DISK_SIZE,USER_DATA_PATH)

#    def __init__(self,VMNAME: str = "node100", CORE: int = 2, MEMORY:int = 8, octet: int = 200, ROOTFS_SIZE: int = 20, EXT_DISK_SIZE: int = 20, USER_DATA_PATH: str = "CONFIG/user-data.yaml" ):
#        super(Genesys).__init__()

        self.CORE = CORE
        self.MEMORY = MEMORY
        self.octet = octet
        self.ROOTFS_SIZE = ROOTFS_SIZE
        self.EXT_DISK_SIZE = EXT_DISK_SIZE
        self.VMNAME = VMNAME
        self.USER_DATA_PATH = USER_DATA_PATH

    def initConfig(self):
        log(message = "Genesys().initConfig()")
        self.user_data_json = yaml.safe_load(self.fileRead(self.USER_DATA_PATH))

    def setVarConfig(self):

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

        self.CMD = self.user_data_json["command"]
        self.PKG = self.user_data_json["pgk"]
        self.node_ssh_key = self.user_data_json["ssh-keys"]
        self.network = self.user_data_json["network"]

    def loadRootCertConfig(self):
        # 1
        self.root_cert_append = self.load_Root_Cert(self.ROOT_CERT_PATH)
        self.root_ssh_authorized_keys = self.fileRead(self.SSH_AUTH_KEYS)

    def loadSshConfig(self):
        # 2
        self.sshd_config = self.fileRead(self.SSD_TPL_PATH)
        self.sshd_config_append = self.sshd_config.split("\n")

    def loadPipConfig(self):
        # 3
        self.pip_conf = self.fileRead(self.PIP_TPL_PATH)
        self.pip_conf_append = self.pip_conf.split("\n")

    def loaduser_data_template_tpl_Config(self):
        # 4
        self.user_data_template_tpl = self.fileRead(self.USER_DATA_TPL_PATH)

    def load_Network_Config_Tpl_Config(self):
        # 5
        self.Network_Config_Tpl = self.fileRead(self.NETWORK_CONFIG_TPL)

    def load_Init_Template_Config(self):
        # 6
        self.Network_Config_Template = Template(self.Network_Config_Tpl)
        self.user_data_template = Template(self.user_data_template_tpl)

    def load_Xml_Template_Config(self):
        # 7
        self.doc = xmltodict.parse( self.fileRead(self.VM_TEMPLATE))


    def postLoadConfig(self):
        # 8
        self.create_image_vm_tpl = self.user_data_json['create-image-vm']
        # 9
        self.vm_disk_tpl = self.user_data_json["vm-disk"]
        self.vm_disk = [ i.format(VMNAME=self.VMNAME)for i in self.vm_disk_tpl ]
        # 10
        self.network_interface = ["enp1s0","enp2s0"]
        self.INTERFACE = [ {'name': i, 'mac': self.random_mac(self.network)   } for i in self.network_interface ]
        # 11
        self.node_ip = copy.deepcopy(self.network["block"])
        # 12
        _,_=(r:={}),[ r.update({ i[2]: {"key":i[1], "type":i[2], "full": i[3]} }) for i in [  i.split(" ")+[i] for i in self.root_ssh_authorized_keys.split("\n") if len(i)>0]]
        self.ssk_key = [ r[i]['full'] for i in self.node_ssh_key if i in r]


    def loadConfig(self):
        self.loadRootCertConfig()
        self.loadSshConfig()
        self.loadPipConfig()
        self.loaduser_data_template_tpl_Config()
        self.load_Network_Config_Tpl_Config()
        self.load_Init_Template_Config()
        self.load_Xml_Template_Config()
        self.postLoadConfig()


