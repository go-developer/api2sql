// Package define ...
//
// Description : 数据结构定义
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-04 9:33 下午
package define

import "time"

// DBInstance DB实例信息存储表
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:36 下午 2021/3/4
type DBInstance struct {
	ID                uint64    `json:"id" gorm:"column:id"`                                   // 数据库ID
	Flag              string    `json:"flag" gorm:"column:flag"`                               // 数据库标识(会作为URL前缀)
	Host              string    `json:"host" gorm:"column:host"`                               // 数据库的地址
	Port              uint      `json:"port" gorm:"column:port"`                               // 数据库的端口
	Status            uint      `json:"status" gorm:"column:status"`                           // 0 - 待上线 1- 生效中 2 - 已下线
	Database          string    `json:"database" gorm:"column:database"`                       // 使用的数据库名称
	Username          string    `json:"username" gorm:"column:username"`                       // 连接的账号
	Password          string    `json:"password" gorm:"column:password"`                       // 连接密码
	DbCharset         string    `json:"db_charset" gorm:"column:db_charset"`                   // 数据库编码
	ReadOnly          uint      `json:"read_only" gorm:"column:read_only"`                     // 是否只读数据库, 0 - 否 1 - 是
	MaxConnection     uint      `json:"max_connection" gorm:"column:max_connection"`           // 最大的连接数,默认 50
	MaxIdleConnection uint      `json:"max_idle_connection" gorm:"column:max_idle_connection"` // 最大空闲连接数，默认25
	Description       string    `json:"description" gorm:"column:description"`                 // 数据库描述
	CreateTime        time.Time `json:"create_time" gorm:"column:create_time"`                 // 创建时间
	ModifyTime        time.Time `json:"modify_time" gorm:"column:modify_time"`                 // 更新时间
}

// DBInstanceTableName 获取表名
func (m *DBInstance) TableName() string {
	return DBInstanceTableName
}

const (
	// DBStatusWaitPublish 待上线
	DBStatusWaitPublish = iota
	// DBStatusUsing 使用中
	DBStatusUsing
	// DBStatusOffline 已下线
	DBStatusOffline
)

// ======================= 以上为db_instance表的常量

// CacheInstance 缓存实例配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:43 下午 2021/3/4
type CacheInstance struct {
	ID                uint64    `json:"id" gorm:"column:id"`                                   // 缓存实例ID
	Host              string    `json:"host" gorm:"column:host"`                               // 数据库的地址
	Port              uint      `json:"port" gorm:"column:port"`                               // 数据库的端口
	Driver            string    `json:"driver" gorm:"column:driver"`                           // 数据库的端口
	Status            uint      `json:"status" gorm:"column:status"`                           // 0 - 待上线 1- 生效中 2 - 已下线
	Database          string    `json:"database" gorm:"column:database"`                       // 使用的数据库名称
	Username          string    `json:"username" gorm:"column:username"`                       // 连接的账号
	Password          string    `json:"password" gorm:"column:password"`                       // 连接密码
	MaxConnection     uint      `json:"max_connection" gorm:"column:max_connection"`           // 最大的连接数,默认 50
	MaxIDleConnection uint      `json:"max_idle_connection" gorm:"column:max_idle_connection"` // 最大空闲连接数，默认25
	CreateTime        time.Time `json:"create_time" gorm:"column:create_time"`                 // 创建时间
	ModifyTime        time.Time `json:"modify_time" gorm:"column:modify_time"`                 // 更新时间
}

// TableName 缓存实例表
func (m *CacheInstance) TableName() string {
	return CacheInstanceTableName
}

// Api 表数据结构
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:50 下午 2021/3/4
type Api struct {
	ID           uint64    `json:"id" gorm:"column:id"`                         // API ID
	DbID         uint64    `json:"db_id" gorm:"column:db_id"`                   // 数据库ID(db_instance 表主键id)
	URI          string    `json:"uri" gorm:"column:uri"`                       // 访问的URI
	SQL          string    `json:"sql" gorm:"column:sql"`                       // api 对应的sql语
	Status       uint      `json:"status" gorm:"column:status"`                 // api状态
	Timeout      uint      `json:"timeout" gorm:"column:timeout"`               // 超时时间,单位ms,操作数据库的超时时间
	EnableCache  uint      `json:"enable_cache" gorm:"column:enable_cache"`     // 是否启用数据缓存 0 - 否 1 - 是
	CacheID      uint      `json:"cache_id" gorm:"column:cache_id"`             // 使用的缓存ID
	CacheTime    uint      `json:"cache_time" gorm:"column:cache_time"`         // 缓存有效期,单位:s,默认半小时
	ParamBindMod uint      `json:"param_bind_mod" gorm:"column:param_bind_mod"` // 0 - 参数名绑定 1 - 索引顺序绑定 2 - 混合模式
	Description  string    `json:"description" gorm:"column:description"`       // API描述
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`       // 创建时间
	ModifyTime   time.Time `json:"modify_time" gorm:"column:modify_time"`       // 更新时间
}

// TableName api表名
func (m *Api) TableName() string {
	return APITableName
}

const (
	// APIStatusWaitPublish 待上线
	APIStatusWaitPublish = iota
	// APIStatusUsing 使用中
	APIStatusUsing
	// APIStatusOffline 已下线
	APIStatusOffline
)

// ============================= 以上为api表枚举定义

// ApiParam api参数信息表
type ApiParam struct {
	ID           uint64    `json:"id" gorm:"column:id"`                       // Param ID
	ApiID        uint64    `json:"api_id" gorm:"column:api_id"`               // Param ID
	Name         string    `json:"name" gorm:"column:name"`                   // 参数名
	DataType     string    `json:"data_type" gorm:"column:data_type"`         // 参数类型
	Sort         uint      `json:"sort" gorm:column:sort`                     // 参数的排序
	DefaultValue string    `json:"default_value" gorm:"column:default_value"` // 默认值
	IsRequired   uint      `json:"is_required" gorm:"column:is_required"`     // 是否必传 0 - 否 1 - 是
	Description  string    `json:"description" gorm:"column:description"`     // 参数描述
	Example      string    `json:"example" gorm:"column:example"`             // 示例值
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`     // 创建时间
	ModifyTime   time.Time `json:"modify_time" gorm:"column:modify_time"`     // 更新时间
}

// TableName 参数配置表
func (m *ApiParam) TableName() string {
	return "api_param"
}
