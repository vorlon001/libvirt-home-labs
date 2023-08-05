package Core

import (
	csi "play.ground/CSI"
	_ "play.ground/Plugin/CSI.Ceph"
	_ "play.ground/Plugin/CSI.Local"
	_ "play.ground/Plugin/CSI.ZFS"

        cni "play.ground/CNI"
        _ "play.ground/Plugin/CNI.OVS"
        _ "play.ground/Plugin/CNI.Bridge"
        _ "play.ground/Plugin/CNI.OVN"

	"math/rand"
	"time"
)

type core struct {

	csiCeph   csi.Csi
	csiZFS    csi.Csi
        csiLocal  csi.Csi

	cniOvn    cni.Cni
	cniOvs    cni.Cni
	cniBridge cni.Cni

	rand    *rand.Rand
}

func (c *core) GetCSICeph() csi.Csi {
	return c.csiCeph
}
func (c *core) GetCSILocal() csi.Csi {
	return c.csiLocal
}

func (c *core) GetCSIZfs() csi.Csi {
        return c.csiZFS
}


func (c *core) GetCNIOvn() cni.Cni {
        return c.cniOvn
}
func (c *core) GetCNIOvs() cni.Cni {
        return c.cniOvs
}

func (c *core) GetCNIBridge() cni.Cni {
        return c.cniBridge
}


func (c *core) GetRand() *rand.Rand {
	return c.rand
}

type Core interface {
	GetCSICeph() csi.Csi
	GetCSILocal() csi.Csi
	GetCSIZfs() csi.Csi
	GetCNIOvn() cni.Cni
	GetCNIOvs() cni.Cni
	GetCNIBridge() cni.Cni
	GetRand() *rand.Rand
}

var _core core

func GetCore() Core {
	return &_core
}
func init() {
	csi_ceph, _ := csi.GetRegisterStruct("csi_csi_ceph")
	csi_zfs, _ := csi.GetRegisterStruct("csi_ceph_zfs")
        csi_local, _ := csi.GetRegisterStruct("csi_csi_local")


        cni_ovn, _ := cni.GetRegisterStruct("cni_cni_ovn")
        cni_ovs, _ := cni.GetRegisterStruct("cni_cni_Ovs")
        cni_bridge, _ := cni.GetRegisterStruct("cni_cni_Bridge")

	source_random := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source_random)
	_core = core{
		csiCeph:  csi_ceph,
		csiZFS:   csi_zfs,
		csiLocal: csi_local,
		cniOvn:   cni_ovn,
		cniOvs:   cni_ovs,
		cniBridge: cni_bridge,
		rand:    random,
	}
}
