package logic

import (
	"goHyper/core/svc/base"
	"goHyper/internal/dao"
)

type Menu struct {
	menu   *dao.Menu
	config *base.Config
}

func NewMenu(menu *dao.Menu, config *base.Config) *Menu {
	return &Menu{
		menu:   menu,
		config: config,
	}
}
