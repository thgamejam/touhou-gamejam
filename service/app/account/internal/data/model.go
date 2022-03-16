package data

import (
    "service/pkg/database"
    "service/pkg/uuid"
)

// Account 账户模型
type Account struct {
    database.Model
    TelCode  uint16    `json:"tel_code" gorm:"column:tel_code"`
    Phone    string    `json:"phone" gorm:"column:phone"`
    Email    string    `json:"email" gorm:"column:email"`
    Status   uint8     `json:"status" gorm:"column:status"`
    UUID     uuid.UUID `json:"uuid" gorm:"column:uuid"`
    Password string    `json:"password" gorm:"column:password"`
    UserID   uint64    `json:"user_id" gorm:"column:user_id"`
}

func (Account) TableName() string {
    return "account"
}

// LockOpener 公钥/密钥对
type LockOpener struct {
    ID      int    `json:"id"`
    Public  string `json:"public"`
    Private string `json:"private"`
}
