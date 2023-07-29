package main

import (

	"fmt"
	"time"
        "strconv"
        "io/ioutil"
        "encoding/xml"

        "github.com/digitalocean/go-libvirt"
        "github.com/digitalocean/go-libvirt/socket/dialers"

        "github.com/sirupsen/logrus"
)



func (ret *VirtualMachine) RootSubCmdvirtualMachineMachineState() {
        core := Singleton[Core]()
	log.WithFields(logrus.Fields{ "core": core,}).Info("Inside RootSubCmdvirtualMachineState Run with args")
	Virtinit().VirtualMachineState(core.VMid)

        doms, err := Virtinit().Libvirt.Domains()
        if err != nil {
                log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                return
        }

        for _,j := range doms {
                if j.Name == core.VMid {

                        domainGetXMLDesc, err := Virtinit().Libvirt.DomainGetXMLDesc(j, libvirt.DomainXMLSecure)
                        if err != nil {
                                log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                                return
                        }

                        t := Domain{}
                        xml.Unmarshal([]byte(domainGetXMLDesc), &t)
			log.Info("--------------------------------------------------")
                        for k,v := range t.Devices.Disk {
				log.WithFields(logrus.Fields{	"key": k, "v.Type": v.Type, "v.Device": v.Device, "v.Source.File": v.Source.File,
								"v.Target.Dev": v.Target.Dev, "v.Target.Bus": v.Target.Bus, }).Info("RootSubCmdvirtualMachineState")
				log.Info("=========================================================")
                        }
			log.Info("--------------------------------------------------")
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
				log.Info("=========================================================")
                        }
			log.Info("--------------------------------------------------")
                }
        }
}

func (ret *VirtualMachine) RootSubCmdvirtualMachineMigrate() {

        core := Singleton[Core]()
        log.WithFields(logrus.Fields{ "core": core,}).Info("Inside RootSubCmdvirtualMachineStart Run with args")

        domains, err := ret.Libvirt.Domains()
        if err != nil {
            log.Fatalf("failed to retrieve domains: %v", err)
        }

        fmt.Println("ID\tName\t\tUUID\t\t\t\tStatus")
        fmt.Printf("---------------------------------------------------------------------------------------------------------------------------------\n")
        for _, d := range domains {
                a, _ := ret.Libvirt.DomainState(d.Name)
                c, _ := strconv.Atoi(fmt.Sprintf("%d",a))
                fmt.Printf("%#v\n",d);
                fmt.Printf("%d\t%s\t%#s\n", d.ID, d.Name, DomainState[c])
        }

        flags := libvirt.MigrateLive |
		            libvirt.MigratePeer2peer |
		            libvirt.MigratePersistDest |
		            libvirt.MigrateChangeProtection |
		            libvirt.MigrateAbortOnError |
		            libvirt.MigrateAutoConverge |
		            libvirt.MigrateNonSharedDisk

        dom, err := ret.Libvirt.DomainLookupByName(core.VMid)
        if err != nil {
            log.Fatalf("failed to lookup domain: %v", err)
        }
        dconnuri := []string{ fmt.Sprintf("qemu+ssh://root@%s/system",core.ToMove)}
        if e, err := ret.Libvirt.DomainMigratePerform3Params( dom, dconnuri,
                                                              []libvirt.TypedParam{}, []byte{}, flags); err != nil {
		    log.Fatalf("unexpected live migration error: %v", err)
        } else {
            fmt.Printf("%#v\n",e);
        }

        domains, err = ret.Libvirt.Domains()
        if err != nil {
                log.Fatalf("failed to retrieve domains: %v", err)
        }

        fmt.Println("ID\tName\t\tUUID\t\t\t\tStatus")
        fmt.Printf("---------------------------------------------------------------------------------------------------------------------------------\n")
        for _, d := range domains {
                a, _ := ret.Libvirt.DomainState(d.Name)
                c, _ := strconv.Atoi(fmt.Sprintf("%d",a))
                fmt.Printf("%#v\n",d);
                fmt.Printf("%d\t%s\t%x\t%#s\n", d.ID, d.Name, d.UUID, DomainState[c])
        }


}


