package CSI

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

type Csi interface {
	Print(string)
}

var (
	driverCsi  sync.Mutex
	AtomicFormats atomic.Value
)

func RegisterStruct(name string, z Csi) {

	driverCsi.Lock()
	defer driverCsi.Unlock()
	formats, _ := AtomicFormats.Load().(map[string]Csi)
	formats[name] = z
	AtomicFormats.Store(formats)
}

func GetRegisterStruct(name string) (Csi, error) {
	driverCsi.Lock()
	defer driverCsi.Unlock()

	Csis := AtomicFormats.Load().(map[string]Csi)
	if val, ok := Csis[name]; ok {
		return val, nil
	}
	return nil, errors.New(fmt.Sprintf("not found driver: %s", name))
}

var Csi_Name string

func init() {
	AtomicFormats.Store(make(map[string]Csi, 0))
	Csi_Name = "Csi"
	fmt.Printf("init.Csi %#v\n", Csi_Name)
}
