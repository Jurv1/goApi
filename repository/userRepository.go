package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rest/models"
)

type UserRepository interface {
	Save(user models.User)
	Update(id string, user models.User)
}

type database struct {
	connection gorm.DB
}

func NewUserRepository() UserRepository {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=admin dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	fmt.Println(*db)
	if err != nil {
		fmt.Println(err)
		panic("No DB")
	}

	db.AutoMigrate(&models.User{})

	return &database{
		connection: *db,
	}
}

func (db *database) Save(user models.User) {
	db.connection.Create(&user).Scan(&user)
}

func (db *database) Update(id string, user models.User) {
	db.connection.Model(&user).Where("id = ?", id).Update("balance", user.Balance)
}
