package mysql

import (
	controller_models "dependency_inversion/models/controller"
	database_models "dependency_inversion/models/database"
	"fmt"
)

type UserMysqlRepository struct { //unexport
}

//factory

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

func (u *UserMysqlRepository) UpdateUser(user controller_models.User) {
	user_db := &database_models.User{}
	user_db.Convert(user)
	fmt.Printf("updated successfully %+v", *user_db)
	// update
}
