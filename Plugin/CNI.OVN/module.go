package cni_OVN

import (
	"fmt"

	cni "play.ground/CNI"
)

var _ cni.Cni = (*cni_ovn)(nil)
var _ = (cni.Cni)((*cni_ovn)(nil))

var Plugin_Name string

type cni_ovn struct {
	Data_cni_ovn string
}

func (a *cni_ovn) Print(z string) {
	fmt.Printf("Plugin.print: %#v %#v\n", z, a.Data_cni_ovn)
}
func newPlugin(z string) cni.Cni {
	return &cni_ovn{Data_cni_ovn: z}
}
func init() {
	Plugin_Name = "cni_cni_ovn"
	cni.RegisterStruct(Plugin_Name, newPlugin("cni_cni_ovn struct"))
	fmt.Printf("init.Plugin: %v\n", Plugin_Name)
}
