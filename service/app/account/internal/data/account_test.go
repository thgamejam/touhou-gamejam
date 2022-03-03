package data

import (
    "service/pkg/cache"
    "service/pkg/conf"
    "service/pkg/database"
    "testing"
)

var (
    Conf = &conf.Service{
        Data: &conf.Data{
            Database: &conf.Data_Database{
                Source:          "",
                MaxIdleConn:     0,
                MaxOpenConn:     0,
                ConnMaxLifetime: nil,
            },
            Redis: &conf.Data_Redis{
                Network:      "",
                Addr:         "",
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
    
    
}
