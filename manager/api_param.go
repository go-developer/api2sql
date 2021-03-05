// Package manager...
//
// Description : api参数管理
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-05 4:36 下午
package manager

import "github.com/go-developer/api2sql/driver/define"

// Param ...
var Param *param

// InitParam 初始化param
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:38 下午 2021/3/5
func InitParam() error {
	Param = &param{}
	return Param.init()
}

type param struct {
	paramList []define.ApiParam
}

func (p *param) init() error {
	var err error
	p.paramList, err = Config.LoadAllAPIParamConfig()
	return err
}

// GetAllParamRule 获取全部的参数规则
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:41 下午 2021/3/5
func (p *param) GetAllParamRule() []define.ApiParam {
	return p.paramList
}
