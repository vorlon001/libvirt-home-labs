package main

import (
        "sync"
        "github.com/sirupsen/logrus"
)

var log *logrus.Logger




var cache sync.Map

var Version string

//**************************************************************

// VirState represents current lifecycle state of a machine
// Pending = VM was just created and there is no state yet
// Running = VM is running
// Blocked = Blocked on resource
// Paused = VM is paused
// Shutdown = VM is being shut down
// Shutoff = VM is shut off
// Crashed = Most likely VM crashed on startup cause something is missing.
// Hybernating = Virtual Machine is hybernating usually due to guest machine request
// TODO:

const (
        VirtStatePending     = VirtState("Pending")     // VM was just created and there is no state yet
        VirtStateRunning     = VirtState("Running")     // VM is running
        VirtStateBlocked     = VirtState("Blocked")     // VM Blocked on resource
        VirtStatePaused      = VirtState("Paused")      // VM is paused
        VirtStateShutdown    = VirtState("Shutdown")    // VM is being shut down
        VirtStateShutoff     = VirtState("Shutoff")     // VM is shut off
        VirtStateCrashed     = VirtState("Crashed")     // Most likely VM crashed on startup cause something is missing.
        VirtStateHybernating = VirtState("Hybernating") // VM is hybernating usually due to guest machine request
)



const (
        VirtualMachineStatusDeleted      = VirtualMachineStatus("deleted")
        VirtualMachineStatusCreated      = VirtualMachineStatus("created")
        VirtualMachineStatusReady        = VirtualMachineStatus("ready")
        VirtualMachineStatusStarting     = VirtualMachineStatus("starting")
        VirtualMachineStatusImaging      = VirtualMachineStatus("imaging")
        VirtualMachineStatusRunning      = VirtualMachineStatus("running")
        VirtualMachineStatusOff          = VirtualMachineStatus("off")
        VirtualMachineStatusShuttingDown = VirtualMachineStatus("shutting_down")


	CDDiskTemplate = `
<disk type="file" device="disk">
        <driver name="qemu" type="raw"></driver>
        <source file="%s"></source>
        <target dev="vda" bus="virtio"></target>
</disk>`

	SCSIDiskTemplate = `<disk type="file" device="disk">
        <driver name='qemu' type='qcow2'/>
        <source file="{{.Config.VMPATH}}/{{.VMNAME}}/{{.VMNAME}}-disk{{.DISKID}}.qcow2"></source>
        <target dev='{{.DISKID}}' bus='scsi'/>
</disk>`

	E1000Networkemplate = `<interface type='network'>
      <source network='cloud_sw1'/>
      <model type='e1000'/>
      <mac address='{{.MacAddress}}'/>
      <address type='pci' domain='0x0000' bus='0x03' slot='{{.Slot}}' function='0x0' multifunction='on'/>
</interface>`






)

