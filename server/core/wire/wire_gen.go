// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"goHyper/core/router"
	"goHyper/internal/api/admin"
	"goHyper/internal/controller/admin"
	"goHyper/internal/dao"
	"goHyper/internal/logic"
	svc2 "goHyper/internal/svc"
	base2 "goHyper/internal/svc/base"
)

// Injectors from wire.go:

func InitializeSvc() (*svc2.Init, error) {
	config, err := base2.NewConfig()
	if err != nil {
		return nil, err
	}
	logger, err := base2.NewLogger(config)
	if err != nil {
		return nil, err
	}
	httpServ, err := svc2.NewHttpServ(config, logger)
	if err != nil {
		return nil, err
	}
	db, err := base2.NewMysql(config, logger)
	if err != nil {
		return nil, err
	}
	admin := dao.NewAdmin(db)
	role := dao.NewRole(db)
	logicAdmin := logic.NewAdmin(admin, role, config)
	ctr_adminAdmin := ctr_admin.NewSystem(logicAdmin)
	perm := dao.NewPerm(db)
	logicRole := logic.NewRole(perm, admin, role, config)
	ctr_adminRole := ctr_admin.NewRole(logicRole)
	route_adminAdmin := route_admin.NewAdmin(ctr_adminAdmin, ctr_adminRole)
	route, err := router.NewRoute(config, logger, route_adminAdmin)
	if err != nil {
		return nil, err
	}
	init, err := svc2.NewInit(logger, httpServ, route)
	if err != nil {
		return nil, err
	}
	return init, nil
}
