package network

import (

        "crypto/rand"

	"fmt"
        "github.com/sirupsen/logrus"
        logs "gitlab.iblog.pro/cobra/libvirt/internal/cobra/logs"

)


type InterfaceName struct {
        Name     string `yaml:"name"`
        MacAddress string `yaml:"mac"`
        Slot    string  `yaml:"slot"`
}


func (j *InterfaceName) CreateMacAddress(MagicMac string) error {
    if len(j.MacAddress)==0 {
        buf := make([]byte, 3)
        _, err := rand.Read(buf)
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("CreateMacAddress")
                return err
        }
        buf[0] |= 2
        j.MacAddress = fmt.Sprintf("%s:%02x:%02x:%02x", MagicMac, buf[0], buf[1], buf[2])
    }
    return nil
}

func (j InterfaceName) GetMacAddress() string {
        return j.MacAddress
}
