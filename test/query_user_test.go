package test

import (
	"fmt"
	"testing"

	"github.com/loyalpartner/gorm-example/database"
	"github.com/loyalpartner/gorm-example/model"
)

//
// 检索单个对象
//
func TestQueryFirst(t *testing.T) {
	user := new(model.User)

	result := database.DB.First(user)
	if result.Error != nil {
		t.Fatalf("query first failed: %v", result)
	}

	t.Logf("user is %v", user.Name)
}

func TestQueryTake(t *testing.T) {
	user := &model.User{}

	result := database.DB.Take(user)
	if result.Error != nil {
		t.Fatalf("query first failed: %v", result)
	}

	t.Logf("user is %v", user.Name)
}

func TestQueryLast(t *testing.T) {
	var user model.User

	result := database.DB.Last(&user)
	if result.Error != nil {
		t.Fatalf("query first failed: %v", result)
	}

	t.Logf("user is %v", user)
	t.Logf("credit card is %v", user.CreditCard)
}

func TestQueryModelFirst(t *testing.T) {
	// user := new(model.User)
	data := map[string]interface{}{}
	result := database.DB.Model(&model.User{}).First(&data)

	if result.Error != nil {
		t.Fatalf("query first failed: %v", result)
	}

	t.Logf("user is %v", data)
}

func TestQueryTableFirst(t *testing.T) {
	data := map[string]interface{}{}
	result := database.DB.Table("users").First(&data)

	if result.Error != nil {
		t.Skipf("出错就对了data is %v", data)
	}
}

//
// 用主键检索
//
func TestQueryFirstById(t *testing.T) {
	user := new(model.User)
	result := database.DB.First(&user, 10)

	if result.Error != nil {
		t.Fatalf("query first failed: %v", result)
	}
	t.Logf("user is %v", user.Name)
}

func TestQueryFirstById2(t *testing.T) {
	user := new(model.User)

	result := database.DB.First(user, "id = ?", 1)

	if result.Error != nil {
		t.Fatalf("query first failed: %v", result)
	}
	t.Logf("user is %v", user.Name)
}

func TestQueryFindById(t *testing.T) {
	users := new([]model.User)
	result := database.DB.Find(users, []int{1, 2, 3})

	if result.Error != nil {
		t.Fatalf("query first failed: %v", result)
	}

	for idx, user := range *users {
		t.Logf("%d:%v", idx, user.Name)
	}
}

//
// 检索全部对象
//
func TestQueryFindAll(t *testing.T) {
	users := new([]model.User)
	result := database.DB.Find(users)

	if result.Error != nil {
		t.Fatalf("query all failed: %v", result)
	}

	t.Logf("all user count is %v", len(*users))
}

//
// 预先加载
//
func TestQueryPreload(t *testing.T) {
	users := []model.User{}

	result := database.DB.Joins("left join credit_cards on users.id = credit_cards.user_id").
		Where("credit_cards.user_id is not null").
		Find(&users)

	if result.Error != nil {
		t.Fatalf("preload failed: %v", result)
	}

	t.Logf("preload user count is %v", len(users))
}

///
/// Scan
///

// TestQueryScan 测试嵌套数据scan
func TestQueryScan(t *testing.T) {
	// users := []model.User{}
	user := &model.User{}

	result := database.DB.Raw("select * from users where id = 1").Scan(user)

	if result.Error != nil {
		t.Fatalf("scan failed: %v", result)
	}

	t.Logf("the first name is %v", user.Name)
}

func a(b []string) []string{
	fmt.Printf("a() address %p\n", b)
	return b
	// fmt.Printf("a() address %p", b)
}

func TestTTT(t *testing.T) {
	c := []string{}

	fmt.Printf("b() address %p\n", c)
	e := a(c)
    
	fmt.Printf("e() address %p\n", e)
}
