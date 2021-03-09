// Package construct...
//
// Description : construct...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-05 12:41 下午
package construct

import (
	"github.com/gin-gonic/gin"
	"github.com/go-developer/api2sql/admin"
	"github.com/go-developer/api2sql/manager"
	"github.com/go-developer/gopkg/middleware/mysql"
)

// Run 构造函数引导服务运行
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:41 下午 2021/3/5
func Run(dbConfig *mysql.DBConfig, logConf *mysql.LogConfig, listenPort int) error {
	// 初始化配置管理
	if err := manager.InitConfig(dbConfig, logConf); nil != err {
		return err
	}
	// 初始化数据库实例
	if err := manager.InitDatabase(); nil != err {
		return err
	}
	// 初始化正则
	if err := manager.InitRegexp(); nil != err {
		return err
	}
	// 初始化可用api列表
	if err := manager.InitAPI(); nil != err {
		return err
	}
	// 初始化可用api 参数列表
	if err := manager.InitParam(); nil != err {
		return err
	}
	// 启动端口监听
	ginRouter := gin.Default()

	// 设置管理员使用的API
	SetAdminApi(ginRouter)

	if err := manager.Run(ginRouter, listenPort); nil != err {
		return err
	}
	return nil
}

// SetAdminApi 设置管理员操作相关的API
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:30 下午 2021/3/9
func SetAdminApi(ginRouter *gin.Engine) {
	adminController := admin.NewDefaultAdminController()
	getDatabaseURI, middlewareList, handler := adminController.GetDatabaseInstanceList()
	ginRouter.GET(getDatabaseURI, handler).Use(middlewareList...)
}
