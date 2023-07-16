package main

import (

        "github.com/sirupsen/logrus"
)


func (j *VMDisk) CreatePath(configAbstract interface{}) error {

    if len(j.Path)==0 {

        render, err := templateRender(j.Tmpl, configAbstract)
        if err != nil {
		log.WithFields(logrus.Fields{ "err": err, }).Info("CreatePath")
                return err
        }
        j.Path = render
    }
    return nil
}
