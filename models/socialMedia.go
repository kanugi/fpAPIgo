package models

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaUrl string `gorm:"not null" json:"socialMediaUrl" form:"socialMediaUrl"`
	UserId         uint   `gorm:"not null" json:"userId" form:"userId"`
}
