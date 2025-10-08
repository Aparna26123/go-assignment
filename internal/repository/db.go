// connects to sql database, initializes global db variables, and automatically creates tables,boards,list for user
package repository

import (
	"task-manager/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("taskmanager.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.User{}, &models.Board{}, &models.List{}, &models.Task{})

}
