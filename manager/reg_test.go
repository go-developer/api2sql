// Package manager...
//
// Description : manager...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-06 3:07 下午
package manager

import (
	"fmt"
	"regexp"
	"testing"
)

func TestReg(t *testing.T) {
	r, err := regexp.Compile(`\?|\{.{1,100}?\}`)
	if nil != err {
		panic("正则编译失败 : " + err.Error())
	}
	sql := "select * from db_instance where id > {id} and name = ? and create_time > {create_time}"
	fmt.Println(r.MatchString(sql))
	fmt.Println(r.FindAllStringSubmatch(sql, -1))
}
