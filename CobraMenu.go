package main

import (

	"sync"
	"fmt"
	"reflect"
	"time"
        "encoding/xml"
        "strconv"
	"github.com/spf13/cobra"


        "github.com/digitalocean/go-libvirt"

        "github.com/sirupsen/logrus"
)


func (f *CobraMenu) RootCmdLibVirt() *cobra.Command {

  var RootCmdLibVirt = &cobra.Command{
    Use:   "LibVirt [sub]",
    Short: "LibVirt commands",
    Run: func(cmd *cobra.Command, args []string) {
	log.WithFields(logrus.Fields{ "args": args,}).Info("Inside RootCmdLibVirt Run with args")
    },
  }
	return RootCmdLibVirt
}

func (f *CobraMenu) RootCmdConfigure() *cobra.Command {
  var RootCmdConfigure = &cobra.Command{
    Use:   "Configure [sub]",
    Short: "Configure VM command",
    Run: func(cmd *cobra.Command, args []string) {
	log.WithFields(logrus.Fields{ "args": args,}).Info("Inside RootCmdConfigure Run with args")
    },
  }
	return RootCmdConfigure
}



func (f *CobraMenu) RootSubCmdvirtualMachineState() *cobra.Command {
  var subvirtualMachineState = &cobra.Command{
    Use:   "MachineState [--id vmname!]",
    Short: "Returns result with a current machine state",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdvirtualMachineState Run with args")

	Virtinit().RootSubCmdvirtualMachineMachineState()

    },
  }
        return subvirtualMachineState
}



func (f *CobraMenu) RootSubCmdvirtualMachineMigrate() *cobra.Command {
  var subvirtualMachineMigrate = &cobra.Command{
    Use:   "MachineMigrate [--id vmname!]",
    Short: "Migrate up a VM. Returns result with a current machine state",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
        log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdvirtualMachineStart Run with args")
        log.WithFields(logrus.Fields{ "core.toMove": core.ToMove, "args": args,}).Info("Inside RootSubCmdvirtualMachineStart Run with args")
        Virtinit().RootSubCmdvirtualMachineMigrate()

    },
  }
        return subvirtualMachineMigrate
}


func (f *CobraMenu) RootSubCmdvirtualMachineSoftReboot() *cobra.Command {
  var subvirtualMachineSoftReboot = &cobra.Command{
    Use:   "MachineSoftReboot [--id vmname!]",
    Short: "reboots a machine gracefully, as chosen by hypervisor. Returns result with a current machine state",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdvirtualMachineSoftReboot Run with args")
	Virtinit().VirtualMachineSoftReboot(core.VMid)
    },
  }
        return subvirtualMachineSoftReboot
}

func (f *CobraMenu) RootSubCmdvirtualMachineHardReboot() *cobra.Command {
  var subvirtualMachineHardReboot = &cobra.Command{
    Use:   "MachineHardReboot [--id vmname!]",
    Short: "sends a VM into hard-reset mode. This is damaging to all ongoing file operations.",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdvirtualMachineHardReboot Run with args")
	Virtinit().VirtualMachineHardReboot(core.VMid)
    },
  }
        return subvirtualMachineHardReboot
}

func (f *CobraMenu) RootSubCmdvirtualMachineShutdown() *cobra.Command {
  var subvirtualMachineShutdown = &cobra.Command{
    Use:   "MachineShutdown [--id vmname!]",
    Short: "gracefully shuts down the VM. Returns result with a current machine state",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdvirtualMachineShutdown Run with args")
	Virtinit().VirtualMachineShutdown(core.VMid)
    },
  }
        return subvirtualMachineShutdown
}

func (f *CobraMenu) RootSubCmdvirtualMachineShutoff() *cobra.Command {
  var subvirtualMachineShutoff = &cobra.Command{
    Use:   "MachineShutoff [--id vmname!]",
    Short: "kills running VM. Equivalent to pulling a plug out of a computer. Returns result with a current machine state",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdvirtualMachineShutoff Run with args")
	Virtinit().VirtualMachineShutoff(core.VMid)
    },
  }
        return subvirtualMachineShutoff
}

func (f *CobraMenu) RootSubCmdvirtualMachineStart() *cobra.Command {
  var subvirtualMachineStart = &cobra.Command{
    Use:   "MachineStart [--id vmname!]",
    Short: "starts up a VM. Returns result with a current machine state",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdvirtualMachineStart Run with args")
	Virtinit().VirtualMachineStart(core.VMid)
    },
  }
        return subvirtualMachineStart
}

