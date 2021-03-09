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
	"reflect"

	"github.com/go-developer/gopkg/gin/util"

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
	iController := reflect.ValueOf(adminController)
	methodCnt := iController.NumMethod()
	fmt.Println(methodCnt)
	for i := 0; i < methodCnt; i++ {
		resultList := iController.Method(i).Call(nil)
		method := resultList[0].String()
		uri := resultList[1].String()
		middlewareList := resultList[2].Interface().([]gin.HandlerFunc)
		if nil == middlewareList {
			middlewareList = make([]gin.HandlerFunc, 0)
		}
		handler := resultList[3].Interface().(gin.HandlerFunc)
		ginRouter.GET(uri, handler).Use(middlewareList...)
		if err := util.RegisterRouter(ginRouter, method, uri, handler); nil != err {
			panic(err.Error())
		}
	}
}
