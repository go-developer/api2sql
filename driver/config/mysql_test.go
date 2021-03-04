// Package config...
//
// Description : config...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-04 10:54 下午
package config

import (
	"fmt"
	"testing"

	"github.com/go-developer/gopkg/logger"

	"github.com/go-developer/gopkg/middleware/mysql"
)

// TestMysqlLoadAllDBInstance 单测获取从mysql加载全部的可用db实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:55 下午 2021/3/4
func TestMysqlLoadAllDBInstance(t *testing.T) {
	var (
		err error
		c   Config
	)
	logConf := &mysql.LogConfig{
		Level:            0,
		ConsoleOutput:    true,
		Encoder:          logger.GetEncoder(),
		ExtractFieldList: nil,
		TraceFieldName:   "",
	}
	logConf.SplitConfig, err = logger.NewRotateLogConfig("/Users/zhangdeman/project/go-project/api2sql/logs", "test-api2sql.log")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c, err = NewMysqlDriver(&mysql.DBConfig{
		Host:              "127.0.0.1",
		Port:              3306,
		Database:          "api2sql",
		Username:          "root",
		Password:          "zhangdeman",
		Charset:           "utf8mb4",
		MaxOpenConnection: 20,
		MaxIdleConnection: 10,
	}, logConf)

	if nil != err {
		panic("数据库初始化失败 :" + err.Error())
	}

	fmt.Println(c.LoadAllDatabaseConfig())
}
