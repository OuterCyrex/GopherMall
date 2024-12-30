package global

import (
	"GopherMall/user_srv/model"
	"fmt"
)

func findUserByField(field string) func(data interface{}) (model.User, error) {
	return func(data interface{}) (model.User, error) {
		var user model.User
		err := DB.Model(&model.User{}).Where(fmt.Sprintf("%s = ?", field), data).First(&user).Error
		return user, err
	}
}

var FindById = findUserByField("id")
var FindByMobile = findUserByField("mobile")
