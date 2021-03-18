package model

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	Model
	Name string
}

func (u *Role) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = strings.ReplaceAll(uuid.NewString(), "-", "")
	return nil
}
