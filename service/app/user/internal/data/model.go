package data

import (
	"service/pkg/database"
	"service/pkg/uuid"
)

// Account 账户模型
type Account struct {
	database.Model
	Username    string    `json:"username" gorm:"column:username; type:varchar(16)"`
	TelCode     uint16    `json:"tel_code" gorm:"column:tel_code"`
	Phone       string    `json:"phone" gorm:"column:phone; type:char(11)"`
	Email       string    `json:"email" gorm:"column:email; type:char(46)"`
	Status      uint8     `json:"status" gorm:"column:status"`
	Description string    `json:"description" gorm:"column:description; type:varchar(64)"`
	UUID        uuid.UUID `json:"uuid" gorm:"column:uuid"`
	Password    string    `json:"password" gorm:"column:password; type:char(40)"`
	Avatar      uuid.UUID `json:"avatar" gorm:"column:avatar"`

	AvatarURL string `json:"avatar_url" gorm:"-"`
}
