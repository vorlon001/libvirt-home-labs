package main

import (
	"github.com/spf13/cobra"
)



// ./cobra LibVirt MachineState --id node190

// ./cobra Configure initVM --CORE 4 --EXT_DISK_SIZE 30 --MEMORY 4096 --Octet 180 --ROOTFS_SIZE 30 --USER_DATA_PATH user-data-jammy.yaml  --VMNAME node180

// ./main Configure attachDiskVM --CORE 4 --EXT_DISK_SIZE 30 --MEMORY 4096 --Octet 180 --ROOTFS_SIZE 30 --USER_DATA_PATH user-data-jammy.yaml  --VMNAME node180
// ./main Configure attachInerfaceVM --CORE 4 --EXT_DISK_SIZE 30 --MEMORY 4096 --Octet 180 --ROOTFS_SIZE 30 --USER_DATA_PATH user-data-jammy.yaml  --VMNAME node180

// ./cobra LibVirt MachineDestroy --id node180

// ./cobra LibVirt MachineCreate --xml  /cloud/KVM/kvm_examples.v20/node180/vm.xml
// ./cobra LibVirt MachineStart --id node180

// ./cobra LibVirt MachineShutoff --id node180
// ./cobra LibVirt MachineDelete --id node180

// https://stackoverflow.com/questions/10200178/call-a-method-from-a-go-template

//***************************************************************************************

// ./cobra.2 Configure sub1 -c 5 -m 4096

