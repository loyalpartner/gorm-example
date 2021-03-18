package model

import (
	"database/sql"
	"time"
)

type User struct {
	Model
	Name         string
	Age          byte `gorm:"default:10"`
	Email        *string
	Birthday     *time.Time
	MemberNumber sql.NullString
	CreditCard   CreditCard `gorm:"foreignKey:UserID;references:ID"`
	CompanyID    int        `gorm:"default:null"`
	Company      Company    `gorm:"foreignKey:CompanyID"`
	Roles        []Role     `gorm:"many2many:user_roles"`
}

