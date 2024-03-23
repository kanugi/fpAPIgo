package models

import (
	"fpGOWebScalable/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username        string        `gorm:"not null;uniqueIndex" json:"username" form:"username"`
	Email           string        `gorm:"not null;uniqueIndex" json:"email"`
	Password        string        `gorm:"not null" json:"password" form:"password"`
	Age             uint          `gorm:"not null" json:"age" form:"age"`
	ProfileImageUrl string        `gorm:"not null" json:"profileImageUrl" form:"profileImageUrl"`
	Photos          []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comments        []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedias    []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socialMedias"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(user)

	if errCreate != nil {
		err = errCreate
		return
	}

	user.Password = helpers.HashPass(user.Password)
	err = nil
	return
}
