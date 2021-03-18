package model

import (
	"strings"

	"github.com/google/uuid"
	"github.com/loyalpartner/gorm-example/database"
	"gorm.io/gorm"
)

func init(){
	db := database.DB

	db.Migrator().AlterColumn(&CreditCard{}, "ID")
	
}

type CreditCard struct {
	Model
	Number string
	UserID string `gorm:"type:varchar(32);"`
}


func (c *CreditCard) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = strings.ReplaceAll(uuid.NewString(), "-", "")
	return nil
}
