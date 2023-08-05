package CSI_CEPH

import (
	"fmt"

	csi "play.ground/CSI"
)

var _ csi.Csi = (*csi_ceph)(nil)
var _ = (csi.Csi)((*csi_ceph)(nil))

var Plugin_Name string

type csi_ceph struct {
	Data_csi_ceph string
}

func (a *csi_ceph) Print(z string) {
	fmt.Printf("Plugin.print: %#v %#v\n", z, a.Data_csi_ceph)
}
func newPlugin(z string) csi.Csi {
	return &csi_ceph{Data_csi_ceph: z}
}
func init() {
	Plugin_Name = "csi_csi_ceph"
	csi.RegisterStruct(Plugin_Name, newPlugin("csi_csi_ceph struct"))
	fmt.Printf("init.Plugin: %v\n", Plugin_Name)
}
