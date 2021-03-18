package test

import (
	"log"
	"math/rand"
	"testing"

	"github.com/icrowley/fake"
	"github.com/loyalpartner/gorm-example/database"
	"github.com/loyalpartner/gorm-example/model"
)

func TestAddUserWithCreditCard(t *testing.T) {
	user := randomUser()
	result := database.DB.Debug().Create(user)

	if result.Error != nil {
		log.Fatalln("插入失败")
	}
	t.Logf("插入成功")
}

func TestAddUserWithCompany(t *testing.T) {
	user := randomUserWithCompany()
	result := database.DB.Debug().Create(user)

	if result.Error != nil {
		log.Fatalln("插入失败")
	}
	t.Logf("插入成功")
}


//
// utils
// 
func randomUser() *model.User {
	randAge := rand.Intn(99)
	return &model.User{
		Name: fake.FullName(),
		Age:  byte(randAge + 1),
		CreditCard: model.CreditCard{
			Number: fake.CreditCardNum(""),
		},
	}
}

func randomUserWithCompany() *model.User {
	user := randomUser()
	user.Company = model.Company{
		Name: fake.Company(),

	}
	return user
}
