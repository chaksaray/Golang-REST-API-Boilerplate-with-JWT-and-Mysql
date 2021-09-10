package app

import (
	"log"
	"os"
	"skeleton_project/app/startup"
	"skeleton_project/config"
	"skeleton_project/mysql/seeds"

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