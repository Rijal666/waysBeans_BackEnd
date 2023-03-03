package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ConnDB *gorm.DB

func DataBaseInit() {
	var err error
	DBurl := "root:@tcp(localhost:3306)/ways_beans?charset=utf8mb4&parseTime=True&loc=Local"
	ConnDB, err = gorm.Open(mysql.Open(DBurl), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("connected to database")
}