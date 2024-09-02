package ent

type User struct {
	Id      int    `gorm:"column:id"`
	Name    string `gorm:"column:name"`
	Age     int    `gorm:"column:age"`
	Address string `gorm:"column:address"`
}

func (User) TableName() string {
	return "user"
}
