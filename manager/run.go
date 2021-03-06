// Package manager...
//
// Description : manager...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-06 9:36 下午
package manager

import (
	"fmt"
	"strings"

	"github.com/go-developer/gopkg/convert"

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
			paramVal    interface{}
			exist       bool
			formatParam interface{}
			err         error
		)
		if paramVal, exist = ctx.GetQuery(param.Name); !exist {
			if param.IsRequired == 1 {
				// 必传,但是没有传
				return nil, errors.New(param.Name + " 要求必传,但是没有传")
			}
			// 非必传,使用默认值
			paramVal = param.DefaultValue
		}
		if formatParam, err = r.ParamTypeCheck(ctx, param.Name, paramVal, param.DataType); nil != err {
			return nil, err
		}
		requestParam = append(requestParam, formatParam)
	}
	return requestParam, nil
}

// ParamTypeCheck 校验参数类型
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:26 下午 2021/3/6
func (r *run) ParamTypeCheck(ctx *gin.Context, paramName string, val interface{}, expectType string) (interface{}, error) {
	var err error
	errInfo := fmt.Errorf("%s 参数类型期望是 %s, 传入的数据格式错误", paramName, expectType)
	switch strings.ToLower(expectType) {
	case "string": // 字符串
		var result string
		if err = convert.ConvertAssign(&result, val); nil != err {
			return nil, errInfo
		}
		return result, nil
	case "int":
		fallthrough
	case "in8":
		fallthrough
	case "int16":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		var result int64
		if err = convert.ConvertAssign(&result, val); nil != err {
			return nil, errInfo
		}
		return result, nil
	case "uint":
		fallthrough
	case "uint8":
		fallthrough
	case "uint16":
		fallthrough
	case "uint32":
		fallthrough
	case "uint64":
		var result uint64
		if err = convert.ConvertAssign(&result, val); nil != err {
			return nil, errInfo
		}
		return result, nil
	case "float32":
		fallthrough
	case "float64":
		var result float64
		if err = convert.ConvertAssign(&result, val); nil != err {
			return nil, errInfo
		}
		return result, nil
	case "[]string":
		return r.getStringSlice(ctx, paramName, val)
	case "[]int":
		fallthrough
	case "[]int8":
		fallthrough
	case "[]int16":
		fallthrough
	case "[]int32":
		fallthrough
	case "[]int64":
		var (
			result    []string
			intResult []int64
		)
		intResult = make([]int64, 0)
		if result, err = r.getStringSlice(ctx, paramName, val); nil != err {
			return nil, errInfo
		}
		for _, item := range result {
			var tmp int64
			if err = convert.ConvertAssign(&tmp, item); nil != err {
				return nil, errInfo
			}
			intResult = append(intResult, tmp)
		}
		return intResult, nil
	case "[]uint":
		fallthrough
	case "[]uint8":
		fallthrough
	case "[]uint16":
		fallthrough
	case "[]uint32":
		fallthrough
	case "[]uint64":
		var (
			result    []string
			intResult []uint64
		)
		intResult = make([]uint64, 0)
		if result, err = r.getStringSlice(ctx, paramName, val); nil != err {
			return nil, errInfo
		}
		for _, item := range result {
			var tmp uint64
			if err = convert.ConvertAssign(&tmp, item); nil != err {
				return nil, errInfo
			}
			intResult = append(intResult, tmp)
		}
		return intResult, nil
	default:
		return nil, errInfo
	}
}

// getStringSlice 解析list类型参数
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:55 下午 2021/3/6
func (r *run) getStringSlice(ctx *gin.Context, paramName string, val interface{}) ([]string, error) {
	var (
		tmpResult string
		err       error
	)
	if err = convert.ConvertAssign(&tmpResult, val); nil != err {
		return nil, fmt.Errorf("%s 参数类型期望是 都好分割的字符串, 传入的数据格式错误", paramName)
	}
	return strings.Split(tmpResult, ","), nil
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
	if nil == result {
		result = make([]map[string]interface{}, 0)
	}
	return result, nil
}
