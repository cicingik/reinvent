package rest

import (
	"fmt"
	"net/http"

	"github.com/cicingik/reinvent/config"
	logger "github.com/cicingik/reinvent/pkg"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmchi"
)

type (
	reinvent func(mux *chi.Mux)

	DeliveryHttpEngine struct {
		mux *chi.Mux
		cfg config.AppConfig
	}
)

var (
	log = logger.GetLogger()
)

func NewHTTPServer(cfg config.AppConfig) *DeliveryHttpEngine {
	c := chi.NewMux()

	trc, err := apm.NewTracer("", config.AppVersion)
	if err != nil {
		log.Errorf("error initialing APM tracer. %s", err)
	} else {
		log.Infoln("initialing APM middleware for Chi")
		c.Use(apmchi.Middleware(apmchi.WithTracer(trc)))
	}

	c.Use(middleware.RequestID)
	c.Use(middleware.AllowContentType("application/json"))
	c.Use(middleware.RealIP)
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	return &DeliveryHttpEngine{
		mux: c,
		cfg: cfg,
	}
}

func (h *DeliveryHttpEngine) RegisterHandler(r reinvent) {
	r(h.mux)
}

func (h *DeliveryHttpEngine) Serve() error {
	err := h.initRoute()
	if err != nil {
		return err
	}

	binding := fmt.Sprintf("%s:%d", h.cfg.HttpHost, h.cfg.HttpPort)

	log.Infoln("Running HTTP server in %s", binding)

	return http.ListenAndServe(binding, h.mux)
}
