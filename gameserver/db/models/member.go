package models

import "github.com/jinzhu/gorm"

// Member ...
type Member struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"type:varchar(100);unique_index;not null" json:"email"`
}

// TableName returns the custom table name.
func (Member) TableName() string {
	return "member"
}
