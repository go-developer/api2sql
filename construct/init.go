// Package construct...
//
// Description : construct...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-05 12:41 下午
package construct

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	ginRouter.Group("admin").GET("getInfo", func(context *gin.Context) {
		r := []byte(`{"code":0,"data":{"name":"admin","roles":["Home","Dashbord","Driver","Driver-index","Permission","PageUser","PageAdmin","Roles","Table","BaseTable","ComplexTable","Icons","Icons-index","Components","Sldie-yz","Upload","Carousel","Echarts","Sldie-chart","Dynamic-chart","Map-chart","Excel","Excel-out","Excel-in","Mutiheader-out","Error","Page404","Github","NavTest","Nav1","Nav2","Nav2-1","Nav2-2","Nav2-2-1","Nav2-2-2","*404"],"introduce":"七没几红必无再住会果须容备南什心受部走太广月层变给局由联从说强间业到那生四明说体与交维连义中反支存那。"},"_res":{"status":200}}`)
		var result map[string]interface{}
		json.Unmarshal(r, &result)
		context.JSON(http.StatusOK, result)
	})
	adminController := admin.NewDefaultAdminController()
	iController := reflect.ValueOf(adminController)
	methodCnt := iController.NumMethod()
	fmt.Println(methodCnt)
	for i := 0; i < methodCnt; i++ {
		resultList := iController.Method(i).Call(nil)
		method := resultList[0].String()
		uri := resultList[1].String()
		if len(method) == 0 || len(uri) == 0 {
			// 未配置,忽略
			continue
		}
		middlewareList := resultList[2].Interface().([]gin.HandlerFunc)
		if nil == middlewareList {
			middlewareList = make([]gin.HandlerFunc, 0)
		}
		handler := resultList[3].Interface().(gin.HandlerFunc)
		if err := util.RegisterRouter(ginRouter, method, uri, handler); nil != err {
			panic(err.Error())
		}
	}
}
