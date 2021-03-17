package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Author Author `gorm:"embedded;embeddedPrefix:author_"`
	Upvotes int32
}
