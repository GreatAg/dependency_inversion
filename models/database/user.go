package database_models

import controller_models "dependency_inversion/models/controller"

type User struct {
	Id            int
	Name          string
	Email         string
	Number        string
	National_code string
}

func (u User) Present() controller_models.User {
	return controller_models.User(u)
}
