package main

import (

        "encoding/json"
        "os"
	"fmt"
        "strings"
        "runtime"
        "github.com/sirupsen/logrus"
)


func PanicRecover() {
    if r := recover(); r != nil {
        logrus.Error("Internal error: %v", r)
        buf := make([]byte, 1<<16)
        stackSize := runtime.Stack(buf, true)
        log.Error("--------------------------------------------------------------------------------")
        log.Error(fmt.Sprintf("Internal error: %s\n", string(buf[0:stackSize])))
        log.Error("--------------------------------------------------------------------------------")

        }
}



func herr(e error) {
        if e != nil {
                log.WithFields(logrus.Fields{ "err": strings.ReplaceAll(e.Error(), "\"", ""), }).Info("herr")
                os.Exit(1)
        }
}

func hok(message string) {
        log.WithFields(logrus.Fields{ "ok": strings.ReplaceAll(message, "\"", ""), }).Info("hok")
}

func hret(i interface{}) {
        ret, err := json.Marshal(i)
        herr(err)
        log.WithFields(logrus.Fields{ "ret": string(ret), }).Info("Inside hret Run")
}
