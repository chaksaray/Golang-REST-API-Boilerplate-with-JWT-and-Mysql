package seeds

import "github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/app/models"

var users = []models.User{
	models.User{
		Name: "saray",
		Username: "chaksaray",
		Password: "123456",
	},
	models.User{
		Name: "saray01",
		Username: "chaksaray01",
		Password: "123456",
	},
}