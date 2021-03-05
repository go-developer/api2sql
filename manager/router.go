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
func Run(listenPort int) error {
	Router = &router{
		ginRouter:     gin.Default(),
		port:          listenPort,
		dbClientTable: make(map[uint64]define2.DatabaseInfo),
		apiTable:      make(map[string]*define2.APIInfo),
	}
	return Router.init()
}

type router struct {
	ginRouter     *gin.Engine
	port          int
	dbClientTable map[uint64]define2.DatabaseInfo
	apiTable      map[string]*define2.APIInfo
}

func (r *router) init() error {
	// 初始化路由与连接
	routerGroupTable := make(map[uint64]*gin.RouterGroup)
	for _, dbInstance := range Database.GetAllDBInstance() {
		if dbClient, err := r.getDatabaseClient(dbInstance); nil != err {
			return err
		} else {
			r.dbClientTable[dbInstance.ID] = define2.DatabaseInfo{
				Client: dbClient,
				Flag:   dbInstance.Flag,
				DbID:   dbInstance.ID,
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
		URI:          apiConfig.URI,
		FullURI:      strings.ReplaceAll("/"+dbInstance.Flag+"/"+apiConfig.URI, "//", "/"),
		SQL:          apiConfig.SQL,
		ApiID:        apiConfig.ID,
		DBInstanceID: dbInstance.DbID,
		RouterGroup:  dbInstance.Flag,
		CacheConfig:  define2.CacheConfig{},
	}
	return info, nil
}
