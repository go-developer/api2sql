// Package manager...
//
// Description : manager...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-05 12:44 下午
package manager

import (
	"github.com/go-developer/api2sql/driver/config"
	"github.com/go-developer/gopkg/middleware/mysql"
)

// Config 配置管理实例
var Config config.Config

// InitConfig 初始化配置管实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:45 下午 2021/3/5
func InitConfig(dbConfig *mysql.DBConfig, logConfig *mysql.LogConfig) error {
	var err error
	if Config, err = config.NewMysqlDriver(dbConfig, logConfig); nil != err {
		return err
	}
	return nil
}
