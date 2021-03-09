// Package driver...
//
// Description : 定义配置加载的数据源的约束
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-04 9:25 下午
package config

import "github.com/go-developer/api2sql/driver/define"

// Config 配置的接口
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:26 下午 2021/3/4
type Config interface {
	// Init 初始化
	Init() error
	// LoadAllDatabaseConfig 加载全部的数据库配置
	LoadAllDatabaseConfig() ([]define.DBInstance, error)
	// LoadAllAPIConfig 加载全部的API配置
	LoadAllAPIConfig() ([]define.Api, error)
	// LoadAllAPIParamConfig 加载全部的API参数
	LoadAllAPIParamConfig() ([]define.ApiParam, error)
	// CreateDatabaseInstance 创建数据库实例
	CreateDatabaseInstance(data define.DBInstance) (uint64, error)
}
