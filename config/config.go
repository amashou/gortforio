package config

import (
	"os"

)

type ConfigList struct {
	ApiKey string
	ApiSecret string
}

var Config ConfigList

func init() {
	Config = ConfigList{
		ApiKey: os.Getenv("BITFLYER_API_KEY"),
		ApiSecret: os.Getenv("BITFLYER_API_SECRET"),
	}
}

