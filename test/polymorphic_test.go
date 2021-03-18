package test

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/loyalpartner/gorm-example/database"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = database.DB
	// db.AutoMigrate(&Toy{})
	// db.AutoMigrate(&Cat{})
	// db.AutoMigrate(&Dog{})
}

type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

func TestCreate(t *testing.T) {
	// db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
	err := db.Create(&Dog{
		Name: fake.MaleFullName(),
		Toy:  Toy{Name: fake.FullName()},
	}).Error
	if err != nil {
		t.Errorf("insert failed")
	}
}

func TestDelete(t *testing.T) {
	dog := Dog{}
	err := db.Last(&dog).Error
	if err != nil {
		return
	}
	err = db.Delete(&dog).Error
	if err != nil {
		t.Errorf("insert failed")
	}
}

func TestQueryToys(t *testing.T) {
	dogs := []Dog{}
	// 关联查询
	err := db.Preload("Toy").Find(&dogs).Error
	if err != nil {
		t.Errorf("query dogs failed")
	}
	for _, dog := range dogs {
		t.Logf("dog:%v is %v", dog.Name, dog.Toy.Name)
	}
}
