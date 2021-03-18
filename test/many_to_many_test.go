package test

import (
	"fmt"
	"testing"

	"github.com/loyalpartner/gorm-example/database"
	"github.com/loyalpartner/gorm-example/model"
)

// var DB *gorm.DB

func init() {
	db := database.DB

	user := &model.User{}
	db.Where("name like ?", "many2many").First(&user)
	fmt.Printf("name is %v \n ", user.Name)
	if user.Name == "" {
		user.Name = "many2many"
		user.Roles = []model.Role{
			{Name: "小学生", Model: model.Model{}},
			{Name: "班长", Model: model.Model{}},
		}
		db.Create(user)
	}
}

// 查询用户以及对应的权限
func TestQueryMany2Many(t *testing.T) {
	db := database.DB

	// 预先加载查询
	user := model.User{}
	err := db.Preload("Roles").Where("name like ?", "many2many").First(&user).Error
	if err != nil {
		panic("ok")
	}

	t.Logf("user %v's rolse", user.Name)
	for i, role := range user.Roles {
		t.Logf("%v:%v", i, role.Name)
	}

	// 关联查询
	err = db.Model(&user).Association("Roles").Find(&user.Roles)
	if err != nil {
		t.Error("query roles error")
	}
	for i, role := range user.Roles {
		t.Logf("%v:%v", i, role.Name)
	}
}
