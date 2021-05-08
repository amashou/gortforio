package config

import (
	"os"

)

type ConfigList struct {
	ApiKey string
	ApiSecret string
}
type TestConfigList struct {
	TestKey string
	TestSecret string
}
var TestConfig TestConfigList;

var Config ConfigList


func init() {
	cfg, err := ini.Load('')
}

// func init() {
// 	Config = ConfigList{
// 		ApiKey: os.Getenv("BITFLYER_API_KEY"),
// 		ApiSecret: os.Getenv("BITFLYER_API_SECRET"),
// 	}
// }

