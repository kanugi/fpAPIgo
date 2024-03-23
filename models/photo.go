package models

type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" form:"title"`
	Caption  string    `gorm:"not null" json:"caption" form:"caption"`
	PhotoUrl string    `gorm:"not null" json:"photoUrl" form:"photoUrl"`
	UserId   uint      `gorm:"not null" json:"userId" form:"userId"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}
