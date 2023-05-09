package controller_models

import api_models "dependency_inversion/models/api"

type User struct {
	Id            int
	Name          string
	Email         string
	Number        string
	National_code string
}

func (u User) Present() api_models.User {
	return api_models.User(u)
}
