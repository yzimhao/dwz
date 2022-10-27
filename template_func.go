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
		"configs":  configs,
		"year":     year,
		"unsafe":   unsafe,
	}
}

func configs(key string) string {
	v := ""
	switch key {
	case "version":
		v = pack.Version
	default:
		v = viper.GetString(key)
	}
	return v
}

func year() string {
	now := time.Now().Format("2006")
	return now
}

func unsafe(str string) template.HTML {
	return template.HTML(str)
}
