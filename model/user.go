package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Age          byte `gorm:"default:10"`
	Email        *string
	Birthday     *time.Time
	MemberNumber sql.NullString
	CreditCard   CreditCard
	CompanyID    int
	Company      Company `gorm:"foreignKey:CompanyID"`
}
