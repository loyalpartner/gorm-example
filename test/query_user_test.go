package test

import (
	"testing"
	"time"

	"github.com/loyalpartner/gorm-example/database"
	"github.com/loyalpartner/gorm-example/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
// 条件
//
func TestQueryWhere(t *testing.T) {

	user := new(model.User)
	users := new([]model.User)
	// 获取第一条匹配的记录
	if err := database.DB.Where("id = ?", 1).First(user).Error; err != nil {
		t.Errorf("id = %v not found", 1)
	}

	// 获取全部记录
	if err := database.DB.Where("id <> ?", 1).Find(users).Error; err != nil {
		t.Errorf("search all error")
	}

	// in
	if err := database.DB.Where("id in ?", []int{1, 2, 3}).Find(users).Error; err != nil {
		t.Errorf("search in [1,2,3] error")
	}

	// like
	if err := database.DB.Where("name like ?", "%a%").Find(users).Error; err != nil {
		t.Errorf("like query error")
	}

	// and
	if err := database.DB.Where("id = ? and birthday is not null", 1).Find(users).Error; err != nil {
		t.Errorf("like query error")
	}

	// time
	if err := database.DB.Where("updated_at < ?", time.Now()).Find(users).Error; err != nil {
		t.Errorf("time query")
	}

	// between
	if err := database.DB.Where("id between 1 and 10").Find(users).Error; err != nil {
		t.Errorf("between query")
	}

	// struct & Map
	// 注意 当使用结构作为条件查询时，GORM 只会查询非零值字段。
	// 这意味着如果您的字段值为 0、''、false 或其他 零值，该字段不会被用于构建查询条件
	if err := database.DB.Where(&model.User{
		Model: gorm.Model{
			ID: 1,
		},
	}).First(user).Error; err != nil {
		t.Errorf("struct query")
	}

	if err := database.DB.Where(map[string]interface{}{"id": 1}).Find(users).Error; err != nil {
		t.Errorf("map query error")
	}

	if err := database.DB.Where([]int64{20, 21, 22}).Find(users).Error; err != nil {
		t.Errorf("slice query error")
	}

}

//
// 内联条件
//
func TestInlineCondition(t *testing.T) {
	user := new(model.User)
	users := new([]model.User)

	if err := database.DB.First(user, "name like ?", "%a%").Error; err != nil {
		t.Errorf("inline first error")
	}

	if err := database.DB.Find(users, "name like ?", "%a%").Error; err != nil {
		t.Errorf("inline find error")
	}

	if err := database.DB.Find(users, "id <> 1 and birthday is not null").Error; err != nil {
		t.Errorf("inline find error")
	}

	if err := database.DB.Find(users, model.User{Model: gorm.Model{ID: 1}}).Error; err != nil {
		t.Errorf("inline find with struct")
	}

	if err := database.DB.Find(users, map[string]interface{}{"id": 1}).Error; err != nil {
		t.Errorf("inline find with map")
	}

}

//
// 指定结构查询
//
func TestSpecifyStruct(t *testing.T) {
	users := new([]model.User)
	if err := database.DB.Where(&model.User{Model: gorm.Model{ID: 1}, Name: ""}, "name").Find(users).Error; err != nil {
		t.Errorf("specify struct query error")
	}

	if err := database.DB.Where(&model.User{Age: byte(0)}, "age").Find(users).Error; err != nil {
		t.Errorf("specify struct query error")
	}
}

//
// Not 条件
//
func TestNotCondition(t *testing.T) {
	// user := new(model.User)
	users := new([]model.User)
	if err := database.DB.Not("id = ?", 2).Find(users).Error; err != nil {
		t.Errorf("not condition error ")
	}

	if err := database.DB.Not([]interface{}{1, 2, 3}).Find(users).Error; err != nil {
		t.Errorf("not in condition")
	}

	if err := database.DB.Not(model.User{Model: gorm.Model{ID: 2}}).Find(users).Error; err != nil {
		t.Errorf("not condition with struct")
	}

	if err := database.DB.Not(map[string]interface{}{"id": 2}).Find(users).Error; err != nil {
		t.Errorf("not condition with struct")
	}
}

//
// Or 条件
//
func TestOrCondition(t *testing.T) {
	// user := new(model.User)
	users := new([]model.User)
	if err := database.DB.Where("id = ?", 1).Or("id = ?", 2).Find(users).Error; err != nil {
		t.Errorf("where or error ")
	}

	if err := database.DB.Where("id = ?", 1).Or(model.User{Model: gorm.Model{ID: 2}}).Find(users).Error; err != nil {
		t.Errorf("whero or with struct")
	}

	if err := database.DB.Where("id = ?", 1).Or(map[string]interface{}{"id": 2}).Find(users).Error; err != nil {
		t.Errorf("whero or with struct")
	}
}

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

//
// 选择特定字段
//
func TestSpecificFields(t *testing.T) {
	users := new(model.User)
	if err := database.DB.Select("name", "age").Find(users).Error; err != nil {
		t.Errorf("specific fields error")
	}

	err := database.DB.Select([]string{"name", "age"}).Find(users).Error
	if err != nil {
		t.Errorf("specific fields with slice")
	}
}

//
// Order 排序
//
func TestOrderSort(t *testing.T) {
	users := new([]model.User)
	result := database.DB.Order("created_at").Find(users)
	if result.Error != nil {
		t.Errorf("order query failed")
	}

	err := database.DB.Order("age desc, created_at asc").Find(users).Error
	if err != nil {
		t.Errorf("order query failed")
	}

	database.DB.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(users)

	t.Logf("%v", len(*users))

}

//
// Limit & Offset
//
func TestLimitAndOffset(t *testing.T) {
	users := []model.User{}
	// creditCards := []model.CreditCard{}
	err := database.DB.Limit(3).Find(&users).Error
	if err != nil {
		t.Errorf("limit query failed")
	}

	// 通过-1 消除limit条件
	t.Logf("-1 消除 limit 条件 enable")
	err = database.DB.
		Limit(3).
		Find(&users).
		Limit(-1).Find(nil).Error
	if err != nil {
		t.Errorf("limit offset query failed")
	}
	t.Logf("-1 消除 limit 条件")

	err = database.DB.Limit(3).Offset(100).Find(&users).Error
	if err != nil {
		t.Errorf("limit offset query failed")
	}

	// err = database.DB.Offset(5).Find(&users).Error
	// if err != nil {
	// 	t.Logf("offset query failed")
	// }
}

//
// Group & Having
//
// TODO

//
// Distinct
//
// TODO

//
// Joins
//
// TODO

//
// Join 预先加载
//
// TODO
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

//
// Scan
//
// TODO
