package CSI_Local

import (
	"fmt"

	csi "play.ground/CSI"
)

var _ csi.Csi = (*csi_local)(nil)
var _ = (csi.Csi)((*csi_local)(nil))

var csi_Name string

type csi_local struct {
	Data_csi_local string
}

func (a *csi_local) Print(z string) {
	fmt.Printf("Plugin.print: %#v %#v\n", z, a.Data_csi_local)
}

func newPlugin(z string) csi.Csi {
	return &csi_local{Data_csi_local: z}
}

func init() {

	csi_Name = "csi_csi_local"
	csi.RegisterStruct(csi_Name, newPlugin("csi_csi_local struct"))
        fmt.Printf("init.Plugin: %v\n", csi_Name)
}