func (f *CobraMenu) RootSubCmdvirtualMachinePause() *cobra.Command {
  var subvirtualMachinePause = &cobra.Command{
    Use:   "MachinePause [--id vmname!]",
    Short: "stops the execution of the VM. CPU is not used, but memory is still occupied.",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdvirtualMachinePause Run with args")
	Virtinit().VirtualMachinePause(core.VMid)
    },
  }
        return subvirtualMachinePause
}

func (f *CobraMenu) RootSubCmdvirtualMachineResume() *cobra.Command {
  var subvirtualMachineResume = &cobra.Command{
    Use:   "MachineResume [--id vmname!]",
    Short: "called after Pause, to resume the invocation of the VM. Returns result with a current machine state",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdvirtualMachineResume Run with args")
	Virtinit().VirtualMachineResume(core.VMid)
    },
  }
        return subvirtualMachineResume
}

func (f *CobraMenu) RootSubCmdvirtualMachineCreate() *cobra.Command {
  var subvirtualMachineCreate = &cobra.Command{
    Use:   "MachineCreate [--xml filename!]",
    Short: "creates a new machine. Requires --xml parameter. Returns result with a current machine state",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.XmlTemplate": core.XmlTemplate, "args": args,}).Info("Inside RootSubCmdvirtualMachineCreate Run with args")
	Virtinit().VirtualMachineCreate(core.XmlTemplate)
    },
  }
        return subvirtualMachineCreate
}

func (f *CobraMenu) RootSubCmdvirtualMachineDelete() *cobra.Command {
  var subvirtualMachineDelete = &cobra.Command{
    Use:   "MachineDelete [--id vmname!]",
    Short: "deletes an existing machine.",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdvirtualMachineDelete Run with args")
        Virtinit().VirtualMachineDelete(core.VMid)
    },
  }
        return subvirtualMachineDelete
}


func (f *CobraMenu) RootSubCmdvirtualMachineDestroy() *cobra.Command {
        var subvirtualMachineDestroy = &cobra.Command{
          Use:   "MachineDestroy [--id VMAname]",
          Short: "Destroy an existing machine.",
          Run: func(cmd *cobra.Command, args []string) {

		core := Singleton[Core]()
		log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdvirtualMachineDestroy Run with args")

		core = Singleton[Core]()
                Virtinit().VirtualMachineShutoff(core.VMid)

                core = Singleton[Core]()
                Virtinit().VirtualMachineDelete(core.VMid)
          },
        }
        return subvirtualMachineDestroy
}


