package virtualmachine

import (

	"fmt"
	"time"
        "strconv"
        "io/ioutil"
        "encoding/xml"

        "github.com/digitalocean/go-libvirt"
        "github.com/digitalocean/go-libvirt/socket/dialers"
        "gitlab.iblog.pro/cobra/libvirt/internal/cobra/store"
        "github.com/sirupsen/logrus"
        logs "gitlab.iblog.pro/cobra/libvirt/internal/cobra/logs"
	PanicRecover "gitlab.iblog.pro/cobra/libvirt/internal/cobra/core/panicrecover"
        Model "gitlab.iblog.pro/cobra/libvirt/internal/cobra/core/model"
)


type VirtualMachineInterface interface {
        VirtualMachineState(id string) (string, error)
        VirtualMachineCreate(xmlTemplate string) (string, error)
        VirtualMachineDelete(id string) (string, error)
        VirtualMachineSoftReboot(id string) (string, error)
        VirtualMachineHardReboot(id string) (string, error)
        VirtualMachineShutdown(id string) (string, error)
        VirtualMachineShutoff(id string) (string, error)
        VirtualMachineStart(id string) (string, error)
        VirtualMachinePause(id string) (string, error)
        VirtualMachineResume(id string) (string, error)
        VirtualMachineMachineState() (string, error)
        VirtualMachineMigrate(string, string) (string, error)
        GetLibvirt() *libvirt.Libvirt
}


type VirtualMachine struct {
        CPUCount uint16
        CPUTime  uint64
        MemoryBytes uint64
        MaxMemoryBytes uint64
        State Model.VirtState
        Libvirt *libvirt.Libvirt
}


func (ret *VirtualMachine) GetLibvirt() *libvirt.Libvirt {
        return ret.Libvirt
}


func (ret *VirtualMachine) VirtualMachineMachineState() (string, error) {
        core := store.Singleton[Model.Core]()
	logs.Log.WithFields(logrus.Fields{ "core": core,}).Info("Inside RootSubCmdvirtualMachineState Run with args")
	virt,err := Virtinit()
        if err != nil {
                logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                return "Error", err
        }
        _, err = virt.VirtualMachineState(core.VMid)
        if err != nil {
            logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%s",err),}).Info("Error init virt.VirtualMachineState()")
                return "Error", err
        }


        doms, err := virt.GetLibvirt().Domains()
        if err != nil {
                logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                return "Error", err
        }

        for _,j := range doms {
                if j.Name == core.VMid {

                        domainGetXMLDesc, err := virt.GetLibvirt().DomainGetXMLDesc(j, libvirt.DomainXMLSecure)
                        if err != nil {
                                logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("Cobra Event Error")
                                return "Error", err
                        }

                        t := Model.Domain{}
                        xml.Unmarshal([]byte(domainGetXMLDesc), &t)
			logs.Log.Info("--------------------------------------------------")
                        for k,v := range t.Devices.Disk {
				logs.Log.WithFields(logrus.Fields{	"key": k, "v.Type": v.Type, "v.Device": v.Device, "v.Source.File": v.Source.File,
								"v.Target.Dev": v.Target.Dev, "v.Target.Bus": v.Target.Bus, }).Info("RootSubCmdvirtualMachineState")
				logs.Log.Info("=========================================================")
                        }
			logs.Log.Info("--------------------------------------------------")
                        for _, v := range t.Devices.Interface {
                                logs.Log.WithFields(logrus.Fields{ "v.Type": v.Type,}).Info("VM interface id")
                                logs.Log.WithFields(logrus.Fields{ "v.Mac.Address": v.Mac.Address,}).Info("VM Interface.Type")
                                logs.Log.WithFields(logrus.Fields{ "v.Source.Network": v.Source.Network,}).Info("VM Interface.Mac.Address")
                                logs.Log.WithFields(logrus.Fields{ "v.Model.Type": v.Model.Type,}).Info("VM Interface.Source.Network")
                                logs.Log.WithFields(logrus.Fields{ "v.Address.Type": v.Address.Type,}).Info("VM Interface.Model.Type")
                                logs.Log.WithFields(logrus.Fields{ "v.Address.Type": v.Address.Type,}).Info("VM Interface.Address.Type")
                                logs.Log.WithFields(logrus.Fields{ "v.Address.Domain": v.Address.Domain,}).Info("VM Interface.Address.Domain")
                                logs.Log.WithFields(logrus.Fields{ "v.Address.Bus": v.Address.Bus,}).Info("VM Interface.Address.Bus")
                                logs.Log.WithFields(logrus.Fields{ "v.Address.Slot": v.Address.Slot,}).Info("VM Interface.Address.Slot")
                                logs.Log.WithFields(logrus.Fields{ "v.Address.Function": v.Address.Function,}).Info("VM Interface.Address.Function")
                                logs.Log.WithFields(logrus.Fields{ "v.Address.Multifunction": v.Address.Multifunction,}).Info("VM Interface.Address.Multifunction")
				logs.Log.Info("=========================================================")
                        }
			logs.Log.Info("--------------------------------------------------")
                }
        }
        return "Done", nil
}

