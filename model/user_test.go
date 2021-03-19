package model_test

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/loyalpartner/gorm-example/database"
	"github.com/loyalpartner/gorm-example/model"
)

func TestCreateUesr(t *testing.T) {
	user := model.User{
		Name:    fake.FullName(),
		Age:     24,
		Email:   fake.EmailAddress(),
		Company: model.Company{Name: fake.FullName()},
	}


	if err := database.DB.Create(&user).Error; err != nil {
		t.Errorf("error: %v", err)
	}

	database.DB.Model(&user).Update("name", "张三")

	database.DB.Model(&user).Association("Company").Clear()
	if err := database.DB.Delete(&user).Error; err != nil {
		t.Errorf("error: %v", err)
	}
}
