package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Title string
	Type  string
}
