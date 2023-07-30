package disk

import (

        "github.com/sirupsen/logrus"
        logs "iblog.pro/cobra/logs"
	coreUtils "iblog.pro/cobra/core/utils"
)

type VMDisk struct {
        Path     string `yaml:"path"`
        Tmpl string `yaml:"tmpl"`
}

func (j *VMDisk) CreatePath(configAbstract interface{}) error {

    if len(j.Path)==0 {

        render, err := coreUtils.TemplateRender(j.Tmpl, configAbstract)
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("CreatePath")
                return err
        }
        j.Path = render
    }
    return nil
}
