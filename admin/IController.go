// Package admin ...
//
// Description : 定义controller接口,自行实现接口可自己实现对数据的管理
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-09 4:39 下午
package admin

import "github.com/gin-gonic/gin"

// IController controller接口定义
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:40 下午 2021/3/9
type IController interface {
	// GetDatabaseInstanceList 查询已注册实例的列表
	//
	// Author : go_developer@163.com<张德满>
	//
	// Date : 4:42 下午 2021/3/9
	GetDatabaseInstanceList() (method string, uri string, middlewareList []gin.HandlerFunc, handler gin.HandlerFunc)
	// GetDatabaseInstanceDetail 获取已注册数据库实例详情
	//
	// Author : go_developer@163.com<张德满>
	//
	// Date : 4:44 下午 2021/3/9
	GetDatabaseInstanceDetail() (method string, uri string, middlewareList []gin.HandlerFunc, handler gin.HandlerFunc)
	// UpdateDatabaseInstance 更新数据库实例的信息
	//
	// Author : go_developer@163.com<张德满>
	//
	// Date : 4:45 下午 2021/3/9
	UpdateDatabaseInstance() (method string, uri string, middlewareList []gin.HandlerFunc, handler gin.HandlerFunc)
}
