package models

import {
	
	"time"

}

type Author struct {
	ID        uint      `json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Gender    string    `gorm:"type:char(1)" json:"Gender"`
	Email     string    `gorm:"type:varchar(100)" json:"Email"`
	Age       int       `gorm:"type:integer" json:"Age"`
	CreatedAt time.Time `json:"created_at gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at gorm:"autoUpdateTime"`
}
