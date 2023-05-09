package api_models

type User struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Number        string `json:"number"`
	National_code string `json:"national_code"`
}
