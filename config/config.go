package config

import (
	"fmt"
	"math/rand"
	logger "reinvent/pkg"

	"github.com/kelseyhightower/envconfig"
)

type (
	AppConfig struct {
		HttpHost  int `envconfig:"http_host"`
		HttpPort  int `envconfig:"http_port"`
		LogConfig struct {
			Logtype    string `envconfig:"type"`
			LogTagName string `envconfig:"tagname"`
		} `envconfig:"log"`
	}
)

var (
	log = logger.GetLogger()
	cfg *AppConfig
)

func init() {
	rand.Seed(rand.Int63())
}

func LoadConfig() *AppConfig {
	var xfg AppConfig
	err := envconfig.Process(AppName, &xfg)
	if err != nil {
		panic(fmt.Sprintf("cannot read config: %s", err))
	}

	cfg = &xfg
	return cfg
}

func GetConfig() *AppConfig {
	if cfg == (&AppConfig{}) {
		LoadConfig()
	}

	return cfg
}
