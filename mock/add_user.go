package main

import (
	"time"

	"github.com/icrowley/fake"
	"github.com/loyalpartner/gorm-example/database"
	"github.com/loyalpartner/gorm-example/model"
)

const (
  MAX_COUNT = 100
)

func main() {
	// 插入100 个用户

	users := [MAX_COUNT]model.User{}

  // fake.SetLang("zh")
  for _, user := range users {
    now := time.Now()
    email := fake.EmailAddress()
    user.Name = fake.FullName()
    user.Birthday = &now
    user.Email = &email
  }
  // database.DB.Session(&gorm.Session{ CreateBatchSize: MAX_COUNT}).Debug().CreateInBatches(users[:], MAX_COUNT)
  database.DB.Debug().CreateInBatches(users[0:], MAX_COUNT)
}
