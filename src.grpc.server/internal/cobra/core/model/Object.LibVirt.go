package object

type VirshDiskDriver struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}


type VirshDiskSource struct {
	Text string `xml:",chardata"`
	File string `xml:"file,attr"`
}

type VirshDiskTarget struct {
	Text string `xml:",chardata"`
	Dev  string `xml:"dev,attr"`
	Bus  string `xml:"bus,attr"`
}


type VirshDisk     struct {
	Text   string `xml:",chardata"`
	Type   string `xml:"type,attr"`
	Device string `xml:"device,attr"`
	Driver VirshDiskDriver `xml:"driver"`
	Source VirshDiskSource `xml:"source"`
	Target VirshDiskTarget `xml:"target"`
}

type VirshNeworkMac  struct {
	Text    string `xml:",chardata"`
	Address string `xml:"address,attr"`
}

type VirshNeworkSource struct {
	Text    string `xml:",chardata"`
	Network string `xml:"network,attr"`
}

type VirshNeworkModel struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type VirshNeworkAddress struct {
	Text          string `xml:",chardata"`
	Type          string `xml:"type,attr"`
	Domain        string `xml:"domain,attr"`
	Bus           string `xml:"bus,attr"`
	Slot          string `xml:"slot,attr"`
	Function      string `xml:"function,attr"`
	Multifunction string `xml:"multifunction,attr"`
}




type  VirshInterface struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
	Mac  VirshNeworkMac `xml:"mac"`
	Source VirshNeworkSource `xml:"source"`
	Model VirshNeworkModel `xml:"model"`
	Address VirshNeworkAddress `xml:"address"`
}


type  VirshMetadata struct {
	Text      string `xml:",chardata"`
	Libosinfo struct {
		Text      string `xml:",chardata"`
		Libosinfo string `xml:"libosinfo,attr"`
		Os        struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"os"`
	} `xml:"libosinfo"`
}

type  VirshMemory struct {
	Text string `xml:",chardata"`
	Unit string `xml:"unit,attr"`
}

type  VirshVcpu struct {
	Text      string `xml:",chardata"`
	Placement string `xml:"placement,attr"`
}

type  VirshOs struct {
	Text string `xml:",chardata"`
	Type struct {
		Text    string `xml:",chardata"`
		Arch    string `xml:"arch,attr"`
		Machine string `xml:"machine,attr"`
	} `xml:"type"`
	Boot struct {
		Text string `xml:",chardata"`
		Dev  string `xml:"dev,attr"`
	} `xml:"boot"`
}


type  VirshFeatures struct {
	Text   string `xml:",chardata"`
	Acpi   string `xml:"acpi"`
	Apic   string `xml:"apic"`
	Vmport struct {
		Text  string `xml:",chardata"`
		State string `xml:"state,attr"`
	} `xml:"vmport"`
}

type  VirshCpu struct {
	Text  string `xml:",chardata"`
	Mode  string `xml:"mode,attr"`
	Check string `xml:"check,attr"`
}

type  VirshClock struct {
	Text   string `xml:",chardata"`
	Offset string `xml:"offset,attr"`
	Timer  []struct {
		Text       string `xml:",chardata"`
		Name       string `xml:"name,attr"`
		Tickpolicy string `xml:"tickpolicy,attr"`
		Present    string `xml:"present,attr"`
	} `xml:"timer"`
}

type  VirshPm    struct {
	Text         string `xml:",chardata"`
	SuspendToMem struct {
		Text    string `xml:",chardata"`
		Enabled string `xml:"enabled,attr"`
	} `xml:"suspend-to-mem"`
	SuspendToDisk struct {
		Text    string `xml:",chardata"`
		Enabled string `xml:"enabled,attr"`
	} `xml:"suspend-to-disk"`
}

type  VirshDevicesController struct {
	Text  string `xml:",chardata"`
	Type  string `xml:"type,attr"`
	Index string `xml:"index,attr"`
	Model string `xml:"model,attr"`
	Ports string `xml:"ports,attr"`
}


type  VirshDevicesSerial struct {
	Text   string `xml:",chardata"`
	Type   string `xml:"type,attr"`
	Target struct {
		Text  string `xml:",chardata"`
		Type  string `xml:"type,attr"`
		Port  string `xml:"port,attr"`
		Model struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"model"`
	} `xml:"target"`
}

type  VirshDevicesConsole struct {
	Text   string `xml:",chardata"`
	Type   string `xml:"type,attr"`
	Target struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Port string `xml:"port,attr"`
	} `xml:"target"`
}


type  VirshDevicesChannel struct {
	Text   string `xml:",chardata"`
	Type   string `xml:"type,attr"`
	Target struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Name string `xml:"name,attr"`
	} `xml:"target"`
}

type  VirshDevicesInput struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
	Bus  string `xml:"bus,attr"`
}


