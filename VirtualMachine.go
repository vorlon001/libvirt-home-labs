package main

import (

	"fmt"
	"time"

        "io/ioutil"

        "github.com/digitalocean/go-libvirt"
        "github.com/digitalocean/go-libvirt/socket/dialers"

        "github.com/sirupsen/logrus"
)


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