func initVM(increment int) {

        core := Singleton[Core]()

	log.WithFields(logrus.Fields{ "core": core,}).Info("Inside RootSubCmdinitVM")

	var c *LibVirtVM
	c = LoadConfigVM( core, increment)
	// ****************************************************************
	c.CreateNetworkConfig()
	// ****************************************************************
	c.CreateUserData()
	// ****************************************************************
	c.PreInitScriptVM()
	// ****************************************************************
	//c.AfterDeployVM()
	// ****************************************************************
	c.CreateVMXML()
	// ****************************************************************

	core = Singleton[Core]()
	log.Info("Run Virtinit().VirtualMachineCreate(core.XmlTemplate) in initVM")
	Virtinit().VirtualMachineCreate(c.XmlTemplate)
	log.Info("Run Virtinit().VirtualMachineCreate(core.XmlTemplate) in initVM Done!")

        core = Singleton[Core]()
        Virtinit().VirtualMachineState(c.VMNAME)

        core = Singleton[Core]()
        Virtinit().VirtualMachineStart(c.VMNAME)

        core = Singleton[Core]()
        Virtinit().VirtualMachineState(c.VMNAME)

	stop := 0
	for stop==0 {

		log.WithFields(logrus.Fields{ "c.VMNAME": c.VMNAME, }).Info("Sleep 8sec.....")
		time.Sleep(8 * time.Second)

		stopCodes, err := Virtinit().Libvirt.DomainState(c.VMNAME)
		if err != nil {
			log.WithFields(logrus.Fields{ "err": err, "c.VMNAME": c.VMNAME, }).Info("Cobra Event Error")
			return
		}
		stopCode, err := strconv.Atoi(fmt.Sprintf("%d",stopCodes))
		if err != nil {
			log.WithFields(logrus.Fields{ "err": err, "c.VMNAME": c.VMNAME, }).Info("Cobra Event Error")
			return
		}
		log.WithFields(logrus.Fields{ "DomainState": DomainState[stopCode], "VMStatusCODE": stopCode, "c.VMNAME": c.VMNAME, }).Info("Inside RootSubCmdinitVM Run with args")
		if DomainState[stopCode]=="DomainShutoff"{
        		stop=1
	    	}
	}
	log.WithFields(logrus.Fields{ "VM Status": c.VMNAME,}).Info("VM is DomainShutoff")

        core = Singleton[Core]()
        Virtinit().VirtualMachineStart(c.VMNAME)

	log.WithFields(logrus.Fields{ "VM Status": c.VMNAME,}).Info("Starting VM")

	stop = 0
        for stop==0 {

		log.WithFields(logrus.Fields{ "c.VMNAME": c.VMNAME, }).Info("Sleep 8sec.....")
                time.Sleep(8 * time.Second)
                stopCodes, err := Virtinit().Libvirt.DomainState(c.VMNAME)
		if err != nil {
			log.WithFields(logrus.Fields{ "err": err, "c.VMNAME": c.VMNAME, }).Info("Cobra Event Error")
			return
		}
                stopCode, err := strconv.Atoi(fmt.Sprintf("%d",stopCodes))
		if err != nil {
			log.WithFields(logrus.Fields{ "err": err, "c.VMNAME": c.VMNAME, }).Info("Cobra Event Error")
			return
		}
		log.WithFields(logrus.Fields{ "VM Status": DomainState[stopCode], "VMStatusCode": stopCode, "c.VMNAME": c.VMNAME, }).Info("Status VM")
                if DomainState[stopCode]=="DomainRunning"{
                        stop = 1
                }
        }
	log.WithFields(logrus.Fields{ "VMName": c.VMNAME,}).Info("VM is Running")

	stop = 0
        for stop==0 {
		log.WithFields(logrus.Fields{ "c.VMNAME": c.VMNAME, }).Info("Sleep 8sec.....")
                time.Sleep(8 * time.Second)
		log.WithFields(logrus.Fields{ "c.VMNAME": c.VMNAME, }).Info("Check SSH port....\n")
        	testConnect :=  Ping("tcp",fmt.Sprintf("192.168.200.%s:22",c.NodeId), 5 * time.Second)
		if testConnect==nil {
			stop = 1
		}
	}

	log.WithFields(logrus.Fields{ "c.VMNAME": c.VMNAME, }).Info("Sleep 20sec.....")
        time.Sleep(20 * time.Second)



	doms, err := Virtinit().Libvirt.Domains()
	if err != nil {
		log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
		return
	}

	for _,j := range doms {
		if j.Name == c.VMNAME {
			dom := j

                        CDDiskTemplate, err := c.ReadFile( &c.Config.CDDiskTemplate)
                        if err != nil {
                                log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                                return
                        }

			disk_seed := fmt.Sprintf(*CDDiskTemplate,c.VMDisk[1].Path)
			log.WithFields(logrus.Fields{ "VMName": c.VMNAME, "XML": disk_seed,}).Info("VM change")
			err = Virtinit().Libvirt.DomainDetachDeviceFlags(dom, disk_seed, uint32(libvirt.DomainDeviceModifyCurrent | libvirt.DomainDeviceModifyConfig | libvirt.DomainDeviceModifyLive))
			if err != nil {
				log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
				return
			}

			c.AfterDeployVM()

			log.WithFields(logrus.Fields{ "c.VMNAME": c.VMNAME, }).Info("Sleep 10sec.....")
			time.Sleep(10 * time.Second)
			log.WithFields(logrus.Fields{ "VMName": c.VMNAME,}).Info("Reboot VM")
			Virtinit().VirtualMachineHardReboot(c.VMNAME)

			stop = 0
			for stop==0 {
				log.WithFields(logrus.Fields{ "c.VMNAME": c.VMNAME, }).Info("Sleep 10sec.....")
				time.Sleep(8 * time.Second)
				log.WithFields(logrus.Fields{ "c.VMNAME": c.VMNAME, }).Info("Check SSH port....")
				testConnect :=  Ping("tcp",fmt.Sprintf("192.168.200.%s:22",c.NodeId), 5 * time.Second)
				if testConnect==nil {
					stop = 1
				}
			}
			log.WithFields(logrus.Fields{ "c.VMNAME": c.VMNAME, }).Info("Sleep 10sec.....")
		        time.Sleep(10 * time.Second)
			log.Info("Done")
		}
	}
}


func workerInitVM(increment int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", increment)
	initVM(increment)
	fmt.Printf("Worker %d done\n", increment)
	wg.Done()
}

