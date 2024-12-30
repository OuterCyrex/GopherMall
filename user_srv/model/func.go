package model

import (
	"GopherMall/user_srv/global"
	"fmt"
)

func findUserByField(field string) func(data interface{}) (User, error) {
	return func(data interface{}) (User, error) {
		var user User
		err := global.DB.Model(&User{}).Where(fmt.Sprintf("%s = ?", field), data).First(&user).Error
		return user, err
	}
}

var FindById = findUserByField("id")
var FindByMobile = findUserByField("mobile")
