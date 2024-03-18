package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rest/models"
)

type QuestRepository interface {
	SaveQuest(quest models.Quest)
	UpdateQuest(id string, quest models.Quest)
}

func NewQuestRepository() QuestRepository {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=admin dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	fmt.Println(*db)
	if err != nil {
		fmt.Println(err)
		panic("No DB")
	}

	db.AutoMigrate(&models.Quest{})

	return &database{
		connection: *db,
	}
}

func (db *database) SaveQuest(quest models.Quest) {
	db.connection.Create(&quest).Scan(&quest)
}

func (db *database) UpdateQuest(id string, user models.Quest) {
	//db.connection.Model(&user).Where("id = ?", id).Update("balance", user.Balance)
}
