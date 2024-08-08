package db

import (
	"log"
	"sync"

	"github.com/dayroromero/storiChallenge/pkg/models"
	"github.com/dayroromero/storiChallenge/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Handler struct {
	DB *gorm.DB
}

var (
	once     sync.Once
	instance Handler
)

func GetInstance() Handler {
	dbUrl := utils.GetEnvVar("DBSTRING_CONNECTION")
	once.Do(func() {
		instance = Init(dbUrl)
	})

	return instance
}

func Init(url string) Handler {
	log.Println("Init Database")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`) //Implements uuid_generate_v4()
	db.AutoMigrate(
		&models.Account{},
		&models.Transaction{},
		&models.User{},
	)

	return Handler{db}
}

func (h *Handler) Close() error {
	db, err := h.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
