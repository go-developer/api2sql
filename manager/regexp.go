// Package manager...
//
// Description : 服务启动时,预编译可能用到的正则
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-06 2:56 下午
package manager

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-developer/api2sql/define"
	define2 "github.com/go-developer/api2sql/driver/define"

	"github.com/pkg/errors"
)

// Regexp ...
var Regexp *reg

// InitRegexp 初始化正则
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:56 下午 2021/3/6
func InitRegexp() error {
	var err error
	Regexp = &reg{}
	// 记得要用懒惰匹配,别用贪婪匹配
	if Regexp.sqlParser, err = regexp.Compile(`\?|\{.{1,100}?\}`); nil != err {
		return errors.New("sql解析正则预编译失败, 失败原因 : " + err.Error())
	}
	return nil
}

type reg struct {
	sqlParser *regexp.Regexp
}

// SQL sql解析
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:21 下午 2021/3/6
func (r *reg) SQL(apiInfo *define.APIInfo) error {
	sqlTemplate := apiInfo.SQL
	matchResult := r.sqlParser.FindAllStringSubmatch(sqlTemplate, -1)
	bindList := make([]string, 0)
	for _, item := range matchResult {
		for _, matchCase := range item {
			sqlTemplate = strings.ReplaceAll(sqlTemplate, matchCase, "?")
		}
		bindList = append(bindList, item...)
	}
	if len(bindList) != len(apiInfo.ParamList) {
		return errors.New(fmt.Sprintf("sql参数绑定需要数量 : %d, 实际配置的接口需要参数数量 : %d", len(bindList), len(apiInfo.ParamList)))
	}
	bindParamTable := make(map[string]bool)
	for _, item := range bindList {
		if item != "?" {
			bindParamTable[item] = true
		}
	}
	paramTable := make(map[string]define2.ApiParam)
	for _, param := range apiInfo.ParamList {
		key := "{" + param.Name + "}"
		if _, exist := bindParamTable[key]; exist {
			paramTable[key] = param
			continue
		}
		paramTable[fmt.Sprintf("%d", param.Sort)] = param
	}
	sortIndex := 0
	realParamList := make([]define2.ApiParam, 0)
	for _, itemParam := range bindList {
		if _, exist := bindParamTable[itemParam]; exist {
			realParamList = append(realParamList, paramTable[itemParam])
			continue
		}
		realParamList = append(realParamList, paramTable[fmt.Sprintf("%d", sortIndex)])
		sortIndex++
	}
	apiInfo.ParamList = realParamList
	apiInfo.SQL = sqlTemplate // 格式化之后的sql
	return nil
}
