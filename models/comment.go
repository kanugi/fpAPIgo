package models

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message"`
	UserId  uint   `gorm:"not null" json:"userId" form:"userId"`
	PhotoId uint   `gorm:"not null" json:"photoId" form:"photoId"`
}