func (ret *VirtualMachine) VirtualMachineMigrate(VMid string, ToMove string) (string, error) {

        core := store.Singleton[Model.Core]()
        logs.Log.WithFields(logrus.Fields{ "core": core,}).Info("Inside RootSubCmdvirtualMachineStart Run with args")

        domains, err := ret.Libvirt.Domains()
        if err != nil {
                logs.Log.Fatalf("failed to retrieve domains: %v", err)
                return "Error", err
        }

        fmt.Println("ID\tName\t\tUUID\t\t\t\tStatus")
        fmt.Printf("---------------------------------------------------------------------------------------------------------------------------------\n")
        for _, d := range domains {
                a, _ := ret.Libvirt.DomainState(d.Name)
                c, _ := strconv.Atoi(fmt.Sprintf("%d",a))
                fmt.Printf("%#v\n",d);
                fmt.Printf("%d\t%s\t%#s\n", d.ID, d.Name, Model.DomainState[c])
        }

        flags := libvirt.MigrateLive |
		            libvirt.MigratePeer2peer |
		            libvirt.MigratePersistDest |
		            libvirt.MigrateChangeProtection |
		            libvirt.MigrateAbortOnError |
		            libvirt.MigrateAutoConverge |
		            libvirt.MigrateNonSharedDisk

        dom, err := ret.Libvirt.DomainLookupByName(VMid)
        if err != nil {
                logs.Log.Fatalf("failed to lookup domain: %v", err)
                return "Error", err
        }
        dconnuri := []string{ fmt.Sprintf("qemu+ssh://root@%s/system",ToMove)}
        if e, err := ret.Libvirt.DomainMigratePerform3Params( dom, dconnuri,
                                                              []libvirt.TypedParam{}, []byte{}, flags); err != nil {
		    logs.Log.Fatalf("unexpected live migration error: %v", err)
        } else {
            fmt.Printf("%#v\n",e);
        }

        domains, err = ret.Libvirt.Domains()
        if err != nil {
                logs.Log.Fatalf("failed to retrieve domains: %v", err)
                return "Error", err
        }

        fmt.Println("ID\tName\t\tUUID\t\t\t\tStatus")
        fmt.Printf("---------------------------------------------------------------------------------------------------------------------------------\n")
        for _, d := range domains {
                a, _ := ret.Libvirt.DomainState(d.Name)
                c, _ := strconv.Atoi(fmt.Sprintf("%d",a))
                fmt.Printf("%#v\n",d);
                fmt.Printf("%d\t%s\t%x\t%#s\n", d.ID, d.Name, d.UUID, Model.DomainState[c])
        }
        return "Done", nil
}


// VirtualMachineState returns current state of a virtual machine.
func (ret *VirtualMachine) VirtualMachineState(id string) (string, error) {
//        var ret VirtualMachine //VirtualMachineState

        d, err := ret.Libvirt.DomainLookupByName(id)
	logs.Log.WithFields(logrus.Fields{ "DomainLookupByName": d, "err": err, "id": id, }).Info("Inside VirtualMachineState Run")
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        state, maxmem, mem, ncpu, cputime, err := ret.Libvirt.DomainGetInfo(d)
	logs.Log.WithFields(logrus.Fields{ "state": state, "maxmem": maxmem, "mem": mem, "ncpu": ncpu, "cputime": cputime, "err": err, }).Info("Inside VirtualMachineState Run")
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        ret.CPUCount = ncpu
        ret.CPUTime = cputime
        // God only knows why they return memory in kilobytes.
        ret.MemoryBytes = mem * 1024
        ret.MaxMemoryBytes = maxmem * 1024
        temp := libvirt.DomainState(state)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        switch temp {
        case libvirt.DomainNostate:
                ret.State = Model.VirtStatePending
        case libvirt.DomainRunning:
                ret.State = Model.VirtStateRunning
        case libvirt.DomainBlocked:
                ret.State = Model.VirtStateBlocked
        case libvirt.DomainPaused:
                ret.State = Model.VirtStatePaused
        case libvirt.DomainShutdown:
                ret.State = Model.VirtStateShutdown
        case libvirt.DomainShutoff:
                ret.State = Model.VirtStateShutoff
        case libvirt.DomainCrashed:
                ret.State = Model.VirtStateCrashed
        case libvirt.DomainPmsuspended:
                ret.State = Model.VirtStateHybernating
        }
        ret.Libvirt.DomainGetState(d, 0)
        PanicRecover.Hret(ret)
        return "Done", nil
}



