package database

import (
	"backEnd/models"
	"backEnd/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.ConnDB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Product{},
		&models.Cart{},
		&models.Transaction{},
		&models.ProductTransaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
