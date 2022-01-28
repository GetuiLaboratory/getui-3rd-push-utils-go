package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Conf struct {
	XmAppId          string
	XmAppKey         string
	XmAppSecret      string
	OppoAppId        string
	OppoAppKey       string
	OppoAppSecret    string
	OppoMasterSecret string
}

func (c *Conf) GetConf(path string) *Conf {
	if path == "" {
		path = "config.toml"
	}
	if _, err := toml.DecodeFile(path, c); err != nil {
		log.Fatal(err)
	}
	return c
}
