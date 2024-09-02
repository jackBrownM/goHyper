package dao

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

func (m *Admin) GetUser() string {
	return "user:xxxxxx"
}