type  VirshDevicesSound struct {
	Text    string `xml:",chardata"`
	Model   string `xml:"model,attr"`
	Address struct {
		Text     string `xml:",chardata"`
		Type     string `xml:"type,attr"`
		Domain   string `xml:"domain,attr"`
		Bus      string `xml:"bus,attr"`
		Slot     string `xml:"slot,attr"`
		Function string `xml:"function,attr"`
	} `xml:"address"`
}

type  VirshDevicesAudio struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
	Type string `xml:"type,attr"`
}

type  VirshDevicesVideo struct {
	Text  string `xml:",chardata"`
	Model struct {
		Text    string `xml:",chardata"`
		Type    string `xml:"type,attr"`
		Ram     string `xml:"ram,attr"`
		Vram    string `xml:"vram,attr"`
		Vgamem  string `xml:"vgamem,attr"`
		Heads   string `xml:"heads,attr"`
		Primary string `xml:"primary,attr"`
	} `xml:"model"`
	Address struct {
		Text     string `xml:",chardata"`
		Type     string `xml:"type,attr"`
		Domain   string `xml:"domain,attr"`
		Bus      string `xml:"bus,attr"`
		Slot     string `xml:"slot,attr"`
		Function string `xml:"function,attr"`
	} `xml:"address"`
}

type  VirshDevicesRedirdev struct {
	Text    string `xml:",chardata"`
	Bus     string `xml:"bus,attr"`
	Type    string `xml:"type,attr"`
	Address struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Bus  string `xml:"bus,attr"`
		Port string `xml:"port,attr"`
	} `xml:"address"`
}

type  VirshDevicesMemballoon struct {
	Text    string `xml:",chardata"`
	Model   string `xml:"model,attr"`
	Address struct {
		Text     string `xml:",chardata"`
		Type     string `xml:"type,attr"`
		Domain   string `xml:"domain,attr"`
		Bus      string `xml:"bus,attr"`
		Slot     string `xml:"slot,attr"`
		Function string `xml:"function,attr"`
	} `xml:"address"`
}

type  VirshDevicesRng struct {
	Text    string `xml:",chardata"`
	Model   string `xml:"model,attr"`
	Backend struct {
		Text  string `xml:",chardata"`
		Model string `xml:"model,attr"`
	} `xml:"backend"`
	Address struct {
		Text     string `xml:",chardata"`
		Type     string `xml:"type,attr"`
		Domain   string `xml:"domain,attr"`
		Bus      string `xml:"bus,attr"`
		Slot     string `xml:"slot,attr"`
		Function string `xml:"function,attr"`
	} `xml:"address"`
}


type  VirshDevicesGraphics struct {
	Text     string `xml:",chardata"`
	Type     string `xml:"type,attr"`
	Autoport string `xml:"autoport,attr"`
	Listen   struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"listen"`
	Image struct {
		Text        string `xml:",chardata"`
		Compression string `xml:"compression,attr"`
	} `xml:"image"`
}

type  VirshDevices struct {
	Text     string `xml:",chardata"`
	Emulator string `xml:"emulator"`
	Disk     []VirshDisk `xml:"disk"`
	Controller []VirshDevicesController `xml:"controller"`
	Interface []VirshInterface `xml:"interface"`
	Serial VirshDevicesSerial `xml:"serial"`
	Console VirshDevicesConsole `xml:"console"`
	Channel VirshDevicesChannel `xml:"channel"`
	Input []VirshDevicesInput `xml:"input"`
	Graphics struct {
		Text     string `xml:",chardata"`
		Type     string `xml:"type,attr"`
		Autoport string `xml:"autoport,attr"`
		Listen   struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"listen"`
		Image struct {
			Text        string `xml:",chardata"`
			Compression string `xml:"compression,attr"`
		} `xml:"image"`
	} `xml:"graphics"`
	Sound VirshDevicesSound `xml:"sound"`
	Audio VirshDevicesAudio `xml:"audio"`
	Video VirshDevicesVideo `xml:"video"`
	Redirdev []VirshDevicesRedirdev `xml:"redirdev"`
	Memballoon VirshDevicesMemballoon `xml:"memballoon"`
	Rng VirshDevicesRng `xml:"rng"`
}

type Domain struct {
	//      XMLName  xml.Name `xml:"domain"`
	Text     string `xml:",chardata"`
	Type     string `xml:"type,attr"`
	Name     string `xml:"name"`
	Metadata VirshMetadata `xml:"metadata"`
	Memory VirshMemory `xml:"memory"`
	Vcpu VirshVcpu `xml:"vcpu"`
	Os VirshOs `xml:"os"`
	Features VirshFeatures `xml:"features"`
	Cpu VirshCpu `xml:"cpu"`
	Clock VirshClock `xml:"clock"`
	OnPoweroff string `xml:"on_poweroff"`
	OnReboot   string `xml:"on_reboot"`
	OnCrash    string `xml:"on_crash"`
	Pm         VirshPm `xml:"pm"`
	Devices VirshDevices `xml:"devices"`
}

