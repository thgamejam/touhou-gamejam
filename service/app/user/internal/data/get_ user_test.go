package data

import (
	"context"
	"github.com/stretchr/testify/assert"
	"service/pkg/cache"
	"service/pkg/conf"
	"service/pkg/database"
	"testing"
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

func TestUserRepo_GetUserByAccountID(t *testing.T) {
	Cache, _ := cache.NewCache(Conf)
	DataBase, _ := database.NewDataBase(Conf)
	data, _, _ := NewData(nil, DataBase, Cache)
	assert.NotNil(t, data)
	ctx, _ := context.WithCancel(context.Background())

	repo := NewUserRepo(data, nil)
	id, err := repo.GetUserByAccountID(ctx, 2)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	t.Log(id)
}
