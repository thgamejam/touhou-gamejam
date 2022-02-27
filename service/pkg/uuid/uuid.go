package uuid

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

// UUID 封装的google uuid
type UUID uuid.UUID

func New() UUID {
	return UUID(uuid.New())
}

// Scan scan uuid
func (b *UUID) Scan(value interface{}) (err error) {
	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	if err != nil {
		return
	}
	*b = UUID(parseByte)
	return
}

// Value get uuid value
func (b UUID) Value() (driver.Value, error) {
	return uuid.UUID(b).MarshalBinary()
}

// ParseUUID 将字符串解析为 UUID
func ParseUUID(s string) (UUID, error) {
	id, err := uuid.Parse(s)
	return UUID(id), err
}

//String 返回uuid的字符串形式, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (b UUID) String() string {
	return uuid.UUID(b).String()
}

//GormDataType 设置数据库类型为binary(16)
func (b UUID) GormDataType() string {
	return "binary(16)"
}

func (b UUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(b)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (b *UUID) UnmarshalJSON(by []byte) (err error) {
	s, err := uuid.ParseBytes(by)
	if err != nil {
		return
	}
	*b = UUID(s)
	return
}
