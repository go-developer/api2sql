// Package admin...
//
// Description : admin...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-09 4:38 下午
package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-developer/api2sql/manager"
	"github.com/go-developer/gopkg/gin/middleware"
	"github.com/go-developer/gopkg/gin/util"
)

// NewDefaultAdminController 默认的管理API
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:16 下午 2021/3/9
func NewDefaultAdminController() IController {
	return &defaultController{}
}

// defaultController 管理API的默认实现,可自行实现
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:47 下午 2021/3/9
type defaultController struct {
}

func (d *defaultController) GetDatabaseInstanceList() (uri string, middlewareList []gin.HandlerFunc, handler func(ctx *gin.Context)) {
	return "/admin/database/list", []gin.HandlerFunc{middleware.InitRequest()}, func(ctx *gin.Context) {
		util.Response(ctx, 0, "请求成功", gin.H{
			"list":  manager.Database.GetAllDBInstance(),
			"total": len(manager.Database.GetAllDBInstance()),
		})
	}
}

func (d *defaultController) GetDatabaseInstanceDetail() (uri string, middlewareList []gin.HandlerFunc, handler func(ctx *gin.Context)) {
	panic("implement me")
}

func (d *defaultController) UpdateDatabaseInstance() (uri string, middlewareList []gin.HandlerFunc, handler func(ctx *gin.Context)) {
	panic("implement me")
}
