// Package manager...
//
// Description : 管理所有的api接口
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-05 2:51 下午
package manager

import "github.com/go-developer/api2sql/driver/define"

// API ...
var API *api

// InitAPI 初始化api表数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:11 下午 2021/3/5
func InitAPI() error {
	API = &api{}
	return API.init()
}

type api struct {
	apiList []define.Api
}

// init 初始化
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:12 下午 2021/3/5
func (a *api) init() error {
	var err error
	a.apiList, err = Config.LoadAllAPIConfig()
	return err
}

// GetAllAPIList 读取全部的api列表
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:13 下午 2021/3/5
func (a *api) GetAllAPIList() []define.Api {
	return a.apiList
}
