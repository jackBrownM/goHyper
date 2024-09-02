package logic

import "goHyper/internal/model"

type Admin struct {
	admin *model.Admin
}

func NewAdmin(admin *model.Admin) *Admin {
	return &Admin{
		admin: admin,
	}
}

func (l *Admin) Login() string {
	return l.admin.GetUser()
}
