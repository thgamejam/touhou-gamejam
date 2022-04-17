package data

import (
	"context"
	"encoding/json"
	"service/pkg/cache"
	"service/pkg/conf"
	"service/pkg/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	Conf = &conf.Service{
		Data: &conf.Data{
			Database: &conf.Data_Database{
				Source:          "root:123456@tcp(127.0.0.1:3306)/touhou_gamejam?charset=utf8mb4&parseTime=True&loc=Local",
				MaxIdleConn:     0,
				MaxOpenConn:     0,
				ConnMaxLifetime: nil,
			},
			Redis: &conf.Data_Redis{
				Network:      "tcp",
				Addr:         "127.0.0.1:6379",
				Password:     "",
				ReadTimeout:  nil,
				WriteTimeout: nil,
			},
		},
	}
)

func TestAccountRepo_GetAccountByID(t *testing.T) {
	Cache, _ := cache.NewCache(Conf)
	DataBase, _ := database.NewDataBase(Conf)
	data, _, _ := NewData(DataBase, Cache, nil)
	assert.NotNil(t, data)
	ctx, _ := context.WithCancel(context.Background())

	_ = DataBase.AutoMigrate(&Account{})
	repo := NewAccountRepo(data, nil)
	user, err := repo.GetAccountByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	u, _ := json.Marshal(user)
	t.Logf("%v\n", string(u))
}
