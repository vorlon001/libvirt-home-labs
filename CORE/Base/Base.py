try:

    import CORE.Core.LOADER as boot
    boot.lm(globals(),"subprocess","xmltodict","libvirt","sys","time","json","ipaddress","copy","random","socket","jinja2", "pprint","CORE.Core.Warp")
    boot.iglob(globals(),[
                            { "module": "subprocess",  "method":"run"                    },
                            { "module": "subprocess",  "method":"STDOUT"                 },
                            { "module": "subprocess",  "method":"PIPE"                   },
                            { "module": "jinja2",      "method":"Template"               },
                            { "module": "pprint",      "method":"pprint",   "as": "dump" },
                            { "module": "CORE.Core.Warp",   "method":"Decorator"         },
                            { "module": "CORE.Core.Warp",   "method":"WARP_DRIVE"        }
                          ]);

except Exception as e:
    print("type error: " + str(e),__file__)
    print(e)
    import traceback
    print("INTERNAL ERROR CORE:",traceback.format_exc());
    import sys
    sys.exit(1)


@Decorator("decorated class BASE")
class Base():
    pass

@Decorator("decorated class Utils")
class Utils(Base):
    def __init_subclass__(self):
        super(Base,self).__init_subclass__()

    @WARP_DRIVE.decorator_void
    def timeSleep(self,sec: int = 0, message: str = ""):
        print(f"sleep {sec}sec - {message}")
        time.sleep( sec )


    @WARP_DRIVE.decorator
    def isSshEnable(self, ipaddress: str = "127.0.0.1", port: int = 22):
        s = socket.socket()
        try:
            print("isSshEnable TEST",ipaddress, port)
            s.connect((ipaddress, port))
        except Exception as e:
            return False
        finally:
            s.close()
            return True

    @WARP_DRIVE.decorator_void
    def getStatusSsh(self, ipaddress: str):
        while True:
            _,sshd_config_active = (tmp := self.isSshEnable(ipaddress=ipaddress, port=22)), tmp["result"] if tmp["code"]==200 else sys.exit(1);
            if sshd_config_active==True:
                print("SSH on VM ENABLE", ipaddress)
                break
            self.timeSleep(sec=5)


    @WARP_DRIVE.decorator
    def fileRead(self,filename: str) -> str:
        with open(filename) as file:
            return file.read()


    @WARP_DRIVE.decorator_void
    def fileWrite(self,filename: str, body: str):
        with open( filename, 'w') as file:
            file.write(body)


@Decorator("decorated class LoadUtils")
class LoadUtils(Utils):
    load_Root_Cert_Read = lambda self,filename: self.fileRead(filename)
    load_Root_Cert = lambda self,filename:  "\n".join([ f"     {i}" for i in self.load_Root_Cert_Read(filename)["result"].split("\n") ])
    def __init_subclass__(self):
        super(Utils,self).__init_subclass__()