// VirtualMachineState returns current state of a virtual machine.
func (ret *VirtualMachine) VirtualMachineState(id string) {
//        var ret VirtualMachine //VirtualMachineState

        d, err := ret.Libvirt.DomainLookupByName(id)
	log.WithFields(logrus.Fields{ "DomainLookupByName": d, "err": err, "id": id, }).Info("Inside VirtualMachineState Run")
        herr(err)

        state, maxmem, mem, ncpu, cputime, err := ret.Libvirt.DomainGetInfo(d)
	log.WithFields(logrus.Fields{ "state": state, "maxmem": maxmem, "mem": mem, "ncpu": ncpu, "cputime": cputime, "err": err, }).Info("Inside VirtualMachineState Run")
        herr(err)

        ret.CPUCount = ncpu
        ret.CPUTime = cputime
        // God only knows why they return memory in kilobytes.
        ret.MemoryBytes = mem * 1024
        ret.MaxMemoryBytes = maxmem * 1024
        temp := libvirt.DomainState(state)
        herr(err)

        switch temp {
        case libvirt.DomainNostate:
                ret.State = VirtStatePending
        case libvirt.DomainRunning:
                ret.State = VirtStateRunning
        case libvirt.DomainBlocked:
                ret.State = VirtStateBlocked
        case libvirt.DomainPaused:
                ret.State = VirtStatePaused
        case libvirt.DomainShutdown:
                ret.State = VirtStateShutdown
        case libvirt.DomainShutoff:
                ret.State = VirtStateShutoff
        case libvirt.DomainCrashed:
                ret.State = VirtStateCrashed
        case libvirt.DomainPmsuspended:
                ret.State = VirtStateHybernating
        }
        ret.Libvirt.DomainGetState(d, 0)
        hret(ret)
}



// VirtualMachineCreate creates a new VM from an xml template file
func (ret *VirtualMachine) VirtualMachineCreate(xmlTemplate string) {

        xml, err := ioutil.ReadFile(xmlTemplate)
        herr(err)

        d, err := ret.Libvirt.DomainDefineXML(string(xml))
        herr(err)

        hret(d)
}

// VirtualMachineDelete deletes a new VM from an xml template file
func (ret *VirtualMachine) VirtualMachineDelete(id string) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        herr(err)
        err = ret.Libvirt.DomainUndefineFlags(d, libvirt.DomainUndefineKeepNvram)
        herr(err)
        hok(fmt.Sprintf("%v was deleted", id))
}

// VirtualMachineSoftReboot reboots a machine gracefully, as chosen by hypervisor.
func (ret *VirtualMachine) VirtualMachineSoftReboot(id string) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        herr(err)

        err = ret.Libvirt.DomainReboot(d, libvirt.DomainRebootDefault)
        herr(err)

        hok(fmt.Sprintf("%v was soft-rebooted successfully", id))
}

// VirtualMachineHardReboot sends a VM into hard-reset mode. This is damaging to all ongoing file operations.
func (ret *VirtualMachine) VirtualMachineHardReboot(id string) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        herr(err)

        err = ret.Libvirt.DomainReset(d, 0)
        herr(err)

        hok(fmt.Sprintf("%v was hard-rebooted successfully", id))
}


// VirtualMachineShutdown gracefully shuts down the VM.
func (ret *VirtualMachine) VirtualMachineShutdown(id string) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        herr(err)

        err = ret.Libvirt.DomainShutdown(d)
        herr(err)

        hok(fmt.Sprintf("%v was shutdown successfully", id))
}

// VirtualMachineShutoff kills running VM. Equivalent to pulling a plug out of a computer.
func (ret *VirtualMachine) VirtualMachineShutoff(id string) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        herr(err)

        err = ret.Libvirt.DomainDestroy(d)
        herr(err)

        hok(fmt.Sprintf("%v was shutoff successfully", id))
}

// VirtualMachineStart starts up a VM.
func (ret *VirtualMachine) VirtualMachineStart(id string) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        herr(err)

        //v.DomainRestore()
        //_, err = ret.Libvirt.DomainCreateWithFlags(d, uint32(libvirt.DomainStartBypassCache))
        err = ret.Libvirt.DomainCreate(d)

        herr(err)

        hok(fmt.Sprintf("%v was started", id))
}

// VirtualMachinePause stops the execution of the VM. CPU is not used, but memory is still occupied.
func (ret *VirtualMachine) VirtualMachinePause(id string) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        herr(err)

        err = ret.Libvirt.DomainSuspend(d)
        herr(err)

        hok(fmt.Sprintf("%v is paused", id))
}

// VirtualMachineResume can be called after Pause, to resume the invocation of the VM.
func (ret *VirtualMachine) VirtualMachineResume(id string) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        herr(err)

        err = ret.Libvirt.DomainResume(d)
        herr(err)

        hok(fmt.Sprintf("%v was resumed", id))
}


func Virtinit() *VirtualMachine {
	v := VirtualMachine{}
        v.Libvirt = libvirt.NewWithDialer(dialers.NewLocal(dialers.WithLocalTimeout(time.Second * 2)))
        if err := v.Libvirt.Connect(); err != nil {
                log.Fatalf("failed to connect: %v", err)
        }
	return &v
}


