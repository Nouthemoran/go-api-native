package models

type Author struct {
	ID uint `json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
	Gender string `gorm:"type:char(1)" json:"name"`
}