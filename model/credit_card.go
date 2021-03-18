package model

type CreditCard struct {
	Model
	Number string
	UserID string `gorm:"type:varchar(32);"`
}
