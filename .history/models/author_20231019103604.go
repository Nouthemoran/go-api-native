package models

type Author struct {
	ID uint `json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
}