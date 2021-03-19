##
# gorm-example
#
# @file
# @version 0.1

run:
	go run ./main.go

test_add_user:
	go test -v -run TestAddUserWithCreditCard ./test/add_user_test.go
test_add_user_company:
	go test -v -run TestAddUserWithCompany ./test/add_user_test.go
test_query_user:
	go test -v -run TestQueryFirst ./test/query_user_test.go
test_query_where_condition:
	go test -v -run TestQueryWhere ./test/query_user_test.go
test_query_inline_condition:
	go test -v -run TestInlineCondition ./test/query_user_test.go
test_query_specify_struct:
	go test -v -run TestSpecifyStruct ./test/query_user_test.go
test_query_not_condition:
	go test -v -run TestNotCondition ./test/query_user_test.go
test_query_or_condition:
	go test -v -run TestOrCondition ./test/query_user_test.go
test_query_preload:
	go test -v -run TestQueryPreload ./test/query_user_test.go
# 测试association查询
test_query_association_company:
	go test -v -run TestQueryAssociation ./test/query_user_test.go
test_query_specific_fields:
	go test -v -run TestSpecificFields ./test/query_user_test.go
test_query_order_sort:
	go test -v -run TestOrderSort ./test/query_user_test.go
test_query_limit_and_offset:
	go test -v -run TestLimitAndOffset ./test/query_user_test.go
test_query_polymorphic:
	go test -v  ./test/polymorphic_test.go

test_query_many_2_many:
	go test -v -run TestQueryMany2Many ./test/many_to_many_test.go
test_user:
	go test -v ./model/user_test.go
test_company:
	go test -v ./model/company_test.go
# end
