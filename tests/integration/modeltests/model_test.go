package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/app/models"
	"github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/app/startup"
	"github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/config"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = startup.App{}
var userInstance = models.User{}
var postInstance = models.Post{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	config := config.GetConfig()
	Database(config)
	os.Exit(m.Run())
}

func Database(config *config.Config) {
	var err error

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	server.DB, err = gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	} else {
		fmt.Printf("database %s connected successfully", dbURI)
	}
}

func refreshUserTable() error {
	server.DB.Exec("SET foreign_key_checks=0")
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	server.DB.Exec("SET foreign_key_checks=1")
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	log.Printf("refreshUserTable routine OK !!!")
	return nil
}

func seedOneUser() (models.User, error) {

	_ = refreshUserTable()

	user := models.User{
		Name: "saray005",
		Username: "chaksaray005",
		Password: "password",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}

	log.Printf("seedOneUser routine OK !!!")
	return user, nil
}

func seedUsers() error {

	users := []models.User{
		models.User{
			Name: "saray003",
			Username: "chaksaray003",
			Password: "password",
		},
		models.User{
			Name: "saray004",
			Username: "chaksaray004",
			Password: "password",
		},
	}

	for i := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}

	log.Printf("seedUsers routine OK !!!")
	return nil
}

func refreshUserAndPostTable() error {

	server.DB.Exec("SET foreign_key_checks=0")
	// NOTE: when deleting first delete Post as Post is depending on User table
	err := server.DB.DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		return err
	}
	server.DB.Exec("SET foreign_key_checks=1")
	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	log.Printf("refreshUserAndPostTable routine OK !!!")
	return nil
}

func seedOneUserAndOnePost() (models.Post, error) {

	err := refreshUserAndPostTable()
	if err != nil {
		return models.Post{}, err
	}
	user := models.User{
		Name: "saray006",
		Username: "chaksaray006",
		Password: "password",
	}
	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Post{}, err
	}
	post := models.Post{
		Title:    "This is the title sam",
		Content:  "This is the content sam",
		AuthorID: user.ID,
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}

	log.Printf("seedOneUserAndOnePost routine OK !!!")
	return post, nil
}

func seedUsersAndPosts() ([]models.User, []models.Post, error) {

	var err error

	if err != nil {
		return []models.User{}, []models.Post{}, err
	}
	var users = []models.User{
		models.User{
			Name: "saray001",
			Username: "chaksaray001",
			Password: "password",
		},
		models.User{
			Name: "saray002",
			Username: "chaksaray002",
			Password: "password",
		},
	}
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

	for i := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
	log.Printf("seedUsersAndPosts routine OK !!!")
	return users, posts, nil
}
