package models

type Author struct {
	ID uint `json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
	Gender string `gorm:"type:char(1)" json:"Gender"`
	Email string `gorm:"type:varchar(100)" json:"Email"`
	Age int `gorm:"type:integer" json:"Age"`
	CreatedAt string `json:"created_at`
	CreatedAt string `json:"updated_at`
}