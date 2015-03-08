package config

import (
	"code.google.com/p/gcfg"
	"fmt"
)

type Config struct {
	Softlayer struct {
		User string
		Key  string
	}
	Slack struct {
		Target     string
		Slacktoken string
	}
}

func LoadConfig() Config {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if err != nil {
		fmt.Println(err)
	}
	return cfg
}
