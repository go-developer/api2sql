// Package driver...
//
// Description : 基于mysql存储配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-04 9:27 下午
package config

import (
	"github.com/go-developer/api2sql/driver/define"
	"github.com/go-developer/gopkg/middleware/mysql"
	"gorm.io/gorm"
)

// NewMysqlDriver mysql 管理配置的驱动
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:19 下午 2021/3/4
func NewMysqlDriver(conf *mysql.DBConfig, logConf *mysql.LogConfig) (Config, error) {
	cInstance := &Mysql{
		dbConf:  conf,
		logConf: logConf,
	}
	return cInstance, cInstance.Init()
}

// Mysql 数据库配置实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:18 下午 2021/3/4
type Mysql struct {
	dbConf  *mysql.DBConfig  // 数据库配置
	logConf *mysql.LogConfig // 日志配置
	client  *gorm.DB         // 数据连接实例
}

// Init 初始化数据库连接
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:38 下午 2021/3/4
func (m *Mysql) Init() error {
	var err error
	m.client, err = mysql.GetDatabaseClient(m.dbConf, m.logConf)
	return err
}

// LoadAllDatabaseConfig 获取全部数据库实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:39 下午 2021/3/4
func (m *Mysql) LoadAllDatabaseConfig() ([]define.DBInstance, error) {
	var (
		err            error
		dbInstanceList []define.DBInstance
	)
	if err = m.client.Where("status = ?", define.DBStatusUsing).Find(&dbInstanceList).Error; nil != err {
		return nil, err
	}
	return dbInstanceList, nil
}

// LoadAllAPIConfig 加载全部API配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:45 下午 2021/3/4
func (m *Mysql) LoadAllAPIConfig() ([]define.Api, error) {
	var (
		err     error
		apiList []define.Api
	)
	if err = m.client.Where("status = ?", define.APIStatusUsing).Find(&apiList).Error; nil != err {
		return nil, err
	}
	return apiList, nil
}

// LoadAllAPIParamConfig 加载全部api参数
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:46 下午 2021/3/4
func (m *Mysql) LoadAllAPIParamConfig() error {
	panic("implement me")
}
