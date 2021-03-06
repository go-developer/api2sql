// Package manager...
//
// Description : manager...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-06 9:36 下午
package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/go-developer/api2sql/define"
	"github.com/pkg/errors"
)

// Runtime 初始化runtime
var Runtime *run

// InitRun 初始化运行时逻辑服务
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:36 下午 2021/3/6
func InitRun() {
	Runtime = &run{}
}

type run struct {
}

// CheckParam 运行时校验入参, TODO : 参数类型校验
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:43 下午 2021/3/6
func (r *run) CheckParam(ctx *gin.Context, apiInfo *define.APIInfo) ([]interface{}, error) {
	requestParam := make([]interface{}, 0)
	// 是否存在的校验
	for _, param := range apiInfo.ParamList {
		var (
			paramVal interface{}
			exist    bool
		)
		if paramVal, exist = ctx.GetQuery(param.Name); !exist {
			if param.IsRequired == 1 {
				// 必传,但是没有传
				return nil, errors.New(param.Name + " 要求必传,但是没有传")
			}
			// 非必传,使用默认值
			paramVal = param.DefaultValue
		}
		requestParam = append(requestParam, paramVal)
	}
	return requestParam, nil
}

// Execute 执行sql查询 TODO : 数据库读写分离校验
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:56 下午 2021/3/6
func (r *run) Execute(ctx *gin.Context, database *define.DatabaseInfo, apiInfo *define.APIInfo) ([]map[string]interface{}, error) {
	var (
		requestParam []interface{}
		err          error
		result       []map[string]interface{}
	)
	if requestParam, err = r.CheckParam(ctx, apiInfo); nil != err {
		return nil, err
	}
	if err = database.Client.Raw(apiInfo.SQL, requestParam...).Scan(&result).Error; nil != err {
		return nil, err
	}
	return result, nil
}
