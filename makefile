##
# gorm-example
#
# @file
# @version 0.1

run:
	go run ./main.go

test_add_user:
	go test -v ./test/add_user_test.go
test_query_user:
	go test -v ./test/query_user_test.go
test_where_condition:
	go test -v -run TestQueryWhere ./test/query_user_test.go
test_inline_condition:
	go test -v -run TestInlineCondition ./test/query_user_test.go
test_specify_struct:
	go test -v -run TestSpecifyStruct ./test/query_user_test.go
test_not_condition:
	go test -v -run TestNotCondition ./test/query_user_test.go
test_or_condition:
	go test -v -run TestOrCondition ./test/query_user_test.go
test_specific_fields:
	go test -v -run TestSpecificFields ./test/query_user_test.go
test_order_sort:
	go test -v -run TestOrderSort ./test/query_user_test.go
test_limit_and_offset:
	go test -v -run TestLimitAndOffset ./test/query_user_test.go
# end
