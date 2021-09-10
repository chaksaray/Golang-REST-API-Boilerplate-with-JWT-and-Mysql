package seeds

import (
	"log"

	"github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/app/models"

	"github.com/jinzhu/gorm"
)

func up(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}

func down(db *gorm.DB) {
	err := db.DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
}

func Run(db *gorm.DB) {
	down(db)
	up(db)

	// seeding data
	for i, _ := range users {
		err := db.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}