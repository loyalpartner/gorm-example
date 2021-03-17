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
# end
