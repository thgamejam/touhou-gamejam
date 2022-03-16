package database

import (
    "gorm.io/gorm"
    "time"
)

type Model struct {
    ID    uint64    `json:"id" gorm:"column:id; primaryKey"`
    Ctime time.Time `json:"ctime" gorm:"column:ctime"` // 状态最后一次更改
    Mtime time.Time `json:"-" gorm:"column:mtime"`     // 数据最后一次修改
}

func createCallback(db *gorm.DB) {
    if db.Statement.Schema != nil {
        nowTime := time.Now()
        createTimeField := db.Statement.Schema.LookUpField("ctime")
        if createTimeField != nil {
            _ = createTimeField.Set(db.Statement.ReflectValue, nowTime)
        }
        modifyTimeField := db.Statement.Schema.LookUpField("mtime")
        if modifyTimeField != nil {
            _ = modifyTimeField.Set(db.Statement.ReflectValue, nowTime)
        }
    }
}

func updateCallback(db *gorm.DB) {
    if db.Statement.Schema != nil {
        nowTime := time.Now()
        modifyTimeField := db.Statement.Schema.LookUpField("mtime")
        if modifyTimeField != nil {
            _ = modifyTimeField.Set(db.Statement.ReflectValue, nowTime)
        }
    }
}

func deleteCallback(db *gorm.DB) {
    if db.Statement.Schema != nil {
        nowTime := time.Now()
        modifyTimeField := db.Statement.Schema.LookUpField("mtime")
        if modifyTimeField != nil {
            _ = modifyTimeField.Set(db.Statement.ReflectValue, nowTime)
        }
    }
}
