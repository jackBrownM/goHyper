package svc

import (
	"goHyper/core/router"
	"goHyper/internal/svc/base"
	"goHyper/libs/fiberLib"
)

type Init struct {
	logger   *base.Logger
	httpServ *fiberLib.HttpServ
}

func NewInit(logger *base.Logger, httpServ *fiberLib.HttpServ, router *router.Route) (*Init, error) {
	router.Register(httpServ.App)
	s := Init{
		logger:   logger,
		httpServ: httpServ,
	}
	return &s, nil
}

func (s *Init) Start() {
	s.httpServ.Start()
}

func (s *Init) Stop() {
	s.logger.Info("base stopping....")
	s.logger.Info("base running cleanup tasks...")
	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	s.logger.Info("base successful stopped.")
}
