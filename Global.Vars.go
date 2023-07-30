package main

import (
        Model "iblog.pro/cobra/core/model"
)

var DomainState = map[int]string{
					0: "DomainNostate",
					1: "DomainRunning",
					2: "DomainBlocked",
					3: "DomainPaused",
					4: "DomainShutdown",
					5: "DomainShutoff",
					6: "DomainCrashed",
					7: "DomainPmsuspended" }



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
        VirtStatePending     = Model.VirtState("Pending")     // VM was just created and there is no state yet
        VirtStateRunning     = Model.VirtState("Running")     // VM is running
        VirtStateBlocked     = Model.VirtState("Blocked")     // VM Blocked on resource
        VirtStatePaused      = Model.VirtState("Paused")      // VM is paused
        VirtStateShutdown    = Model.VirtState("Shutdown")    // VM is being shut down
        VirtStateShutoff     = Model.VirtState("Shutoff")     // VM is shut off
        VirtStateCrashed     = Model.VirtState("Crashed")     // Most likely VM crashed on startup cause something is missing.
        VirtStateHybernating = Model.VirtState("Hybernating") // VM is hybernating usually due to guest machine request
)


const (
        VirtualMachineStatusDeleted      = Model.VirtualMachineStatus("deleted")
        VirtualMachineStatusCreated      = Model.VirtualMachineStatus("created")
        VirtualMachineStatusReady        = Model.VirtualMachineStatus("ready")
        VirtualMachineStatusStarting     = Model.VirtualMachineStatus("starting")
        VirtualMachineStatusImaging      = Model.VirtualMachineStatus("imaging")
        VirtualMachineStatusRunning      = Model.VirtualMachineStatus("running")
        VirtualMachineStatusOff          = Model.VirtualMachineStatus("off")
        VirtualMachineStatusShuttingDown = Model.VirtualMachineStatus("shutting_down")

/*
	CDDiskTemplate = `
<disk type="file" device="disk">
        <driver name="qemu" type="raw"></driver>
        <source file="%s"></source>
        <target dev="vda" bus="virtio"></target>
</disk>`
*/

/*	SCSIDiskTemplate = `<disk type="file" device="disk">
        <driver name='qemu' type='qcow2'/>
        <source file="{{.Config.VMPATH}}/{{.VMNAME}}/{{.VMNAME}}-disk{{.DISKID}}.qcow2"></source>
        <target dev='{{.DISKID}}' bus='scsi'/>
</disk>`
*/
/*
	E1000Networkemplate = `<interface type='network'>
      <source network='cloud_sw1'/>
      <model type='e1000'/>
      <mac address='{{.MacAddress}}'/>
      <address type='pci' domain='0x0000' bus='0x03' slot='{{.Slot}}' function='0x0' multifunction='on'/>
</interface>`
*/





)

