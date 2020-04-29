// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"go-wire/internal/dao"
	"go-wire/internal/server/h"
	"go-wire/internal/service"

	"github.com/google/wire"
)

// Initialize 声明injector的函数签名
func InitApp() (app *App, closeFunc func(), err error) {
	wire.Build(dao.Provider, service.Provider, h.New, NewApp)
	return
}

// // Initialize 声明injector的函数签名
// func InitRedis() (r *redis.Client, err error) {
// 	wire.Build(dao.ProviderRedis)
// 	return
// }

// // Initialize 声明injector的函数签名
// func InitDB() (r *sql.DB, err error) {
// 	wire.Build(dao.NewDB)
// 	return
// }

// // Initialize 声明injector的函数签名
// func InitDAO() (r dao.Dao, cf func(), err error) {
// 	wire.Build(dao.Provider)
// 	return
// }