func (f *CobraMenu) RootSubCmdinitVM() *cobra.Command {
  var subinitVM = &cobra.Command{
    Use:   "initVM [--help]",
    Short: "init VM",
    Run: func(cmd *cobra.Command, args []string) {

        core := Singleton[Core]()

        log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdinitVM Run with args")

        log.WithFields(logrus.Fields{ "core.VMNAME": core.VMNAME, "core.CORE": core.CORE, "core.MEMORY": core.MEMORY,}).Info("Inside RootSubCmdinitVM")
        log.WithFields(logrus.Fields{ "core.ROOTFS_SIZE": core.ROOTFS_SIZE, "core.Octet": core.Octet, "core.EXT_DISK_SIZE": core.EXT_DISK_SIZE,}).Info("Inside RootSubCmdinitVM")
        log.WithFields(logrus.Fields{ "core.USER_DATA_PATH": core.USER_DATA_PATH,}).Info("Inside RootSubCmdinitVM")


        log.WithFields(logrus.Fields{ "core": core,}).Info("Inside RootSubCmdinitVM")

	numVM := 1
	var wg sync.WaitGroup
	for i := 0; i <= numVM-1; i++ {
		wg.Add(1)
		go workerInitVM(i, &wg)
	}
	wg.Wait()

    },
  }
	return subinitVM
}






func (f *CobraMenu) RootSubCmdinitVMs() *cobra.Command {
  var subinitVMs = &cobra.Command{
    Use:   "initVMs [--id vmname!]",
    Short: "init VMs",
    Run: func(cmd *cobra.Command, args []string) {

        core := Singleton[Core]()

        log.WithFields(logrus.Fields{ "core.VMid": core.VMid, "args": args,}).Info("Inside RootSubCmdinitVM Run with args")

        log.WithFields(logrus.Fields{ "core.VMNAME": core.VMNAME, "core.CORE": core.CORE, "core.MEMORY": core.MEMORY,}).Info("Inside RootSubCmdinitVM")
        log.WithFields(logrus.Fields{ "core.ROOTFS_SIZE": core.ROOTFS_SIZE, "core.Octet": core.Octet, "core.EXT_DISK_SIZE": core.EXT_DISK_SIZE,}).Info("Inside RootSubCmdinitVM")
        log.WithFields(logrus.Fields{ "core.USER_DATA_PATH": core.USER_DATA_PATH,}).Info("Inside RootSubCmdinitVM")


        log.WithFields(logrus.Fields{ "core": core,}).Info("Inside RootSubCmdinitVM")

        var wg sync.WaitGroup
        for i := 0; i <= core.NumVM-1; i++ {
                wg.Add(1)
                go workerInitVM(i, &wg)
        }
        wg.Wait()


    },
  }
        return subinitVMs
}

