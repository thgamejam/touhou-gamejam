package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"service/pkg/conf"
)

// NewDataBase 初始化数据库
func NewDataBase(c *conf.Service) (*gorm.DB, error) {
	db, err := gorm.Open(
		mysql.Open(c.Data.Database.Source),
		&gorm.Config{
			//Logger: , // TODO 绑定 Log 未完成
		})
	if err != nil {
		return nil, err
	}

	selDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	// 设置连接池
	// 空闲
	selDb.SetMaxIdleConns(int(c.Data.Database.MaxIdleConn))
	// 打开
	selDb.SetMaxOpenConns(int(c.Data.Database.MaxOpenConn))
	// 超时 time.Second * 30
	selDb.SetConnMaxLifetime(c.Data.Database.ConnMaxLifetime.AsDuration())

	return db, nil
}