func initCoreCobra() {

	core := Singleton[Core]()

	core.Menu = []Menu{
                                        Menu{
                                                Name:"RootCmdLibVirt",
                                                SubMenu: []SubMenu{
                                                    SubMenu{ Name: "RootSubCmdvirtualMachineState",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMid, "id", "", "","Libvirt VMname")
                                                    }},
                                                    SubMenu{ Name: "RootSubCmdvirtualMachineSoftReboot",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMid, "id", "", "","Libvirt VMname")
                                                    }},
                                                    SubMenu{ Name: "RootSubCmdvirtualMachineHardReboot",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMid, "id", "", "","Libvirt VMname")
                                                    }},
                                                    SubMenu{ Name: "RootSubCmdvirtualMachineShutdown",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMid, "id", "", "","Libvirt VMname")
                                                    }},
                                                    SubMenu{ Name: "RootSubCmdvirtualMachineShutoff",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMid, "id", "", "","Libvirt VMname")
                                                    }},
                                                    SubMenu{ Name: "RootSubCmdvirtualMachineStart",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMid, "id", "", "","Libvirt VMname")
                                                    }},
                                                    SubMenu{ Name: "RootSubCmdvirtualMachinePause",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMid, "id", "", "","Libvirt VMname")
                                                    }},
                                                    SubMenu{ Name: "RootSubCmdvirtualMachineResume",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMid, "id", "", "","Libvirt VMname")
                                                    }},
                                                    SubMenu{ Name: "RootSubCmdvirtualMachineCreate",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.XmlTemplate, "xml", "", "","Libvirt Create VM from Xml Template")
                                                    }},
                                                    SubMenu{ Name: "RootSubCmdvirtualMachineDelete",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMid, "id", "", "","Libvirt VMname")
                                                    }},
                                                    SubMenu{ Name: "RootSubCmdvirtualMachineDestroy",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMid, "id", "", "","Libvirt VMname")
                                                    }},
                                                  },
					},
                                        Menu{
                                                Name:"RootCmdConfigure",
                                                SubMenu: []SubMenu{

                                                    SubMenu{ Name: "RootSubCmdinitVM",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMNAME, "VMNAME", "", "","Libvirt VM hostname")
                                                        c.PersistentFlags().IntVarP(&core.CORE, "CORE", "", 0,"Libvirt CORE")
                                                        c.PersistentFlags().IntVarP(&core.MEMORY, "MEMORY", "", 0,"Libvirt MEMORY Mb")
                                                        c.PersistentFlags().IntVarP(&core.ROOTFS_SIZE, "ROOTFS_SIZE", "", 0,"Libvirt ROOTFS SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.Octet, "Octet", "", "","Libvirt network Octet")
                                                        c.PersistentFlags().IntVarP(&core.EXT_DISK_SIZE, "EXT_DISK_SIZE", "", 0,"Libvirt EXT DISK SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.USER_DATA_PATH, "USER_DATA_PATH", "", "","Libvirt USER-DATA PATH")
                                                    }},


                                                    SubMenu{ Name: "RootSubCmdinitVMs",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMNAME, "VMNAME", "", "","Libvirt VM hostname")
                                                        c.PersistentFlags().IntVarP(&core.CORE, "CORE", "", 0,"Libvirt CORE")
                                                        c.PersistentFlags().IntVarP(&core.MEMORY, "MEMORY", "", 0,"Libvirt MEMORY Mb")
                                                        c.PersistentFlags().IntVarP(&core.ROOTFS_SIZE, "ROOTFS_SIZE", "", 0,"Libvirt ROOTFS SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.Octet, "Octet", "", "","Libvirt network Octet")
                                                        c.PersistentFlags().IntVarP(&core.EXT_DISK_SIZE, "EXT_DISK_SIZE", "", 0,"Libvirt EXT DISK SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.USER_DATA_PATH, "USER_DATA_PATH", "", "","Libvirt USER-DATA PATH")
							c.PersistentFlags().IntVarP(&core.NumVM, "NumVM", "", 0,"Libvirt NumVM")
                                                    }},


                                                    SubMenu{ Name: "RootSubCmdattachDiskVM",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMNAME, "VMNAME", "", "","Libvirt VM hostname")
                                                        c.PersistentFlags().IntVarP(&core.CORE, "CORE", "", 0,"Libvirt CORE")
                                                        c.PersistentFlags().IntVarP(&core.MEMORY, "MEMORY", "", 0,"Libvirt MEMORY Mb")
                                                        c.PersistentFlags().IntVarP(&core.ROOTFS_SIZE, "ROOTFS_SIZE", "", 0,"Libvirt ROOTFS SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.Octet, "Octet", "", "","Libvirt network Octet")
                                                        c.PersistentFlags().IntVarP(&core.EXT_DISK_SIZE, "EXT_DISK_SIZE", "", 0,"Libvirt EXT DISK SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.USER_DATA_PATH, "USER_DATA_PATH", "", "","Libvirt USER-DATA PATH")
                                                    }},

                                                    SubMenu{ Name: "RootSubCmdDetachDiskVM",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMNAME, "VMNAME", "", "","Libvirt VM hostname")
                                                        c.PersistentFlags().IntVarP(&core.CORE, "CORE", "", 0,"Libvirt CORE")
                                                        c.PersistentFlags().IntVarP(&core.MEMORY, "MEMORY", "", 0,"Libvirt MEMORY Mb")
                                                        c.PersistentFlags().IntVarP(&core.ROOTFS_SIZE, "ROOTFS_SIZE", "", 0,"Libvirt ROOTFS SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.Octet, "Octet", "", "","Libvirt network Octet")
                                                        c.PersistentFlags().IntVarP(&core.EXT_DISK_SIZE, "EXT_DISK_SIZE", "", 0,"Libvirt EXT DISK SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.USER_DATA_PATH, "USER_DATA_PATH", "", "","Libvirt USER-DATA PATH")
                                                        c.PersistentFlags().StringVarP(&core.DetachDiskName, "DetachDisk", "", "","Libvirt Detach Disk <sda,sdb,sdc,...>")
                                                    }},

                                                    SubMenu{ Name: "RootSubCmdattachInerfaceVM",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMNAME, "VMNAME", "", "","Libvirt VM hostname")
                                                        c.PersistentFlags().IntVarP(&core.CORE, "CORE", "", 0,"Libvirt CORE")
                                                        c.PersistentFlags().IntVarP(&core.MEMORY, "MEMORY", "", 0,"Libvirt MEMORY Mb")
                                                        c.PersistentFlags().IntVarP(&core.ROOTFS_SIZE, "ROOTFS_SIZE", "", 0,"Libvirt ROOTFS SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.Octet, "Octet", "", "","Libvirt network Octet")
                                                        c.PersistentFlags().IntVarP(&core.EXT_DISK_SIZE, "EXT_DISK_SIZE", "", 0,"Libvirt EXT DISK SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.USER_DATA_PATH, "USER_DATA_PATH", "", "","Libvirt USER-DATA PATH")
                                                    }},


                                                    SubMenu{ Name: "RootSubCmdaDetachInerfaceVM",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMNAME, "VMNAME", "", "","Libvirt VM hostname")
                                                        c.PersistentFlags().IntVarP(&core.CORE, "CORE", "", 0,"Libvirt CORE")
                                                        c.PersistentFlags().IntVarP(&core.MEMORY, "MEMORY", "", 0,"Libvirt MEMORY Mb")
                                                        c.PersistentFlags().IntVarP(&core.ROOTFS_SIZE, "ROOTFS_SIZE", "", 0,"Libvirt ROOTFS SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.Octet, "Octet", "", "","Libvirt network Octet")
                                                        c.PersistentFlags().IntVarP(&core.EXT_DISK_SIZE, "EXT_DISK_SIZE", "", 0,"Libvirt EXT DISK SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.USER_DATA_PATH, "USER_DATA_PATH", "", "","Libvirt USER-DATA PATH")
                                                        c.PersistentFlags().StringVarP(&core.NetworkSlot, "NetworkSlot", "", "","Libvirt NetworkSlot")
                                                    }},


                                                    SubMenu{ Name: "RootSubCmddestroyVM",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMNAME, "VMNAME", "", "","Libvirt VM hostname")
                                                        c.PersistentFlags().IntVarP(&core.CORE, "CORE", "", 0,"Libvirt CORE")
                                                        c.PersistentFlags().IntVarP(&core.MEMORY, "MEMORY", "", 0,"Libvirt MEMORY Mb")
                                                        c.PersistentFlags().IntVarP(&core.ROOTFS_SIZE, "ROOTFS_SIZE", "", 0,"Libvirt ROOTFS SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.Octet, "Octet", "", "","Libvirt network Octet")
                                                        c.PersistentFlags().IntVarP(&core.EXT_DISK_SIZE, "EXT_DISK_SIZE", "", 0,"Libvirt EXT DISK SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.USER_DATA_PATH, "USER_DATA_PATH", "", "","Libvirt USER-DATA PATH")
                                                    }},


                                                    SubMenu{ Name: "RootSubCmdlistInterfaceVM",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMNAME, "VMNAME", "", "","Libvirt VM hostname")
                                                        c.PersistentFlags().IntVarP(&core.CORE, "CORE", "", 0,"Libvirt CORE")
                                                        c.PersistentFlags().IntVarP(&core.MEMORY, "MEMORY", "", 0,"Libvirt MEMORY Mb")
                                                        c.PersistentFlags().IntVarP(&core.ROOTFS_SIZE, "ROOTFS_SIZE", "", 0,"Libvirt ROOTFS SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.Octet, "Octet", "", "","Libvirt network Octet")
                                                        c.PersistentFlags().IntVarP(&core.EXT_DISK_SIZE, "EXT_DISK_SIZE", "", 0,"Libvirt EXT DISK SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.USER_DATA_PATH, "USER_DATA_PATH", "", "","Libvirt USER-DATA PATH")
                                                    }},


                                                    SubMenu{ Name: "RootSubCmdlistDiskVM",
                                                    Run: func(c *cobra.Command) {
                                                        core := Singleton[Core]()
                                                        c.PersistentFlags().StringVarP(&core.VMNAME, "VMNAME", "", "","Libvirt VM hostname")
                                                        c.PersistentFlags().IntVarP(&core.CORE, "CORE", "", 0,"Libvirt CORE")
                                                        c.PersistentFlags().IntVarP(&core.MEMORY, "MEMORY", "", 0,"Libvirt MEMORY Mb")
                                                        c.PersistentFlags().IntVarP(&core.ROOTFS_SIZE, "ROOTFS_SIZE", "", 0,"Libvirt ROOTFS SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.Octet, "Octet", "", "","Libvirt network Octet")
                                                        c.PersistentFlags().IntVarP(&core.EXT_DISK_SIZE, "EXT_DISK_SIZE", "", 0,"Libvirt EXT DISK SIZE Gb")
                                                        c.PersistentFlags().StringVarP(&core.USER_DATA_PATH, "USER_DATA_PATH", "", "","Libvirt USER-DATA PATH")
                                                    }},


                                                },
					},
		}
}


func initMenu() *cobra.Command {

	initCoreCobra()

        var rootCmd = &cobra.Command{Use: "HomeLabs"}
        core := Singleton[Core]()

	for _,v := range core.Menu {
		rootMenu := GetCobraMenu(v.Name)
		for _,j := range v.SubMenu {
			subCmd := GetCobraMenu(j.Name)
			j.Run(subCmd)
			rootMenu.AddCommand(subCmd)
		}
		rootCmd.AddCommand(rootMenu)
	}
	return rootCmd
}
