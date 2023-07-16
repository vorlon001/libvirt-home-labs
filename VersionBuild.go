package main

import (
	"github.com/sirupsen/logrus"
)

func VersionBuild(version string) {
	log.WithFields(logrus.Fields{ "Version": version,}).Info("https://github.com/vorlon001, (C) Vorlon001")
	log.Info("HomeLabs LibVirt Connector (Golang version)")
}

