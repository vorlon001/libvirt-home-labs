package CSI_ZFS

import (
	"fmt"

	csi "play.ground/CSI"
)

var _ csi.Csi = (*ceph_zfs)(nil)
var _ = (csi.Csi)((*ceph_zfs)(nil))

var csi_Name string

type ceph_zfs struct {
	Data_ceph_zfs string
}

func (a *ceph_zfs) Print(z string) {
	fmt.Printf("Plugin.print: %#v %#v\n", z, a.Data_ceph_zfs)
}

func newPlugin(z string) csi.Csi {
	return &ceph_zfs{Data_ceph_zfs: z}
}

func init() {

	csi_Name = "csi_ceph_zfs"
	csi.RegisterStruct(csi_Name, newPlugin("csi_ceph_zfs struct"))
	fmt.Printf("init.Plugin 2 %#v\n", csi_Name)
}
