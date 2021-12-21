package service

import (
	"fmt"

	"github.com/tkhs1121/go-sns/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := "root:root@tcp(db:3306)/dbname?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("DB Error")

		return
	}

	fmt.Println("init ok")
}

func Migrate() {
	DB.AutoMigrate(
		model.User{},
		model.Profile{},
		model.Upvote{},
	)
}
