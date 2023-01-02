package query

import (
	"gorm.io/gorm"
	"gorm_pro/create"
)

func GoC(db *gorm.DB) *create.User {
	user := &create.User{}
	db.First(&user)
	return user
}

func GoC2(db *gorm.DB) map[string]interface{} {
	result := map[string]interface{}{}
	db.Model(&create.User{}).First(&result)
	return result
}