func (f *CobraMenu) RootSubCmdattachDiskVM() *cobra.Command {
  var subattachDiskVM = &cobra.Command{
    Use:   "attachDiskVM [--help]",
    Short: "attach Disk VM",
    Run: func(cmd *cobra.Command, args []string) {

        core := Singleton[Core]()

        log.WithFields(logrus.Fields{ "core.VMNAME": core.VMNAME, "core.CORE": core.CORE, "core.MEMORY": core.MEMORY,}).Info("Inside RootSubCmdattachDiskVM")
        log.WithFields(logrus.Fields{ "core.ROOTFS_SIZE": core.ROOTFS_SIZE, "core.Octet": core.Octet, "core.EXT_DISK_SIZE": core.EXT_DISK_SIZE,}).Info("Inside RootSubCmdattachDiskVM")
        log.WithFields(logrus.Fields{ "core.USER_DATA_PATH": core.USER_DATA_PATH,}).Info("Inside RootSubCmdattachDiskVM")


        log.WithFields(logrus.Fields{ "core": core,}).Info("Inside RootSubCmdattachDiskVM")

        var c *LibVirtVM
        c = LoadConfigVM(core,0)

        doms, err := Virtinit().Libvirt.Domains()
	if err != nil {
		log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
		return
	}

        for _,j := range doms {
                if j.Name == core.VMNAME {

                        dom := j
			fmt.Printf("%#v\n", dom);

                        domainGetXMLDesc, err := Virtinit().Libvirt.DomainGetXMLDesc(j, libvirt.DomainXMLSecure)
			if err != nil {
				log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
				return
			}

                        t := Domain{}
                        xml.Unmarshal([]byte(domainGetXMLDesc), &t)

			useDisk := make([]string, len(t.Devices.Disk))
                        for k,v := range t.Devices.Disk {
				log.WithFields(logrus.Fields{ "key": k, "v.Type": v.Type, "v.Device": v.Device, "v.Source.File": v.Source.File, "v.Target.Dev": v.Target.Dev, "v.Target.Bus": v.Target.Bus, }).Info("RootSubCmdattachDiskVM")
				useDisk[k] = v.Target.Dev
                        }


			c.DISKID = SearchDisk(useDisk)
			c.EXT_DISK_SIZE = core.EXT_DISK_SIZE

			SCSIDiskTemplate, err := c.ReadFile( &c.Config.SCSIDiskTemplate)
                        if err != nil {
                                log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                                return
                        }


			renderXML, err := templateRender( *SCSIDiskTemplate, *c)
			if err != nil {
				log.WithFields(logrus.Fields{ "err": err, }).Info("RootSubCmdattachDiskVM")
				return
			}
			log.WithFields(logrus.Fields{ "renderXML": renderXML, }).Info("RootSubCmdattachDiskVM")

			cmdtmpl := `qemu-img create -f qcow2 {{.Config.VMPATH}}/{{.VMNAME}}/{{.VMNAME}}-disk{{.DISKID}}.qcow2 {{.EXT_DISK_SIZE}}G`
			renderCmd, err := templateRender(cmdtmpl, *c)
			if err != nil {
				log.WithFields(logrus.Fields{ "err": err, }).Info("RootSubCmdattachDiskVM")
				return
                        }
			log.WithFields(logrus.Fields{ "renderCmd": renderCmd, }).Info("RootSubCmdattachDiskVM")

                        out, errout, err := c.Shellout(renderCmd)
                        if err != nil {
                        	log.WithFields(logrus.Fields{ "err": err, }).Info("RootSubCmdattachDiskVM run renderCmd")
                        	return
                        }

			log.WithFields(logrus.Fields{ "out": out, }).Info("RootSubCmdattachDiskVM run renderCmd")
			log.WithFields(logrus.Fields{ "errout": errout, }).Info("RootSubCmdattachDiskVM run renderCmd")
                        log.WithFields(logrus.Fields{ "VMName": core.VMNAME, "XML": renderXML,}).Info("VM change")

			fmt.Printf("%s\n",renderXML)
			fmt.Printf("%#v\n", dom)
                        err = Virtinit().Libvirt.DomainAttachDeviceFlags(dom, renderXML, uint32(libvirt.DomainDeviceModifyCurrent | libvirt.DomainDeviceModifyConfig | libvirt.DomainDeviceModifyLive))

			if err != nil {
				log.WithFields(logrus.Fields{ "err": err, }).Info("RootSubCmdattachDiskVM run renderCmd")
				return
			}

                }
        }


    },
  }
        return subattachDiskVM
}


