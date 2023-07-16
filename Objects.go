package main

import (

        "github.com/spf13/cobra"
	"github.com/digitalocean/go-libvirt"
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


type CobraMenu struct { }


type VMDisk struct {
        Path     string `yaml:"path"`
        Tmpl string `yaml:"tmpl"`
}

type InterfaceName struct {
        Name     string `yaml:"name"`
        MacAddress string `yaml:"mac"`
	Slot	string  `yaml:"slot"`
}


type VirtualMachineStatus string

type VirtualMachine struct {
        CPUCount uint16
        CPUTime  uint64
        MemoryBytes uint64
        MaxMemoryBytes uint64
        State VirtState
        Libvirt *libvirt.Libvirt
}


type Config struct {
        INTERFACEINIT    []InterfaceName `yaml:"INTERFACE_INIT"`
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
        STEP2            bool     `yaml:"STEP2"`
        STEP3            bool     `yaml:"STEP3"`
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
type LibVirtVM struct {

        VMPath          string
        VMNAME          string  `yaml:"VMNAME"`
        VMNAME_FQDN     string  `yaml:"VMNAME_FQDN"`
        NodeId        	string  `yaml:"nodeid"`
        ROOTFS_SIZE   	int     `yaml:"ROOTFS_SIZE"`
	DISKID		string	`yaml:"DISKID"`
	EXT_DISK_SIZE	int 	`yaml:"EXT_DISK_SIZE"`
	DetachDiskName	string
	XmlTemplate     string

        MEMORY                  int
        CORE                    int
        DISKSDA                 string
        DISKSDBCLOUDINIT        string

        AfterDeploy   []string `yaml:"after-deploy"`
        Command       []string `yaml:"command"`
        Config        Config   `yaml:"config"`
        CreateImageVM []string `yaml:"create-image-vm"`
        Network       Network  `yaml:"network"`
        Pgk   []string `yaml:"pgk"`
        SSHKeys       []string `yaml:"ssh-keys"`
        VMDisk        []VMDisk `yaml:"vm-disk"`
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
	XmlTemplate	string
	myMapFlag1 int
	myMapFlag2 string
}
