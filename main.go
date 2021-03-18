package main

import (
	"fmt"

	"github.com/loyalpartner/gorm-example/database"
	"github.com/loyalpartner/gorm-example/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func init(){
	db = database.DB
}

func main() {
	// dsn := "root:a@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Printf("%v", db)

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Author{})
	db.AutoMigrate(&model.Blog{})
	db.AutoMigrate(&model.CreditCard{})
	db.AutoMigrate(&model.Company{})
	db.AutoMigrate(&model.Role{})
}