// VirtualMachineCreate creates a new VM from an xml template file
func (ret *VirtualMachine) VirtualMachineCreate(xmlTemplate string) (string, error) {

        xml, err := ioutil.ReadFile(xmlTemplate)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        d, err := ret.Libvirt.DomainDefineXML(string(xml))
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        PanicRecover.Hret(d)
        return "Done", nil
}

// VirtualMachineDelete deletes a new VM from an xml template file
func (ret *VirtualMachine) VirtualMachineDelete(id string) (string, error) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        err = ret.Libvirt.DomainUndefineFlags(d, libvirt.DomainUndefineKeepNvram)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        PanicRecover.Hok(fmt.Sprintf("%v was deleted", id))
        return "Done", nil
}

// VirtualMachineSoftReboot reboots a machine gracefully, as chosen by hypervisor.
func (ret *VirtualMachine) VirtualMachineSoftReboot(id string) (string, error) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        err = ret.Libvirt.DomainReboot(d, libvirt.DomainRebootDefault)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        PanicRecover.Hok(fmt.Sprintf("%v was soft-rebooted successfully", id))
        return "Done", nil
}

// VirtualMachineHardReboot sends a VM into hard-reset mode. This is damaging to all ongoing file operations.
func (ret *VirtualMachine) VirtualMachineHardReboot(id string) (string, error) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        err = ret.Libvirt.DomainReset(d, 0)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        PanicRecover.Hok(fmt.Sprintf("%v was hard-rebooted successfully", id))
        return "Done", nil
}


// VirtualMachineShutdown gracefully shuts down the VM.
func (ret *VirtualMachine) VirtualMachineShutdown(id string) (string, error) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        err = ret.Libvirt.DomainShutdown(d)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        PanicRecover.Hok(fmt.Sprintf("%v was shutdown successfully", id))
        return "Done", nil
}

// VirtualMachineShutoff kills running VM. Equivalent to pulling a plug out of a computer.
func (ret *VirtualMachine) VirtualMachineShutoff(id string) (string, error) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        err = ret.Libvirt.DomainDestroy(d)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        PanicRecover.Hok(fmt.Sprintf("%v was shutoff successfully", id))
        return "Done", nil
}

// VirtualMachineStart starts up a VM.
func (ret *VirtualMachine) VirtualMachineStart(id string) (string, error) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        //v.DomainRestore()
        //_, err = ret.Libvirt.DomainCreateWithFlags(d, uint32(libvirt.DomainStartBypassCache))
        err = ret.Libvirt.DomainCreate(d)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        PanicRecover.Hok(fmt.Sprintf("%v was started", id))
        return "Done", nil
}

// VirtualMachinePause stops the execution of the VM. CPU is not used, but memory is still occupied.
func (ret *VirtualMachine) VirtualMachinePause(id string) (string, error) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        err = ret.Libvirt.DomainSuspend(d)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        PanicRecover.Hok(fmt.Sprintf("%v is paused", id))
        return "Done", nil
}

// VirtualMachineResume can be called after Pause, to resume the invocation of the VM.
func (ret *VirtualMachine) VirtualMachineResume(id string) (string, error) {
        d, err := ret.Libvirt.DomainLookupByName(id)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        err = ret.Libvirt.DomainResume(d)
        if err != nil {
                PanicRecover.Herr(err)
                return "Error", err
        }

        PanicRecover.Hok(fmt.Sprintf("%v was resumed", id))
        return "Done", nil
}

/*
func Virtinit() *VirtualMachine {
	v := VirtualMachine{}
        v.Libvirt = libvirt.NewWithDialer(dialers.NewLocal(dialers.WithLocalTimeout(time.Second * 2)))
        if err := v.Libvirt.Connect(); err != nil {
                logs.Log.Fatalf("failed to connect: %v", err)
        }
	return &v
}
*/



func Virtinit() (VirtualMachineInterface, error) {
        v := VirtualMachine{}
        v.Libvirt = libvirt.NewWithDialer(dialers.NewLocal(dialers.WithLocalTimeout(time.Second * 2)))
        if err := v.Libvirt.Connect(); err != nil {
                logs.Log.Fatalf("failed to connect: %v", err)
                return nil, err
        }
        return &v, nil
}