func (f *CobraMenu) RootSubCmdDetachDiskVM() *cobra.Command {
  var subattachDiskVM = &cobra.Command{
    Use:   "detachDiskVM [--help]",
    Short: "detach Disk VM",
    Run: func(cmd *cobra.Command, args []string) {

        core := Singleton[Core]()

        log.WithFields(logrus.Fields{ "core.VMNAME": core.VMNAME, "core.CORE": core.CORE, "core.MEMORY": core.MEMORY,}).Info("Inside RootSubCmdDetachDiskVM")
        log.WithFields(logrus.Fields{ "core.ROOTFS_SIZE": core.ROOTFS_SIZE, "core.Octet": core.Octet, "core.EXT_DISK_SIZE": core.EXT_DISK_SIZE,}).Info("Inside RootSubCmdDetachDiskVM")
        log.WithFields(logrus.Fields{ "core.USER_DATA_PATH": core.USER_DATA_PATH,"core.DetachDiskName": core.DetachDiskName,}).Info("Inside RootSubCmdDetachDiskVM")

        log.WithFields(logrus.Fields{ "core": core,}).Info("Inside RootSubCmdDetachDiskVM")

        var c *LibVirtVM
        c = LoadConfigVM(core,0)

        doms, err := Virtinit().Libvirt.Domains()
        if err != nil {
                log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                return
        }

        for _,j := range doms {
                if j.Name == core.VMNAME {

                        dom := j
                        fmt.Printf("%#v\n", dom);

                        domainGetXMLDesc, err := Virtinit().Libvirt.DomainGetXMLDesc(j, libvirt.DomainXMLSecure)
                        if err != nil {
                                log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                                return
                        }

                        t := Domain{}
                        xml.Unmarshal([]byte(domainGetXMLDesc), &t)
                        for k,v := range t.Devices.Disk {
				log.WithFields(logrus.Fields{ "k": k, "v.Type": v.Type, "v.Device": v.Device, "v.Source.File": v.Source.File, "v.Target.Dev": v.Target.Dev, "v.Target.Bus": v.Target.Bus, }).Info("RootSubCmdDetachDiskVM")
                        }

                        c.DISKID = core.DetachDiskName

                        SCSIDiskTemplate, err := c.ReadFile( &c.Config.SCSIDiskTemplate)
                        if err != nil {
                                log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                                return
                        }

                        renderXML, err := templateRender( *SCSIDiskTemplate, *c)
                        if err != nil {
                                log.WithFields(logrus.Fields{ "err": err, }).Info("RootSubCmdDetachDiskVM")
                                return
                        }
                        log.WithFields(logrus.Fields{ "renderXML": renderXML, }).Info("RootSubCmdDetachDiskVM")

                        fmt.Printf("%s\n",renderXML)
                        fmt.Printf("%#v\n", dom)

                        err = Virtinit().Libvirt.DomainDetachDeviceFlags(dom, renderXML, uint32(libvirt.DomainDeviceModifyCurrent | libvirt.DomainDeviceModifyConfig | libvirt.DomainDeviceModifyLive))
                        if err != nil {
                                log.WithFields(logrus.Fields{ "err": err, }).Info("RootSubCmdDetachDiskVM run renderCmd")
                                return
                        }
                }
        }


    },
  }
        return subattachDiskVM
}

func (f *CobraMenu) RootSubCmdattachInerfaceVM() *cobra.Command {
  var subattachInerfaceVM = &cobra.Command{
    Use:   "attachInerfaceVM [--id vmname!]",
    Short: "attach Inerface VM",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()

        log.WithFields(logrus.Fields{ "core.VMNAME": core.VMNAME, "core.CORE": core.CORE, "core.MEMORY": core.MEMORY,}).Info("Inside RootSubCmdattachInerfaceVM")
        log.WithFields(logrus.Fields{ "core.ROOTFS_SIZE": core.ROOTFS_SIZE, "core.Octet": core.Octet, "core.EXT_DISK_SIZE": core.EXT_DISK_SIZE,}).Info("Inside RootSubCmdattachInerfaceVM")
        log.WithFields(logrus.Fields{ "core.USER_DATA_PATH": core.USER_DATA_PATH,}).Info("Inside RootSubCmdattachInerfaceVM")


        log.WithFields(logrus.Fields{ "core": core,}).Info("Inside RootSubCmdattachInerfaceVM")

        var c *LibVirtVM
        c = LoadConfigVM(core,0)


        doms, err := Virtinit().Libvirt.Domains()
	if err != nil {
		log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
		return
	}

        for _,j := range doms {
                if j.Name == core.VMNAME {
                        dom := j


                        domainGetXMLDesc, err := Virtinit().Libvirt.DomainGetXMLDesc(j, libvirt.DomainXMLSecure)
			if err != nil {
				log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
				return
			}

                        fmt.Printf("%#v\n", domainGetXMLDesc);

                        t := Domain{}
                        xml.Unmarshal([]byte(domainGetXMLDesc), &t)

			pciSlotUse := []string{}
                        for _, v := range t.Devices.Interface {
                                log.WithFields(logrus.Fields{ "v.Type": v.Type,}).Info("VM interface id")
                                log.WithFields(logrus.Fields{ "v.Mac.Address": v.Mac.Address,}).Info("VM Interface.Type")
                                log.WithFields(logrus.Fields{ "v.Source.Network": v.Source.Network,}).Info("VM Interface.Mac.Address")
                                log.WithFields(logrus.Fields{ "v.Model.Type": v.Model.Type,}).Info("VM Interface.Source.Network")
                                log.WithFields(logrus.Fields{ "v.Address.Type": v.Address.Type,}).Info("VM Interface.Model.Type")
                                log.WithFields(logrus.Fields{ "v.Address.Type": v.Address.Type,}).Info("VM Interface.Address.Type")
                                log.WithFields(logrus.Fields{ "v.Address.Domain": v.Address.Domain,}).Info("VM Interface.Address.Domain")
                                log.WithFields(logrus.Fields{ "v.Address.Bus": v.Address.Bus,}).Info("VM Interface.Address.Bus")
                                log.WithFields(logrus.Fields{ "v.Address.Slot": v.Address.Slot,}).Info("VM Interface.Address.Slot")
                                log.WithFields(logrus.Fields{ "v.Address.Function": v.Address.Function,}).Info("VM Interface.Address.Function")
                                log.WithFields(logrus.Fields{ "v.Address.Multifunction": v.Address.Multifunction,}).Info("VM Interface.Address.Multifunction")

                                 if v.Address.Bus=="0x03" {
                                        pciSlotUse = append( pciSlotUse, v.Address.Slot)
                                }
                        }

			log.WithFields(logrus.Fields{ "pciSlotUse": pciSlotUse, }).Info("RootSubCmdattachInerfaceVM run renderCmd")
			freePciSlot := SearchSlot(pciSlotUse)
			log.WithFields(logrus.Fields{ "freePciSlot": freePciSlot, }).Info("RootSubCmdattachInerfaceVM run renderCmd")

			interfaceVM := InterfaceName{}
			interfaceVM.CreateMacAddress(c.Network.MagicMac)
			fmt.Printf("%s\n", interfaceVM.MacAddress)
			interfaceVM.Slot = fmt.Sprintf("0x%02x",  freePciSlot)
			fmt.Printf("%s\n", interfaceVM.Slot)

                        E1000Networkemplate, err := c.ReadFile( &c.Config.E1000Networkemplate)
                        if err != nil {
                                log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                                return
                        }

                        renderInterface, err := templateRender( *E1000Networkemplate, interfaceVM)
                        if err != nil {
                            log.WithFields(logrus.Fields{ "err": err, }).Info("CreateUserData")
                            return
                        }

                        log.WithFields(logrus.Fields{ "VMName": core.VMNAME, "XML": renderInterface,}).Info("VM change")
                        err = Virtinit().Libvirt.DomainAttachDeviceFlags(dom, renderInterface, uint32(libvirt.DomainDeviceModifyCurrent | libvirt.DomainDeviceModifyConfig | libvirt.DomainDeviceModifyLive))
                        if err != nil {
				log.WithFields(logrus.Fields{ "err": err, }).Info("RootSubCmdattachInerfaceVM run renderCmd")
				return
                        }

                }
        }

    },
  }
        return subattachInerfaceVM
}


