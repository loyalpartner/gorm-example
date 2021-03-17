package main

import (
	"github.com/icrowley/fake"
	"github.com/loyalpartner/gorm-example/database"
	"github.com/loyalpartner/gorm-example/model"
)

func test_create_user_with_credit_card() {
	database.DB.Debug().Create(&model.User{
		Name: fake.FullName(),
		CreditCard: model.CreditCard{
			Number: fake.CreditCardNum(""),
		},
	})
}
