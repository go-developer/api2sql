// Package manager...
//
// Description : manager...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-05 4:54 下午
package manager

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/go-developer/gopkg/logger"

	"github.com/go-developer/gopkg/middleware/mysql"

	define2 "github.com/go-developer/api2sql/define"
	"github.com/go-developer/api2sql/driver/define"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// Router ...
var Router *router

// Run 启动服务
func Run(ginRouter *gin.Engine, listenPort int) error {
	Router = &router{
		ginRouter:     ginRouter,
		port:          listenPort,
		dbClientTable: make(map[uint64]*define2.DatabaseInfo),
		apiTable:      make(map[string]*define2.APIInfo),
	}
	return Router.init()
}

type router struct {
	ginRouter     *gin.Engine
	port          int
	dbClientTable map[uint64]*define2.DatabaseInfo
	apiTable      map[string]*define2.APIInfo
}

func (r *router) init() error {
	// 初始化路由与连接
	routerGroupTable := make(map[uint64]*gin.RouterGroup)
	for _, dbInstance := range Database.GetAllDBInstance() {
		if dbClient, err := r.getDatabaseClient(dbInstance); nil != err {
			return err
		} else {
			r.dbClientTable[dbInstance.ID] = &define2.DatabaseInfo{
				Client:   dbClient,
				Flag:     dbInstance.Flag,
				DbID:     dbInstance.ID,
				ReadOnly: dbInstance.ReadOnly == 1,
			}
		}
		routerGroupTable[dbInstance.ID] = r.ginRouter.Group(dbInstance.Flag)
	}

	// 初始化API
	// 参数使用API ID进行分组
	paramGroup := make(map[uint64][]define.ApiParam)
	for _, p := range Param.GetAllParamRule() {
		if _, exist := paramGroup[p.ApiID]; !exist {
			paramGroup[p.ApiID] = make([]define.ApiParam, 0)
		}
		paramGroup[p.ApiID] = append(paramGroup[p.ApiID], p)
	}
	// API处理
	for _, uri := range API.GetAllAPIList() {
		paramList := make([]define.ApiParam, 0)
		if paramListInGroup, exist := paramGroup[uri.ID]; exist {
			paramList = paramListInGroup
		}
		apiInfo, err := r.buildApi(uri, paramList)
		if nil != err {
			return err
		}
		r.apiTable[apiInfo.FullURI] = apiInfo
		switch strings.ToUpper(apiInfo.Method) {
		case http.MethodGet:
			routerGroupTable[apiInfo.DBInstanceID].GET(apiInfo.URI, r.proxy)
		case http.MethodPost:
			routerGroupTable[apiInfo.DBInstanceID].POST(apiInfo.URI, r.proxy)
		case http.MethodDelete:
			routerGroupTable[apiInfo.DBInstanceID].DELETE(apiInfo.URI, r.proxy)
		case http.MethodHead:
			routerGroupTable[apiInfo.DBInstanceID].HEAD(apiInfo.URI, r.proxy)
		case http.MethodOptions:
			routerGroupTable[apiInfo.DBInstanceID].OPTIONS(apiInfo.URI, r.proxy)
		case http.MethodPatch:
			routerGroupTable[apiInfo.DBInstanceID].PATCH(apiInfo.URI, r.proxy)
		case http.MethodPut:
			routerGroupTable[apiInfo.DBInstanceID].PUT(apiInfo.URI, r.proxy)
		case "ANY": // 一次性注册全部请求方法的路由
			routerGroupTable[apiInfo.DBInstanceID].Any(apiInfo.URI, r.proxy)
		default:
			// 不是一个函数,数名method配置错误
			return fmt.Errorf("api_id=%d uri=%s method=%s 请求方法配置错误", apiInfo.ApiID, apiInfo.URI, apiInfo.Method)
		}
	}

	return r.ginRouter.Run(fmt.Sprintf(":%d", r.port))
}

// getDatabaseClient 获取数据库连接
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:56 下午 2021/3/5
func (r *router) getDatabaseClient(dbInstance define.DBInstance) (*gorm.DB, error) {
	dbConfig := &mysql.DBConfig{
		Host:              dbInstance.Host,
		Port:              dbInstance.Port,
		Database:          dbInstance.Database,
		Username:          dbInstance.Username,
		Password:          dbInstance.Password,
		Charset:           dbInstance.DbCharset,
		MaxOpenConnection: dbInstance.MaxConnection,
		MaxIdleConnection: dbInstance.MaxIdleConnection,
	}
	logConf := &mysql.LogConfig{
		Level:            0,
		ConsoleOutput:    true,
		Encoder:          nil,
		SplitConfig:      nil,
		ExtractFieldList: nil,
		TraceFieldName:   "trace_id",
	}
	var err error
	if logConf.SplitConfig, err = logger.NewRotateLogConfig("./logs", dbInstance.Flag+".log"); nil != err {
		return nil, err
	}

	return mysql.GetDatabaseClient(dbConfig, logConf)
}

// buildApi 构建API信息
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:33 下午 2021/3/5
func (r *router) buildApi(apiConfig define.Api, paramList []define.ApiParam) (*define2.APIInfo, error) {
	dbInstance, exist := r.dbClientTable[apiConfig.DbID]
	if !exist {
		return nil, errors.New("api关联的数据库配置不存在")
	}
	info := &define2.APIInfo{
		Method:       apiConfig.Method,
		URI:          apiConfig.URI,
		FullURI:      strings.ReplaceAll("/"+dbInstance.Flag+"/"+apiConfig.URI, "//", "/"),
		SQL:          apiConfig.SQL,
		ApiID:        apiConfig.ID,
		DBInstanceID: dbInstance.DbID,
		RouterGroup:  dbInstance.Flag,
		CacheConfig:  define2.CacheConfig{},
		ParamList:    paramList,
	}
	if err := Regexp.SQL(info); nil != err {
		return nil, err
	}
	return info, nil
}

// proxy 统一请求处理
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:53 下午 2021/3/5
func (r *router) proxy(ctx *gin.Context) {
	var (
		err       error
		result    []map[string]interface{}
		apiConfig *define2.APIInfo
		exist     bool
	)
	if apiConfig, exist = r.apiTable[ctx.Request.URL.Path]; !exist {
		ctx.JSON(http.StatusOK, gin.H{
			"code":     404,
			"message":  "api not found",
			"data":     nil,
			"trace_id": "",
		})
		return
	}

	if result, err = Runtime.Execute(ctx, r.dbClientTable[apiConfig.DBInstanceID], apiConfig); nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code":     500,
			"message":  "sql 执行失败 : " + err.Error(),
			"data":     nil,
			"trace_id": "",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "请求成功",
		"data": map[string]interface{}{
			"list": result,
		},
		"trace_id": "",
	})
}