func (f *CobraMenu) RootSubCmdaDetachInerfaceVM() *cobra.Command {
  var subattachInerfaceVM = &cobra.Command{
    Use:   "detachInerfaceVM [--id vmname!]",
    Short: "detach Inerface VM",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()

        log.WithFields(logrus.Fields{ "core.VMNAME": core.VMNAME, "core.CORE": core.CORE, "core.MEMORY": core.MEMORY,}).Info("Inside RootSubCmdDetachInerfaceVM")
        log.WithFields(logrus.Fields{ "core.ROOTFS_SIZE": core.ROOTFS_SIZE, "core.Octet": core.Octet, "core.EXT_DISK_SIZE": core.EXT_DISK_SIZE,}).Info("Inside RootSubCmdDetachInerfaceVM")
        log.WithFields(logrus.Fields{ "core.USER_DATA_PATH": core.USER_DATA_PATH, "core.NetworkSlot": core.NetworkSlot, }).Info("Inside RootSubCmdDetachInerfaceVM")


        log.WithFields(logrus.Fields{ "core": core,}).Info("Inside RootSubCmdDetachInerfaceVM")

        var c *LibVirtVM
        c = LoadConfigVM(core,0)


        doms, err := Virtinit().Libvirt.Domains()
        if err != nil {
                log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                return
        }

        for _,j := range doms {
                if j.Name == core.VMNAME {
                        dom := j


                        domainGetXMLDesc, err := Virtinit().Libvirt.DomainGetXMLDesc(j, libvirt.DomainXMLSecure)
                        if err != nil {
                                log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                                return
                        }

                        fmt.Printf("%#v\n", domainGetXMLDesc);

                        t := Domain{}
                        xml.Unmarshal([]byte(domainGetXMLDesc), &t)

                        for _, v := range t.Devices.Interface {
                                log.WithFields(logrus.Fields{ "v.Type": v.Type,}).Info("VM interface id")
                                log.WithFields(logrus.Fields{ "v.Mac.Address": v.Mac.Address,}).Info("VM Interface.Type")
                                log.WithFields(logrus.Fields{ "v.Source.Network": v.Source.Network,}).Info("VM Interface.Mac.Address")
                                log.WithFields(logrus.Fields{ "v.Model.Type": v.Model.Type,}).Info("VM Interface.Source.Network")
                                log.WithFields(logrus.Fields{ "v.Address.Type": v.Address.Type,}).Info("VM Interface.Model.Type")
                                log.WithFields(logrus.Fields{ "v.Address.Type": v.Address.Type,}).Info("VM Interface.Address.Type")
                                log.WithFields(logrus.Fields{ "v.Address.Domain": v.Address.Domain,}).Info("VM Interface.Address.Domain")
                                log.WithFields(logrus.Fields{ "v.Address.Bus": v.Address.Bus,}).Info("VM Interface.Address.Bus")
                                log.WithFields(logrus.Fields{ "v.Address.Slot": v.Address.Slot,}).Info("VM Interface.Address.Slot")
                                log.WithFields(logrus.Fields{ "v.Address.Function": v.Address.Function,}).Info("VM Interface.Address.Function")
                                log.WithFields(logrus.Fields{ "v.Address.Multifunction": v.Address.Multifunction,}).Info("VM Interface.Address.Multifunction")

                                if v.Address.Bus=="0x03" && core.NetworkSlot == v.Address.Slot{
                                        interfaceVM := InterfaceName{}
                                        interfaceVM.MacAddress = v.Mac.Address
                                        interfaceVM.Slot = v.Address.Slot
                                        log.WithFields(logrus.Fields{ "interfaceVM": interfaceVM, }).Info("interface")


					E1000Networkemplate, err := c.ReadFile( &c.Config.E1000Networkemplate)
					if err != nil {
						log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
						return
					}

                                        renderInterface, err := templateRender( *E1000Networkemplate, interfaceVM)
                                        if err != nil {
                                                log.WithFields(logrus.Fields{ "err": err, }).Info("RootSubCmdDetachInerfaceVM")
                                                return
                                        }

                                        log.WithFields(logrus.Fields{ "VMName": core.VMNAME, "XML": renderInterface,}).Info("VM change")
                                        err = Virtinit().Libvirt.DomainDetachDeviceFlags(dom, renderInterface, uint32(libvirt.DomainDeviceModifyCurrent | libvirt.DomainDeviceModifyConfig | libvirt.DomainDeviceModifyLive))
                                        if err != nil {
                                                log.WithFields(logrus.Fields{ "err": err, }).Info("RootSubCmdDetachInerfaceVM run renderCmd")
                                                return
                                        }
                                }
                        }
                }
        }

    },
  }
        return subattachInerfaceVM
}


