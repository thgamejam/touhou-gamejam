package database

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

    err = registerCallbacks(db)
    if err != nil {
        return nil, err
    }

    return db, nil
}

func registerCallbacks(db *gorm.DB) (err error) {
    err = db.Callback().Create().Before("gorm:create").Replace("gorm:create_time_stamp", createCallback)
    if err != nil {
        return err
    }
    err = db.Callback().Update().Before("gorm:update").Replace("gorm:update_time_stamp", updateCallback)
    if err != nil {
        return err
    }
    err = db.Callback().Delete().Before("gorm:delete").Replace("gorm:delete_time_stamp", deleteCallback)
    if err != nil {
        return err
    }
    return err
}
