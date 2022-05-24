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
	Password []byte    `json:"password" gorm:"column:password"`
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

// PrepareCreateEMailAccountCache 预创建用户缓存
type PrepareCreateEMailAccountCache struct {
	Email      string `json:"email"`
	KeyHash    string `json:"hash"`       // 秘钥哈希键
	Ciphertext string `json:"ciphertext"` // 密文密码
}
