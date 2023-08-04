package virsh

import (

        "bytes"
        "os/exec"
        "html/template"
        "gopkg.in/yaml.v2"

	"fmt"

        "encoding/json"
        "io/ioutil"
        "os"
        "strings"
        "gitlab.iblog.pro/cobra/store"
        "github.com/sirupsen/logrus"
        logs "gitlab.iblog.pro/cobra/libvirtgrpc/internal/cobra/logs"
	coreUtils "gitlab.iblog.pro/cobra/libvirtgrpc/internal/cobra/core/utils"
        Model "gitlab.iblog.pro/cobra/libvirtgrpc/internal/cobra/core/model"
	InterfaceDisk "gitlab.iblog.pro/cobra/libvirtgrpc/internal/cobra/core/interfacedisk"
)

type LibVirtVM struct {

        VMPath          string
        VMNAME          string  `yaml:"VMNAME"`
        VMNAME_FQDN     string  `yaml:"VMNAME_FQDN"`
        NodeId          string  `yaml:"nodeid"`
        ROOTFS_SIZE     int     `yaml:"ROOTFS_SIZE"`
        DISKID          string  `yaml:"DISKID"`
        EXT_DISK_SIZE   int     `yaml:"EXT_DISK_SIZE"`
        DetachDiskName  string
        XmlTemplate     string

        MEMORY                  int
        CORE                    int
        DISKSDA                 string
        DISKSDBCLOUDINIT        string

        AfterDeploy   	[]string `yaml:"after-deploy"`
        Command       	[]string `yaml:"command"`
        Config        	Model.Config   `yaml:"config"`
        CreateImageVM 	[]string `yaml:"create-image-vm"`
        Network       	Model.Network  `yaml:"network"`
        Pgk		[]string `yaml:"pgk"`
        SSHKeys       	[]string `yaml:"ssh-keys"`
        VMDisk        	[]InterfaceDisk.VMDisk `yaml:"vm-disk"`
}


func (a LibVirtVM) ReadFromFile(filename string) []string {
        contentBytes, err := ioutil.ReadFile(filename)
        if err != nil {
            return []string{}
        }
        contentString := string(contentBytes)
        return strings.Split(contentString, "\n")
    }


func (a *LibVirtVM) SetROOTFS_SIZE(ROOTFS_SIZE int){
	a.ROOTFS_SIZE = ROOTFS_SIZE
}

func (a *LibVirtVM) SetNodeId(NodeId string){
	a.NodeId = NodeId
}

func (a *LibVirtVM) SetVMNAME(VMNAME string){
	a.VMNAME = VMNAME
}

func (a *LibVirtVM) SetVMNAME_FQDN(VMNAME_FQDN string){
	a.VMNAME_FQDN = VMNAME_FQDN
}


func (a *LibVirtVM) SetMEMORY(MEMORY int){
	a.MEMORY = MEMORY
}

func (a *LibVirtVM) SetCORE(CORE int){
	a.CORE = CORE
}


func (a *LibVirtVM) CreateNetworkConfig() {

	var networkConfigTpl string
	networkConfigTpl = a.Config.NETWORKCONFIGTPL

	logs.Log.WithFields(logrus.Fields{ "networkConfigTpl": networkConfigTpl, }).Info("CreateNetworkConfig")

	d, err := a.ReadFile(&networkConfigTpl)
	if err != nil {
		logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
		return
	}

	logs.Log.WithFields(logrus.Fields{ "networkConfig": *d, }).Info("CreateNetworkConfig")

        render, err := a.TemplateRender(*d, *a)
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("CreateNetworkConfig")
                return
        }


        networkConfig := "network-config"
        networkConfig = fmt.Sprintf("%s%s",a.VMPath, networkConfig)
        _ = a.WriteInFile(&networkConfig, &render)

}

func (a *LibVirtVM) CreateUserData() {

        var userDataTpl string
        userDataTpl = a.Config.USERDATATPLPATH //"user-data.tpl"

	logs.Log.WithFields(logrus.Fields{ "userDataTpl": userDataTpl, }).Info("CreateUserData")

        d2, err := a.ReadFile(&userDataTpl)
	if err != nil {
		logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
		return
	}

        render2, err := a.TemplateRender(*d2, *a)
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("CreateUserData")
                return
        }

        userData := "user-data"
        userData = fmt.Sprintf("%s%s",a.VMPath, userData)

        _ = a.WriteInFile(&userData, &render2)


}


func (a *LibVirtVM) PreInitScriptVM() {

        for _,v := range a.CreateImageVM {
                render, err := a.TemplateRender(v, *a)
		if err != nil {
			logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
			return
		}

		logs.Log.WithFields(logrus.Fields{ "render": render, }).Info("PreInitScriptVM")

                out, errout, err := a.Shellout(render)
                if err != nil {
			logs.Log.WithFields(logrus.Fields{ "out": out, "errout": errout, "err": err, }).Info("PreInitScriptVM")
			return
                }

		logs.Log.WithFields(logrus.Fields{ "out": out, }).Info("PreInitScriptVM")
		logs.Log.WithFields(logrus.Fields{ "errout": errout, }).Info("PreInitScriptVM")

        }

	logs.Log.Info("PreInitScriptVM Done!")
}


func (a *LibVirtVM) AfterDeployVM() {

        for _,v := range a.AfterDeploy {
                render, err := a.TemplateRender(v, *a)
		if err != nil {
			logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
			return
		}


                out, errout, err := a.Shellout(render)
                if err != nil {
                        logs.Log.WithFields(logrus.Fields{ "out": out, "errout": errout, "err": err, }).Info("PreInitScriptVM")
                        return
                }

		logs.Log.WithFields(logrus.Fields{ "out": out, }).Info("AfterDeployVM")
		logs.Log.WithFields(logrus.Fields{ "errout": errout, }).Info("AfterDeployVM")

        }
}

