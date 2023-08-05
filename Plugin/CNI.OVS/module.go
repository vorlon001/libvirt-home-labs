package cni_OVS

import (
	"fmt"

	cni "play.ground/CNI"
)

var _ cni.Cni = (*cni_Ovs)(nil)
var _ = (cni.Cni)((*cni_Ovs)(nil))

var Plugin_Name string

type cni_Ovs struct {
	Data_cni_Ovs string
}

func (a *cni_Ovs) Print(z string) {
	fmt.Printf("Plugin.print: %#v %#v\n", z, a.Data_cni_Ovs)
}
func newPlugin(z string) cni.Cni {
	return &cni_Ovs{Data_cni_Ovs: z}
}
func init() {
	Plugin_Name = "cni_cni_Ovs"
	cni.RegisterStruct(Plugin_Name, newPlugin("cni_cni_Ovs struct"))
	fmt.Printf("init.Plugin: %v\n", Plugin_Name)
}
