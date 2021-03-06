// Package example...
//
// Description : example...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-06 1:40 下午
package main

import (
	"fmt"

	"github.com/go-developer/api2sql/construct"
	"github.com/go-developer/gopkg/logger"
	"github.com/go-developer/gopkg/middleware/mysql"
)

func main() {
	var err error
	dbConfig := &mysql.DBConfig{
		Host:              "127.0.0.1",
		Port:              3306,
		Database:          "api2sql",
		Username:          "root",
		Password:          "zhangdeman",
		Charset:           "utf8mb4",
		MaxOpenConnection: 20,
		MaxIdleConnection: 10,
	}
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
	construct.Run(dbConfig, logConf, 19808)

}
