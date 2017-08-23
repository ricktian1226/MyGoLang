package main

import (
	//_ "mysql-master"
	"test"
)

const (
	TEST_ENUM_MYSQL = iota
	TEST_ENUM_CONTENT_REPLACE
	TEST_ENUM_JSON
)

func main() {

	enum := TEST_ENUM_JSON
	var entity test.TestEntityInterface
	switch enum {
	case TEST_ENUM_MYSQL: //mysql测试
		entity = test.NewTestMysql("mysql", "root:@tcp(127.0.0.1:3333)/test?charset=utf8")
	case TEST_ENUM_CONTENT_REPLACE: //文件内容替换
		entity = test.NewTestContentReplace("E:/MyGoLang/MyGoLang/MyGoLang/GoTest/src", "test.txt", "content1", "content2")
	case TEST_ENUM_JSON: //json文件操作
		entity = test.NewTestJson("E:/MyGoLang/MyGoLang/MyGoLang/GoTest/src")
	default:
		return
	}

	if entity != nil {
		entity.Do()
	}
}
