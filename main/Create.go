package main

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func CreateUser(db *gorm.DB) {
	//db.AutoMigrate(&User{})
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	db.Create(&user) // 通过数据的指针来创建

	db.Select("Name", "Age", "CreatedAt").Create(&user)
}

func CreateUserMul(db *gorm.DB) {
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

	for _, user := range users {
		println(user.ID)
	}
}
