package pkg

import (
	"context"
	"database/sql"
	"demo/internal/conf"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Data struct {
	Db       *sql.DB
	RedisCli redis.Cmdable
}

// 创建数据库连接
func newDb(conf *conf.Data) (*sql.DB, error) {
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
func newRedisCli(conf *conf.Data) (redis.Cmdable, error) {
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

func NewData(conf *conf.Data) (*Data, error) {
	data := &Data{}
	if db, err := newDb(conf); err != nil {
		return data, err
	} else {
		data.Db = db
	}

	if redisCli, err := newRedisCli(conf); err != nil {
		return data, err
	} else {
		data.RedisCli = redisCli
	}

	return data, nil
}
