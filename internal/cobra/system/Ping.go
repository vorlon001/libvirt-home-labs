package system

import (
        "net"
        "time"
        "github.com/sirupsen/logrus"
        logs "gitlab.iblog.pro/cobra/libvirt/internal/cobra/logs"
)

func Ping(network, address string, timeout time.Duration) error {
	logs.Log.WithFields(logrus.Fields{ "network": network, "address": address, "timeout":timeout, }).Info("Ping")
        conn, err := net.DialTimeout(network, address, timeout)
        defer conn.Close()
        if conn != nil {
                return nil
        }
        return err
}
