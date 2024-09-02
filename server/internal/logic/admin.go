package logic

import "goHyper/internal/dao"

type Admin struct {
	admin *dao.Admin
}

func NewAdmin(admin *dao.Admin) *Admin {
	return &Admin{
		admin: admin,
	}
}

func (l *Admin) Login() string {
	return l.admin.GetUser()
}
