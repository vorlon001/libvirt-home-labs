package main

type Domain struct {
	//      XMLName  xml.Name `xml:"domain"`
	Text     string `xml:",chardata"`
	Type     string `xml:"type,attr"`
	Name     string `xml:"name"`
	Metadata struct {
		Text      string `xml:",chardata"`
		Libosinfo struct {
			Text      string `xml:",chardata"`
			Libosinfo string `xml:"libosinfo,attr"`
			Os        struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"os"`
		} `xml:"libosinfo"`
	} `xml:"metadata"`
	Memory struct {
		Text string `xml:",chardata"`
		Unit string `xml:"unit,attr"`
	} `xml:"memory"`
	Vcpu struct {
		Text      string `xml:",chardata"`
		Placement string `xml:"placement,attr"`
	} `xml:"vcpu"`
	Os struct {
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
	} `xml:"os"`
	Features struct {
		Text   string `xml:",chardata"`
		Acpi   string `xml:"acpi"`
		Apic   string `xml:"apic"`
		Vmport struct {
			Text  string `xml:",chardata"`
			State string `xml:"state,attr"`
		} `xml:"vmport"`
	} `xml:"features"`
	Cpu struct {
		Text  string `xml:",chardata"`
		Mode  string `xml:"mode,attr"`
		Check string `xml:"check,attr"`
	} `xml:"cpu"`
	Clock struct {
		Text   string `xml:",chardata"`
		Offset string `xml:"offset,attr"`
		Timer  []struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"name,attr"`
			Tickpolicy string `xml:"tickpolicy,attr"`
			Present    string `xml:"present,attr"`
		} `xml:"timer"`
	} `xml:"clock"`
	OnPoweroff string `xml:"on_poweroff"`
	OnReboot   string `xml:"on_reboot"`
	OnCrash    string `xml:"on_crash"`
	Pm         struct {
		Text         string `xml:",chardata"`
		SuspendToMem struct {
			Text    string `xml:",chardata"`
			Enabled string `xml:"enabled,attr"`
		} `xml:"suspend-to-mem"`
		SuspendToDisk struct {
			Text    string `xml:",chardata"`
			Enabled string `xml:"enabled,attr"`
		} `xml:"suspend-to-disk"`
	} `xml:"pm"`
	Devices struct {
		Text     string `xml:",chardata"`
		Emulator string `xml:"emulator"`
		Disk     []struct {
			Text   string `xml:",chardata"`
			Type   string `xml:"type,attr"`
			Device string `xml:"device,attr"`
			Driver struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
				Type string `xml:"type,attr"`
			} `xml:"driver"`
			Source struct {
				Text string `xml:",chardata"`
				File string `xml:"file,attr"`
			} `xml:"source"`
			Target struct {
				Text string `xml:",chardata"`
				Dev  string `xml:"dev,attr"`
				Bus  string `xml:"bus,attr"`
			} `xml:"target"`
		} `xml:"disk"`
		Controller []struct {
			Text  string `xml:",chardata"`
			Type  string `xml:"type,attr"`
			Index string `xml:"index,attr"`
			Model string `xml:"model,attr"`
			Ports string `xml:"ports,attr"`
		} `xml:"controller"`
		Interface []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Mac  struct {
				Text    string `xml:",chardata"`
				Address string `xml:"address,attr"`
			} `xml:"mac"`
			Source struct {
				Text    string `xml:",chardata"`
				Network string `xml:"network,attr"`
			} `xml:"source"`
			Model struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"model"`
			Address struct {
				Text          string `xml:",chardata"`
				Type          string `xml:"type,attr"`
				Domain        string `xml:"domain,attr"`
				Bus           string `xml:"bus,attr"`
				Slot          string `xml:"slot,attr"`
				Function      string `xml:"function,attr"`
				Multifunction string `xml:"multifunction,attr"`
			} `xml:"address"`
		} `xml:"interface"`
		Serial struct {
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
		} `xml:"serial"`
		Console struct {
			Text   string `xml:",chardata"`
			Type   string `xml:"type,attr"`
			Target struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Port string `xml:"port,attr"`
			} `xml:"target"`
		} `xml:"console"`
		Channel struct {
			Text   string `xml:",chardata"`
			Type   string `xml:"type,attr"`
			Target struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Name string `xml:"name,attr"`
			} `xml:"target"`
		} `xml:"channel"`
		Input []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Bus  string `xml:"bus,attr"`
		} `xml:"input"`
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
		Sound struct {
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
		} `xml:"sound"`
		Audio struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Type string `xml:"type,attr"`
		} `xml:"audio"`
		Video struct {
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
		} `xml:"video"`
		Redirdev []struct {
			Text    string `xml:",chardata"`
			Bus     string `xml:"bus,attr"`
			Type    string `xml:"type,attr"`
			Address struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Bus  string `xml:"bus,attr"`
				Port string `xml:"port,attr"`
			} `xml:"address"`
		} `xml:"redirdev"`
		Memballoon struct {
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
		} `xml:"memballoon"`
		Rng struct {
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
		} `xml:"rng"`
	} `xml:"devices"`
}

