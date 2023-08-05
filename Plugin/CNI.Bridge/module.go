package cni_Bridge

import (
	"fmt"

	cni "play.ground/CNI"
)

var _ cni.Cni = (*cni_Bridge)(nil)
var _ = (cni.Cni)((*cni_Bridge)(nil))

var Plugin_Name string

type cni_Bridge struct {
	Data_cni_Bridge string
}

func (a *cni_Bridge) Print(z string) {
	fmt.Printf("Plugin.print: %#v %#v\n", z, a.Data_cni_Bridge)
}
func newPlugin(z string) cni.Cni {
	return &cni_Bridge{Data_cni_Bridge: z}
}
func init() {
	Plugin_Name = "cni_cni_Bridge"
	cni.RegisterStruct(Plugin_Name, newPlugin("cni_cni_Bridge struct"))
	fmt.Printf("init.Plugin: %v\n", Plugin_Name)
}