func (f *CobraMenu) RootSubCmddestroyVM() *cobra.Command {
  var subdestroyVM = &cobra.Command{
    Use:   "destroyVM [--VMNAME <name>]",
    Short: "destroy VM",
    Run: func(cmd *cobra.Command, args []string) {

		core := Singleton[Core]()
		log.WithFields(logrus.Fields{ "core.VMNAME": core.VMNAME,}).Info("Inside RootSubCmddestroyVM")


		core = Singleton[Core]()
		Virtinit().VirtualMachineShutoff(core.VMNAME)

		core = Singleton[Core]()
		Virtinit().VirtualMachineDelete(core.VMNAME)

    },
  }
        return subdestroyVM
}

func (f *CobraMenu) RootSubCmdlistInterfaceVM() *cobra.Command {
  var sublistInterfaceVM = &cobra.Command{
    Use:   "listInterfaceVM [--id vmname!]",
    Short: "list Interface VM",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.VMid": core.VMid,}).Info("RootSubCmdlistInterfaceVM")
    },
  }
        return sublistInterfaceVM
}


func (f *CobraMenu) RootSubCmdlistDiskVM() *cobra.Command {
  var sublistDiskVM = &cobra.Command{
    Use:   "listDiskVM [--id vmname!]",
    Short: "list Disk VM",
    Run: func(cmd *cobra.Command, args []string) {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core.VMid": core.VMid,}).Info("RootSubCmdlistDiskVM")
    },
  }
        return sublistDiskVM
}



func GetCobraMenu(name string) *cobra.Command {

	CobraMenu := CobraMenu{}
	CobraMenuType := reflect.TypeOf(&CobraMenu)
	CobraMenuValue := reflect.ValueOf(&CobraMenu)

	method, err := CobraMenuType.MethodByName(name)
	if err !=true {
		return nil
	}

	val := CobraMenuValue.MethodByName(method.Name).Call([]reflect.Value{})

	if len(val)==1 {
		return val[0].Interface().(*cobra.Command)
	}
	return nil
}

