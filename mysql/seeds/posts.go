package seeds

import "github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/app/models"

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}