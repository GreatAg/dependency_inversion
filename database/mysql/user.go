package mysql

import (
	controller_models "dependency_inversion/models/controller"
	database_models "dependency_inversion/models/database"
)

type UserMysqlRepository struct {
}

func (u *UserMysqlRepository) GetUser(id int) controller_models.User {
	user := database_models.User{
		Id:            id,
		Name:          "ali",
		Email:         "ali@gmail.com",
		Number:        "09140903378",
		National_code: "4420926120",
	}
	return user.Present()
}
