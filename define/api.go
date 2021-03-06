// Package define...
//
// Description : api 的数据结构
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-05 5:49 下午
package define

import (
	"github.com/go-developer/api2sql/driver/define"
	"gorm.io/gorm"
)

// DatabaseInfo ...
//
// Author : zhangdeman001@ke.com<张德满>
//
// Date : 6:35 下午 2021/3/5
type DatabaseInfo struct {
	Client   *gorm.DB
	Flag     string
	DbID     uint64
	ReadOnly bool
}

// APIInfo API数据结构
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:52 下午 2021/3/5
type APIInfo struct {
	URI          string            `json:"uri"`            // 最终注册的uri(不包含分组前缀)
	FullURI      string            `json:"full_uri"`       // 完整URI,包括分组前缀
	SQL          string            `json:"sql"`            // 绑定的SQL语句
	ApiID        uint64            `json:"api_id"`         // api 的 id
	DBInstanceID uint64            `json:"db_instance_id"` // db 的ID
	RouterGroup  string            `json:"router_group"`   // 路由分组
	CacheConfig  CacheConfig       `json:"cache_config"`   // 缓存配置
	ParamList    []define.ApiParam `json:"param_list"`     // 参数列表
}

// CacheConfig 缓存的配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:52 下午 2021/3/5
type CacheConfig struct {
	Enable          bool   `json:"enable"`            // 缓存是否启用
	CacheInstanceID uint64 `json:"cache_instance_id"` // 缓存实例ID
}

// ParamConfig 参数的配置规则
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:53 下午 2021/3/5
type ParamConfig struct {
	Model     uint              `json:"model"`      // 参数绑定模式
	ParamList []define.ApiParam `json:"param_list"` // 参数规则: 0 - 参数名绑定 1 - 索引顺序绑定 2 - 混合模式
}

// SQLConfig 最终的sql配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:31 下午 2021/3/6
type SQLConfig struct {
	SQL        string   // 执行的sql语句
	ParamOrder []string // 排好序的参数顺序
}
