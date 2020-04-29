package dao

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/google/wire"
)

// wire 定义Dao需要的依赖
var Provider = wire.NewSet(New, NewDB, NewRedis)

type Dao interface {
	Ping() (err error)
	Close()
}

type dao struct {
	db    *sql.DB
	redis *redis.Client
}

func New(db *sql.DB, redis *redis.Client) (d Dao, cf func(), err error) {
	d = &dao{
		db:    db,
		redis: redis,
	}
	cf = d.Close
	return
}

func (d *dao) Ping() (err error) {
	fmt.Println("ping...")
	return
}

func (d *dao) Close() {
	fmt.Println("dao close...")
	return
}