func (a *LibVirtVM) CreateVMXML() {

	logs.Log.Info("Init CreateVMXML")

        var vmTemplateXml string
        vmTemplateXml = a.Config.VMTEMPLATE //"vm.template.xml"

        d3, err := a.ReadFile(&vmTemplateXml)
	if err != nil {
		logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
		return
	}


        render3, err := a.TemplateRender(*d3, *a)
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("CreateVMXML")
                return
        }


        vmXml := "vm.xml"
        vmXml = fmt.Sprintf("%s%s",a.VMPath, vmXml)

        _ = a.WriteInFile(&vmXml, &render3)

	core := store.Singleton[Model.Core]()
	core.XmlTemplate = vmXml
	a.XmlTemplate = vmXml

	logs.Log.Info("CreateVMXML Done!")
}

func (a *LibVirtVM) getConf(filename string) *LibVirtVM {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logs.Log.WithFields(logrus.Fields{ "yamlFile.Get-err": err, }).Info("getConf")
	}

	err = yaml.Unmarshal(yamlFile, a)
	if err != nil {
		logs.Log.Fatalf("Unmarshal: %v", err)
	}

	return a

}

func (a *LibVirtVM) TemplateRender(templateMust string, configAbstract interface{}) (string, error) {

        tpl := template.Must(template.New("").Parse(templateMust))
        var tplBuffer bytes.Buffer
        if err := tpl.Execute(&tplBuffer, configAbstract); err != nil {
		logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("TemplateRender")
                return  "",err
        }
        render := tplBuffer.String()
        replaceContent := [][]string{ []string{"&#34;","\""}, []string{"&#43;","+"}}
        for _,v := range replaceContent {
                render = strings.ReplaceAll( render, v[0], v[1])
        }
        return render,nil

}


func (a *LibVirtVM) Var_dump(expression ...interface{} ) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

func (a *LibVirtVM) prettyPrint(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
		return ""
	}

	return string(s)
}



func (a *LibVirtVM) ReadFile(filename *string) (*string, error) {
	contentBytes, err := ioutil.ReadFile(*filename)
	if err != nil {
		return nil, err
	}
	contentString := string(contentBytes)
	return &contentString, nil
}


func (a *LibVirtVM) WriteInFile(filename *string, data *string) error {
        return ioutil.WriteFile(*filename, []byte(*data), 0644)
}





const ShellToUse = "bash"

func (a *LibVirtVM) Shellout(command string) (string, string, error) {
        var stdout bytes.Buffer
        var stderr bytes.Buffer
        cmd := exec.Command(ShellToUse, "-c", command)
        cmd.Stdout = &stdout
        cmd.Stderr = &stderr
        err := cmd.Run()
        return stdout.String(), stderr.String(), err
}



func LoadConfigVM(core *Model.Core, increment int) *LibVirtVM {

        var c LibVirtVM


        logs.Log.WithFields(logrus.Fields{ "core.VMNAME": core.VMNAME, "core.CORE": core.CORE, "core.MEMORY": core.MEMORY,}).Info("Inside LoadConfigVM")
        logs.Log.WithFields(logrus.Fields{ "core.ROOTFS_SIZE": core.ROOTFS_SIZE, "core.Octet": core.Octet, "core.EXT_DISK_SIZE": core.EXT_DISK_SIZE,}).Info("Inside LoadConfigVM")
        logs.Log.WithFields(logrus.Fields{ "core.USER_DATA_PATH": core.USER_DATA_PATH,}).Info("Inside LoadConfigVM")


        c.getConf(core.USER_DATA_PATH)

        c.SetNodeId(coreUtils.VMOctetAddIncrement(core.Octet, increment))
        c.SetVMNAME(coreUtils.VMOctetAddIncrement(core.VMNAME, increment))

	logs.Log.WithFields(logrus.Fields{ "LibVirtVM": c, "core": core,}).Info("Inside LoadConfigVM")

        c.SetVMNAME_FQDN( fmt.Sprintf("%s.%s",core.VMNAME, c.Config.VMNAMEFQDN ))
        c.SetROOTFS_SIZE(core.ROOTFS_SIZE)
        c.SetMEMORY(core.MEMORY*1024)
        c.SetCORE(core.CORE)

        for k,_ := range c.Config.INTERFACEINIT {
                _ = c.Config.INTERFACEINIT[k].CreateMacAddress(c.Network.MagicMac)
		logs.Log.WithFields(logrus.Fields{ "c.Config.INTERFACEINIT[k].GetMacAddress()": c.Config.INTERFACEINIT[k].GetMacAddress(),
						"c.Config.INTERFACEINIT[k].MacAddress": c.Config.INTERFACEINIT[k].MacAddress,}).Info("Inside LoadConfigVM Run with args")
        }


        for k,_ := range c.VMDisk {
                _ = c.VMDisk[k].CreatePath(c)
		logs.Log.WithFields(logrus.Fields{ "c.VMDisk[k]": c.VMDisk[k],}).Info("Inside LoadConfigVM Run with args")
        }

        VMPathTmpl := "{{.Config.VMPATH}}/{{.VMNAME}}/"
        VMPath, err := c.TemplateRender(VMPathTmpl, c)
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "err": err,}).Info("Inside LoadConfigVM Run with args")
                return nil
        }

	logs.Log.WithFields(logrus.Fields{ "VMPath": VMPath,}).Info("Inside LoadConfigVM Run with args")

        if err := os.MkdirAll(VMPath, os.ModePerm); err != nil {
                logs.Log.Fatal(err)
        }

        c.VMPath = VMPath

        return &c
}
