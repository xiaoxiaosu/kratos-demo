package data

import (
	"context"
	"database/sql"
	"demo/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"time"
)

type Data struct {
	Db       *sql.DB
	RedisCli redis.Cmdable
}

var ProviderSet = wire.NewSet(NewData, NewDb, NewRedisCli)

// 创建数据库连接
func NewDb(conf *conf.Data) (*sql.DB, error) {
	db, err := sql.Open(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// 连接redis
func NewRedisCli(conf *conf.Data) (redis.Cmdable, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        conf.Redis.Addr,
		DialTimeout: time.Second * 2,
		PoolSize:    10,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewData(dbClient *sql.DB, redisCmd redis.Cmdable) (*Data, func(), error) {
	data := &Data{
		Db:       dbClient,
		RedisCli: redisCmd,
	}

	return data, func() {
		if err := data.Db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
