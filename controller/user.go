package controller

import (
	"dependency_inversion/database/mysql"
	api_models "dependency_inversion/models/api"
	controller_models "dependency_inversion/models/controller"
	"fmt"

	"github.com/spf13/cast"
)

type Iuser interface {
	GetUser(int) controller_models.User
	UpdateUser(controller_models.User)
}

type UserRepository struct {
}

func newUserRepository() Iuser {
	return &mysql.UserMysqlRepository{}
}

func (u *UserRepository) GetUser(id string) api_models.User {
	return newUserRepository().GetUser(cast.ToInt(id)).Present()
}

func (u *UserRepository) UpdateUser(user api_models.User) {
	user_model := &controller_models.User{}
	user_model.Convert(user)
	fmt.Printf("in controller %+v", user_model)
	newUserRepository().UpdateUser(*user_model)
}
