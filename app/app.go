package app

import (
	"fmt"

	"github.com/cicingik/reinvent/config"
	logger "github.com/cicingik/reinvent/pkg"
	"github.com/cicingik/reinvent/service/delivery/rest"
	"github.com/cicingik/reinvent/service"
)

var (
	log = logger.GetLogger()
)

type Application struct {
	cfg        *config.AppConfig
	HttpServer *rest.DeliveryHttpEngine
}

func New(cfg *config.AppConfig) (*Application, error) {
	apx := &Application{
		cfg:        cfg,
		HttpServer: rest.NewHTTPServer(*cfg),
	}

	err := apx.initService()
	return apx, err
}

func (a *Application) validate() error {
	if a.cfg.HttpPort <= 0 {
		return fmt.Errorf("invalid port: %d", a.cfg.HttpPort)
	}

	return nil
}

func (a *Application) Start() error {
	err := a.validate()

	if err != nil {
		return err
	}

	return a.HttpServer.Serve()
}
