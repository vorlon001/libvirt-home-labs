package main

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

        "github.com/sirupsen/logrus"
)


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

	log.WithFields(logrus.Fields{ "networkConfigTpl": networkConfigTpl, }).Info("CreateNetworkConfig")

	d, err := a.ReadFile(&networkConfigTpl)
	if err != nil {
		log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
		return
	}

	log.WithFields(logrus.Fields{ "networkConfig": *d, }).Info("CreateNetworkConfig")

        render, err := templateRender(*d, *a)
        if err != nil {
		log.WithFields(logrus.Fields{ "err": err, }).Info("CreateNetworkConfig")
                return
        }


        networkConfig := "network-config"
        networkConfig = fmt.Sprintf("%s%s",a.VMPath, networkConfig)
        _ = a.WriteInFile(&networkConfig, &render)

}

func (a *LibVirtVM) CreateUserData() {

        var userDataTpl string
        userDataTpl = a.Config.USERDATATPLPATH //"user-data.tpl"

	log.WithFields(logrus.Fields{ "userDataTpl": userDataTpl, }).Info("CreateUserData")

        d2, err := a.ReadFile(&userDataTpl)
	if err != nil {
		log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
		return
	}

        render2, err := templateRender(*d2, *a)
        if err != nil {
		log.WithFields(logrus.Fields{ "err": err, }).Info("CreateUserData")
                return
        }

        userData := "user-data"
        userData = fmt.Sprintf("%s%s",a.VMPath, userData)

        _ = a.WriteInFile(&userData, &render2)


}


func (a *LibVirtVM) PreInitScriptVM() {

        for _,v := range a.CreateImageVM {
                render, err := templateRender(v, *a)
		if err != nil {
			log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
			return
		}

		log.WithFields(logrus.Fields{ "render": render, }).Info("PreInitScriptVM")

                out, errout, err := a.Shellout(render)
                if err != nil {
			log.WithFields(logrus.Fields{ "out": out, "errout": errout, "err": err, }).Info("PreInitScriptVM")
			return
                }

		log.WithFields(logrus.Fields{ "out": out, }).Info("PreInitScriptVM")
		log.WithFields(logrus.Fields{ "errout": errout, }).Info("PreInitScriptVM")

        }

	log.Info("PreInitScriptVM Done!")
}


func (a *LibVirtVM) AfterDeployVM() {

        for _,v := range a.AfterDeploy {
                render, err := templateRender(v, *a)
		if err != nil {
			log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
			return
		}


                out, errout, err := a.Shellout(render)
                if err != nil {
                        log.WithFields(logrus.Fields{ "out": out, "errout": errout, "err": err, }).Info("PreInitScriptVM")
                        return
                }

		log.WithFields(logrus.Fields{ "out": out, }).Info("AfterDeployVM")
		log.WithFields(logrus.Fields{ "errout": errout, }).Info("AfterDeployVM")

        }
}

func (a *LibVirtVM) CreateVMXML() {

	log.Info("Init CreateVMXML")

        var vmTemplateXml string
        vmTemplateXml = a.Config.VMTEMPLATE //"vm.template.xml"

        d3, err := a.ReadFile(&vmTemplateXml)
	if err != nil {
		log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
		return
	}


        render3, err := templateRender(*d3, *a)
        if err != nil {
		log.WithFields(logrus.Fields{ "err": err, }).Info("CreateVMXML")
                return
        }


        vmXml := "vm.xml"
        vmXml = fmt.Sprintf("%s%s",a.VMPath, vmXml)

        _ = a.WriteInFile(&vmXml, &render3)

	core := Singleton[Core]()
	core.XmlTemplate = vmXml
	a.XmlTemplate = vmXml

	log.Info("CreateVMXML Done!")
}

func (a *LibVirtVM) getConf(filename string) *LibVirtVM {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.WithFields(logrus.Fields{ "yamlFile.Get-err": err, }).Info("getConf")
	}

	err = yaml.Unmarshal(yamlFile, a)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return a

}

func templateRender(templateMust string, configAbstract interface{}) (string, error) {

        tpl := template.Must(template.New("").Parse(templateMust))
        var tplBuffer bytes.Buffer
        if err := tpl.Execute(&tplBuffer, configAbstract); err != nil {
		log.WithFields(logrus.Fields{ "err": err, }).Info("templateRender")
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
		log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
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



func LoadConfigVM(core *Core, increment int) *LibVirtVM {

        var c LibVirtVM


        log.WithFields(logrus.Fields{ "core.VMNAME": core.VMNAME, "core.CORE": core.CORE, "core.MEMORY": core.MEMORY,}).Info("Inside LoadConfigVM")
        log.WithFields(logrus.Fields{ "core.ROOTFS_SIZE": core.ROOTFS_SIZE, "core.Octet": core.Octet, "core.EXT_DISK_SIZE": core.EXT_DISK_SIZE,}).Info("Inside LoadConfigVM")
        log.WithFields(logrus.Fields{ "core.USER_DATA_PATH": core.USER_DATA_PATH,}).Info("Inside LoadConfigVM")


        c.getConf(core.USER_DATA_PATH)

        c.SetNodeId(VMOctetAddIncrement(core.Octet, increment))
        c.SetVMNAME(VMOctetAddIncrement(core.VMNAME, increment))

	log.WithFields(logrus.Fields{ "LibVirtVM": c, "core": core,}).Info("Inside LoadConfigVM")

        c.SetVMNAME_FQDN( fmt.Sprintf("%s.%s",core.VMNAME, c.Config.VMNAMEFQDN ))
        c.SetROOTFS_SIZE(core.ROOTFS_SIZE)
        c.SetMEMORY(core.MEMORY*1024)
        c.SetCORE(core.CORE)

        for k,_ := range c.Config.INTERFACEINIT {
                _ = c.Config.INTERFACEINIT[k].CreateMacAddress(c.Network.MagicMac)
		log.WithFields(logrus.Fields{ "c.Config.INTERFACEINIT[k].GetMacAddress()": c.Config.INTERFACEINIT[k].GetMacAddress(),
						"c.Config.INTERFACEINIT[k].MacAddress": c.Config.INTERFACEINIT[k].MacAddress,}).Info("Inside LoadConfigVM Run with args")
        }


        for k,_ := range c.VMDisk {
                _ = c.VMDisk[k].CreatePath(c)
		log.WithFields(logrus.Fields{ "c.VMDisk[k]": c.VMDisk[k],}).Info("Inside LoadConfigVM Run with args")
        }

        VMPathTmpl := "{{.Config.VMPATH}}/{{.VMNAME}}/"
        VMPath, err := templateRender(VMPathTmpl, c)
        if err != nil {
		log.WithFields(logrus.Fields{ "err": err,}).Info("Inside LoadConfigVM Run with args")
                return nil
        }

	log.WithFields(logrus.Fields{ "VMPath": VMPath,}).Info("Inside LoadConfigVM Run with args")

        if err := os.MkdirAll(VMPath, os.ModePerm); err != nil {
                log.Fatal(err)
        }

        c.VMPath = VMPath

        return &c
}
