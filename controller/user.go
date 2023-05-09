package controller

import (
	"dependency_inversion/database/mysql"
	api_models "dependency_inversion/models/api"
	controller_models "dependency_inversion/models/controller"

	"github.com/spf13/cast"
)

type Iuser interface {
	GetUser(id int) controller_models.User
}

type UserRepository struct {
}

func newUserRepository() Iuser {
	return &mysql.UserMysqlRepository{}
}

func (u *UserRepository) GetUser(id string) api_models.User {
	return newUserRepository().GetUser(cast.ToInt(id)).Present()
}
