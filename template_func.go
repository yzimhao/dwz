package dwz

import (
	"html/template"
	"time"

	"github.com/spf13/viper"
	"github.com/yzimhao/utilgo/pack"
)

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"appInfo": appInfo,
		"year":    year,
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

func year() string {
	now := time.Now().Format("2006")
	return now
}
