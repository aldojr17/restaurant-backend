package initialize

import (
	"final-project-backend/config"
	"final-project-backend/logger"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Application struct {
	DB *gorm.DB
}

func App() *Application {
	if err := config.LoadConfig(); err != nil {
		handleInitError(err)
	}

	var app = &Application{
		DB: initializeDB(),
	}

	return app
}

func initializeDB() *gorm.DB {
	dbConfig := config.Database()

	db, err := gorm.Open(postgres.Open(dbConfig.String()), &gorm.Config{
		Logger: logger.NewGormLogger(),
	})
	if err != nil {
		panic(err)
	}

	log.Println("connected to database")
	return db
}

func handleInitError(e error) {
	if e != nil {
		panic(e)
	}
}
