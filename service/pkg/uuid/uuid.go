package uuid

import (
    "database/sql/driver"
    "encoding/hex"
    "github.com/google/uuid"
)

// UUID 封装的google uuid
type UUID uuid.UUID

// New 生成UUIDv4
func New() UUID {
    return UUID(uuid.New())
}

// NewUUID1 生成UUIDv1
func NewUUID1() (u1 UUID, err error) {
    u, err := uuid.NewUUID()
    if err != nil {
        return
    }
    u1 = UUID(u)
    return
}

// NewOrderedUUID 生成有序UUID
func NewOrderedUUID() (ou UUID, err error) {
    ou, err = NewUUID1()
    if err != nil {
        return
    }

    // 重新排序uuid
    ou[4], ou[6] = ou[6], ou[4]
    ou[5], ou[7] = ou[7], ou[5]
    ou[0], ou[4] = ou[4], ou[0]
    ou[1], ou[5] = ou[5], ou[1]
    ou[2], ou[6] = ou[6], ou[2]
    ou[3], ou[7] = ou[7], ou[3]

    return
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

// Parse 将字符串解析为 UUID
func Parse(s string) (UUID, error) {
    id, err := uuid.Parse(s)
    return UUID(id), err
}

// String 返回uuid字符串
func (b UUID) String() string {
    return hex.EncodeToString(b[:])
}

// Format 返回uuid字符串, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (b UUID) Format() string {
    return uuid.UUID(b).String()
}

// GormDataType 数据库的数据类型
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
