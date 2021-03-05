// Package construct...
//
// Description : construct...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-05 12:41 下午
package construct

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-developer/api2sql/manager"
	"github.com/go-developer/gopkg/middleware/mysql"
)

// Run 构造函数引导服务运行
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:41 下午 2021/3/5
func Run(dbConfig *mysql.DBConfig, logConf *mysql.LogConfig, listenPort int) error {
	if err := manager.InitConfig(dbConfig, logConf); nil != err {
		return err
	}
	// 初始化路由实例
	router := gin.Default()
	// 启动端口监听
	if err := router.Run(fmt.Sprintf(":%d", listenPort)); nil != err {
		return err
	}
	return nil
}
