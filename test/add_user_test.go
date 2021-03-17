package test

import (
	"log"
	"testing"

	"github.com/icrowley/fake"
	"github.com/loyalpartner/gorm-example/database"
	"github.com/loyalpartner/gorm-example/model"
)

func TestAddUserWithCreditCard(t *testing.T) {
	age := new(byte)
	*age = 0
	user := &model.User{
		Name: fake.FullName(),
		Age: age,
		CreditCard: model.CreditCard{
			Number: fake.CreditCardNum(""),
		},
	}
	// *(user.Age) = byte(0);
	// t.Logf("hello%v", user.Age)
	result := database.DB.Debug().Create(user)

	if result.Error != nil {
		log.Fatalln("插入失败")
	}
	t.Logf("插入成功")
}
