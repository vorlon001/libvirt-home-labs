package utils

import (
        "bytes"
        "strings"
        "github.com/sirupsen/logrus"
        "html/template"
        logs "gitlab.iblog.pro/cobra/libvirtgrpc/internal/cobra/logs"
)

func TemplateRender(templateMust string, configAbstract interface{}) (string, error) {

        tpl := template.Must(template.New("").Parse(templateMust))
        var tplBuffer bytes.Buffer
        if err := tpl.Execute(&tplBuffer, configAbstract); err != nil {
                logs.Log.WithFields(logrus.Fields{ "err": err, }).Info("templateRender")
                return  "",err
        }
        render := tplBuffer.String()
        replaceContent := [][]string{ []string{"&#34;","\""}, []string{"&#43;","+"}}
        for _,v := range replaceContent {
                render = strings.ReplaceAll( render, v[0], v[1])
        }
        return render,nil

}
