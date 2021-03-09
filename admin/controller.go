// Package admin...
//
// Description : admin...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-09 4:38 下午
package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-developer/api2sql/driver/define"
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

func (d *defaultController) GetDatabaseInstanceList() (method string, uri string, middlewareList []gin.HandlerFunc, handler gin.HandlerFunc) {
	return http.MethodGet, "/admin/database/list", []gin.HandlerFunc{middleware.InitRequest()}, func(ctx *gin.Context) {
		util.Response(ctx, 0, "请求成功", gin.H{
			"list":  manager.Database.GetAllDBInstance(),
			"total": len(manager.Database.GetAllDBInstance()),
		})
	}
}

func (d *defaultController) GetDatabaseInstanceDetail() (method string, uri string, middlewareList []gin.HandlerFunc, handler gin.HandlerFunc) {
	return http.MethodGet, "/admin/database/detail", []gin.HandlerFunc{middleware.InitRequest()}, func(ctx *gin.Context) {
		dbID := ctx.DefaultQuery("db_id", "")
		if len(dbID) == 0 {
			util.Response(ctx, -1, "请求参数错误", gin.H{})
			return
		}
		var detail define.DBInstance
		for _, item := range manager.Database.GetAllDBInstance() {
			if fmt.Sprintf("%d", item.ID) == dbID {
				detail = item
				break
			}
		}
		if detail.ID == 0 {
			util.Response(ctx, -1, "数据库实例不存在", gin.H{})
			return
		}
		util.Response(ctx, 0, "请求成功", detail)
	}
}

func (d *defaultController) UpdateDatabaseInstance() (method string, uri string, middlewareList []gin.HandlerFunc, handler gin.HandlerFunc) {
	return
	// panic("implement me")
}
