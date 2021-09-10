package app

import (
	"log"
	"os"

	"github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/app/startup"
	"github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/config"
	"github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/mysql/seeds"

	"github.com/joho/godotenv"
)

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	logger := &startup.Log{}
	logger.InitialLog()

	startup.InfoLog.Println("Starting the application...")

	config := config.GetConfig()
	app := &startup.App{}
	app.Initialize(config)

	// seeding data to database
	seeds.Run(app.DB)

	app.Run(":" + os.Getenv("PORT"))
}