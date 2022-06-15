package data

import "service/pkg/database"

type User struct {
	Id               uint32 `json:"id" gorm:"column:id"`
	Name             string `json:"name" gorm:"column:name"`
	AccountID        uint32 `json:"account_id" gorm:"column:account_id"`
	AvatarID         string `json:"avatar_id" gorm:"column:avatar_id"`
	TagString        string `json:"tags" gorm:"column:tag_string"`
	AllowSyndication bool   `json:"allow_syndication" gorm:"column:allow_syndication"`
	WorksCount       uint32 `json:"works_count" gorm:"column:works_count"`
	FansCount        uint32 `json:"fans_count" gorm:"column:fans_count"`
	database.Model
}

func (User) TableName() string {
	return "user"
}
