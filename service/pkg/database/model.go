package database

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID    uint32    `json:"id" gorm:"column:id; primaryKey"`
	Ctime time.Time `json:"ctime" gorm:"column:ctime"` // 状态最后一次更改
	Mtime time.Time `json:"-" gorm:"column:mtime"`     // 数据最后一次修改
}

type DeleteModel struct {
	IsDelete bool `json:"del" gorm:"column:del"` // 数据软删除
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
		modifyDeleteField := db.Statement.Schema.LookUpField("del")
		if modifyDeleteField != nil {
			_ = modifyDeleteField.Set(db.Statement.ReflectValue, true)
		}

		nowTime := time.Now()
		modifyTimeField := db.Statement.Schema.LookUpField("mtime")
		if modifyTimeField != nil {
			_ = modifyTimeField.Set(db.Statement.ReflectValue, nowTime)
		}
	}
}
