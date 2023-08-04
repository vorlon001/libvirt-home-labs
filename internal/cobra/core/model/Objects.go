package object

import (

        "github.com/spf13/cobra"
        InterfaceNetwork "gitlab.iblog.pro/cobra/libvirt/internal/cobra/core/interfacenetwork"
)


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
type VirtState string

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

var DomainState = map[int]string{
                                        0: "DomainNostate",
                                        1: "DomainRunning",
                                        2: "DomainBlocked",
                                        3: "DomainPaused",
                                        4: "DomainShutdown",
                                        5: "DomainShutoff",
                                        6: "DomainCrashed",
                                        7: "DomainPmsuspended" }


type VirtualMachineStatus string



type Config struct {
	IMAGENAME	 string   `yaml:"IMAGENAME"`
        INTERFACEINIT    []InterfaceNetwork.InterfaceName `yaml:"INTERFACE_INIT"`
        NETWORKCONFIGTPL string   `yaml:"NETWORK_CONFIG_TPL"`
        NEXUSREPO        string   `yaml:"NEXUS_REPO"`
        NEXUSREPOSEC     string   `yaml:"NEXUS_REPO_SEC"`
        PIPTPLPATH       string   `yaml:"PIP_TPL_PATH"`
        ROOTCERTPATH     string   `yaml:"ROOT_CERT_PATH"`
        SSDTPLPATH       string   `yaml:"SSD_TPL_PATH"`
        SSHAUTHKEYS      string   `yaml:"SSH_AUTH_KEYS"`
        USERDATATPLPATH  string   `yaml:"USER_DATA_TPL_PATH"`
        VMNAMEFQDN       string   `yaml:"VMNAME_FQDN"`
        VMREPO           string   `yaml:"VM_REPO"`
        VMTEMPLATE       string   `yaml:"VM_TEMPLATE"`
        VMPATH           string   `yaml:"VMPATH"`
        VMIMAGEPATH      string   `yaml:"VMIMAGEPATH"`
	CDDiskTemplate        string   `yaml:"CDDiskTemplate"`
	SCSIDiskTemplate      string   `yaml:"SCSIDiskTemplate"`
	E1000Networkemplate   string   `yaml:"E1000Networkemplate"`


}

type Nameservers struct {
        IP     string `yaml:"ip"`
        Search string `yaml:"search"`
}
type NetworkPort struct {
        Gateway4    string      `yaml:"gateway4"`
        Nameservers Nameservers `yaml:"nameservers"`
        Netmask     int         `yaml:"netmask"`
        Network     string      `yaml:"network"`
        Netmask6    int         `yaml:"netmask6"`
        Network6    string      `yaml:"network6"`
}


type Network struct {
        Block map[string]NetworkPort `yaml:"block"`
        MagicMac string `yaml:"magic_mac"`
}

type SubMenu struct {
    Name string
    Run func(c *cobra.Command)
}

type Menu struct {
	Name    string   `yaml:"name"`
	SubMenu []SubMenu `yaml:"subMenu"`
}
type Core struct {
	Menu []Menu `yaml:"menu"`

	VMNAME		string
	CORE		int
	MEMORY		int
	ROOTFS_SIZE	int
	Octet		string
	EXT_DISK_SIZE	int
	USER_DATA_PATH	string
	DetachDiskName	string
	NetworkSlot	string
	NumVM		int

	VMid		string
	ToMove		string
	XmlTemplate	string
	myMapFlag1 int
	myMapFlag2 string
}
