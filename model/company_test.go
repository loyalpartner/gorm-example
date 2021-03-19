package model_test

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/loyalpartner/gorm-example/database"
	"github.com/loyalpartner/gorm-example/model"
)

func TestCreateCompany(t *testing.T) {
	company := model.Company{
		Name: fake.FullName(),
		Users: []model.User{
			{Name: fake.FullName()},
			{Name: fake.FullName()},
		},
	}

	t.Logf("%v", company)
	if err := database.DB.Create(&company).Error; err != nil {
		t.Errorf("error: %v", err)
	}

	database.DB.Model(&company).Association("Users").Delete(company.Users)
	if err := database.DB.Delete(&company).Error; err != nil {
		t.Errorf("error: %v", err)
	}
}
