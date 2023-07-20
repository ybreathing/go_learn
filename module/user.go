package module

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Phone    string `gorm:"type:varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}
