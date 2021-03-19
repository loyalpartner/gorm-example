package model

import (
	"database/sql"
	"time"
)

type User struct {
	Model
	Name         string
	Age          byte `gorm:"default:10"`
	Email        string
	Birthday     *time.Time
	MemberNumber sql.NullString
	CreditCard   CreditCard `gorm:"foreignKey:UserID;references:ID"`
	CompanyID    string     `gorm:"default:null"`
	Company      Company    `gorm:"references:ID"`
	Roles        []Role     `gorm:"many2many:user_roles"`
}
