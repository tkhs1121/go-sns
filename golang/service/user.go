package service

import (
	"github.com/tkhs1121/go-sns/model"
)


func Register(link string) (uint, error) {

	var user model.User

	tx := DB.Begin()

	if result := tx.Create(&user); result.Error != nil {
		tx.Rollback()

		return 0, result.Error
	}

	profile := model.Profile{UserId: user.ID, Url: link}

	if result := tx.Create(&profile); result.Error != nil {
		tx.Rollback()

		return 0, result.Error
	}

	tx.Commit()

	return user.ID, nil
}
