package data

import (
	"service/pkg/database"
	"service/pkg/uuid"
)

// Account 账户模型
type Account struct {
	database.Model
	UUID     uuid.UUID `json:"uuid" gorm:"column:uuid"`
	TelCode  uint16    `json:"tel_code" gorm:"column:tel_code"`
	Phone    string    `json:"phone" gorm:"column:phone; type:char(11)"`
	Email    string    `json:"email" gorm:"column:email; type:char(46)"`
	Status   uint8     `json:"status" gorm:"column:status"`
	Password string    `json:"password" gorm:"column:password; type:char(40)"`
}
