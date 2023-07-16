package main

import (
        "net"
        "time"
        "github.com/sirupsen/logrus"
)

func Ping(network, address string, timeout time.Duration) error {
	log.WithFields(logrus.Fields{ "network": network, "address": address, "timeout":timeout, }).Info("Ping")
        conn, err := net.DialTimeout(network, address, timeout)
        if conn != nil {
                defer conn.Close()
        }
        return err
}
