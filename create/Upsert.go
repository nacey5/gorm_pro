package create

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User struct {
	gorm.Model
	Name       string `gorm:"default:galeone"`
	CreditCard CreditCard
}

func GoC() {
	db.Create(&User{
		Name:       "jinzhu",
		CreditCard: CreditCard{Number: "4546542132132"},
	})
}

// 可以通过 Select、 Omit 跳过关联保存
func GoCOmit() {
	user := &User{}
	db.Omit("CreditCard").Create(&user)

	// 跳过所有关联
	db.Omit(clause.Associations).Create(&user)

}
