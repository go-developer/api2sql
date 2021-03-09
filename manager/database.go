// Package manager ...
//
// Description : 管理并操作配置的数据库实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-05 12:38 下午
package manager

import "github.com/go-developer/api2sql/driver/define"

// Database ...
var Database *database

// InitDatabase 初始化配置管理实例 TODO 当前直接使用mysql,后续升级成多驱动
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:59 下午 2021/3/5
func InitDatabase() error {
	Database = &database{
		instanceList: make([]define.DBInstance, 0),
	}
	return Database.init()
}

// database 数据库实例管理
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:00 下午 2021/3/5
type database struct {
	instanceList []define.DBInstance
}

// init 初始化数据库配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 1:59 下午 2021/3/5
func (d *database) init() error {
	var err error
	d.instanceList, err = Config.LoadAllDatabaseConfig()
	return err
}

// GetAllDBInstance 获取全部的数据库实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:05 下午 2021/3/5
func (d *database) GetAllDBInstance() []define.DBInstance {
	return d.instanceList
}

// CreateDatabaseInstance  创建数据库实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:47 下午 2021/3/9
func (d *database) CreateDatabaseInstance(data define.DBInstance) (uint64, error) {
	return Config.CreateDatabaseInstance(data)
}
