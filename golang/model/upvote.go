package model

type Upvote struct {
	ID        int `gorm:"id PrimaryKey"`
	UserId    uint
	ProfileId uint
}
