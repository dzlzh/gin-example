package models

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	Account  string
	Password string
	Role     uint
}

// func (Admin) TableName() string {
//     return "admin"
// }
