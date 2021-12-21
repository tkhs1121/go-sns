package model

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Url  string
	UserId uint
	Upvote []Upvote
}
