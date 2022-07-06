package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
	cli "github.com/urfave/cli/v2"
	"github.com/yzimhao/dwz"
	"github.com/yzimhao/utilgo/pack"
)

func main() {
	viper.SetConfigFile("./config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	app := &cli.App{
		Name:  "dwz",
		Usage: "",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {

			dwz.Run()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show version",
				Action: func(c *cli.Context) error {
					pack.ShowVersion()
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
