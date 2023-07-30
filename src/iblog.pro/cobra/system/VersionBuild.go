package system

import (
	"github.com/sirupsen/logrus"
        logs "iblog.pro/cobra/logs"
)

func VersionBuild(version string) {
	logs.Log.WithFields(logrus.Fields{ "Version": version,}).Info("https://github.com/vorlon001, (C) Vorlon001")
	logs.Log.Info("HomeLabs LibVirt Connector (Golang version)")
}

