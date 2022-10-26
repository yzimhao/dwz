package dwz

import (
	"html/template"
	"time"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/spf13/viper"
	"github.com/yzimhao/utilgo/pack"
)

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"Localize": ginI18n.GetMessage,
		"appInfo":  appInfo,
		"year":     year,
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
