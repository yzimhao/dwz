package dwz

import (
	"html/template"

	"github.com/spf13/viper"
	"github.com/yzimhao/utilgo/pack"
)

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"appInfo": appInfo,
	}
}

func appInfo(key string) string {
	v := ""
	switch key {
	case "appname":
		v = viper.GetString("app.name")
	case "version":
		v = pack.Version
	}
	return v
}
