package CNI

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

type Cni interface {
	Print(string)
}

var (
	driverCNI  sync.Mutex
	AtomicFormats atomic.Value
)

func RegisterStruct(name string, z Cni) {

	driverCNI.Lock()
	defer driverCNI.Unlock()
	formats, _ := AtomicFormats.Load().(map[string]Cni)
	formats[name] = z
	AtomicFormats.Store(formats)
}

func GetRegisterStruct(name string) (Cni, error) {
	driverCNI.Lock()
	defer driverCNI.Unlock()

	CNIs := AtomicFormats.Load().(map[string]Cni)
	if val, ok := CNIs[name]; ok {
		return val, nil
	}
	return nil, errors.New(fmt.Sprintf("not found driver: %s", name))
}

var CNI_Name string

func init() {
	AtomicFormats.Store(make(map[string]Cni, 0))
	CNI_Name = "CNI"
	fmt.Printf("init.CNI %#v\n", CNI_Name)
}
