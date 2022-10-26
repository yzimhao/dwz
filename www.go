package dwz

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"net/url"
	"regexp"

	"github.com/BurntSushi/toml"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/marksalpeter/token"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

//go:embed templates/*
var emfs embed.FS

var (
	rdc *redis.Client
)

func Run() {
	startGin()
}

func startGin() {
	viper.SetDefault("app.host", "0.0.0.0:8080")
	viper.SetDefault("redis.host", "0.0.0.0:6379")
	rdc = redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.host"),
	})

	router := SetupRouter(emfs)
	router.Run(viper.GetString("app.host"))
}

// router.Use(GinI18nLocalize())
func GinI18nLocalize() gin.HandlerFunc {
	return ginI18n.Localize(
		ginI18n.WithBundle(&ginI18n.BundleCfg{
			RootPath:         "./lang",
			AcceptLanguage:   []language.Tag{language.Chinese, language.English},
			DefaultLanguage:  language.English,
			FormatBundleFile: "toml",
			UnmarshalFunc:    toml.Unmarshal,
		}),
		ginI18n.WithGetLngHandle(
			func(context *gin.Context, defaultLng string) string {
				lng := context.Query("lang")
				if lng == "" {
					return defaultLng
				}
				return lng
			},
		),
	)
}

func SetupRouter(emfs embed.FS) *gin.Engine {
	router := gin.Default()

	router.Use(GinI18nLocalize())

	if viper.GetBool("app.debug") {
		gin.SetMode(gin.DebugMode)
		router.SetFuncMap(templateFuncMap())
		router.LoadHTMLGlob("../templates/default/*.html")
		router.StaticFS("statics", http.Dir("../templates/default/statics"))
	} else {
		// embed files
		tmpl := template.New("").Funcs(templateFuncMap())
		tmpl = template.Must(tmpl.ParseFS(emfs, "templates/default/*.html"))
		router.SetHTMLTemplate(tmpl)
		fp, _ := fs.Sub(emfs, "templates/default/statics")
		router.StaticFS("/statics", http.FS(fp))
	}
	router.GET("/", index)
	router.GET("/ping", ping)
	router.GET("/:key", redirect)
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/create", create)
		apiV1.GET("/index_links", index_links)
	}
	return router
}

func ping(c *gin.Context) {
	c.String(200, "pong")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func redirect(c *gin.Context) {
	code := c.Param("key")

	url := rdc.Get(rdc.Context(), "code:"+code).Val()

	if url == "" {
		fmt.Println(url)
	}
	c.Redirect(http.StatusMovedPermanently, url)
}

type resp struct {
	Ok   bool        `json:"ok"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func create(c *gin.Context) {
	callback := ""
	paramUrl := ""
	if c.Request.Method == "GET" {
		callback = c.Query("callback")
		paramUrl, _ = url.QueryUnescape(c.Query("url"))
	}
	if c.Request.Method == "POST" {
		paramUrl = c.PostForm("url")
	}

	pattern := `^((https|http|ftp|rtsp|mms)?://)?(([0-9a-z_!~*\'().&=+$%-]+: )?[0-9a-z_!~*\'().&=+$%-]+@)?`
	reg := regexp.MustCompile(pattern)
	if reg.FindString(paramUrl) == "" {
		result(c, callback, resp{Ok: false, Msg: "网址不正确"})
		return
	}

	code := token.New(5).Encode()
	data := map[string]interface{}{
		"dwz": code,
	}

	rdc.Set(rdc.Context(), "code:"+code, paramUrl, -1)
	rdc.SAdd(rdc.Context(), "urls", code)

	result(c, callback, resp{Ok: true, Data: &data})
}

func index_links(c *gin.Context) {
	callback := c.Query("callback")
	cmd := rdc.SRandMemberN(rdc.Context(), "urls", 100)
	rows := cmd.Val()
	result(c, callback, resp{Ok: true, Data: rows})
}

func result(c *gin.Context, callback string, resp2 resp) {
	if callback == "" {
		c.JSON(http.StatusOK, &resp2)
	} else {
		c.JSONP(http.StatusOK, &resp2)
	}
}
