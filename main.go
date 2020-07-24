package main

import (
	"os"

	"github.com/cicingik/reinvent/app"
	"github.com/cicingik/reinvent/config"
	logger "github.com/cicingik/reinvent/pkg"

	"github.com/amoghe/distillog"
	env "github.com/joho/godotenv"
)

var (
	cfg *config.AppConfig
	log = distillog.NewStderrLogger(config.AppName)
)

func init() {
	_, err := os.Stat(".env")
	if err == nil {
		log.Infoln("loading ENV_VAR from .env")
		if err := env.Load(); err != nil {
			log.Errorf(err.Error())
		}
	}

	cfg = config.LoadConfig()
	log = logger.GetLogger()

	logger.InitLogger(cfg.LogConfig.Logtype, cfg.LogConfig.LogTagName)

}

func main() {
	apx, err := app.New(cfg)

	if err != nil {
		log.Errorf("error initializing application: %s", err)
		os.Exit(1)
	}

	if err := apx.Start(); err != nil {
		log.Errorf("found error: %s", err)
	}

	log.Infoln("done.")
}
