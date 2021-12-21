package service

import (
	"github.com/tkhs1121/go-sns/model"
)

func GetRandProfile(id uint) (uint, error) {

	profile := model.Profile{}

	if err := DB.Select("id").Where("id != ?", id).Order("RAND()").First(&profile).Error; err != nil {

		return 0, err
	}

	return profile.ID, nil
}

func UpdateRecommendation(userID uint, link string) error {

	if err := DB.Model(&model.Profile{}).Where("id = ?", userID).Update("url", link).Error; err != nil {

		return err
	}

	return nil
}
