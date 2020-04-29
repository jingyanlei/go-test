package di

import (
	"fmt"
	"go-wire/internal/service"
	"net/http"
)

type App struct {
	svc        *service.Service
	httpServer *http.Server
}

// App
func NewApp(svc *service.Service, httpServer *http.Server) (app *App, closeFunc func(), err error) {
	app = &App{
		svc:        svc,
		httpServer: httpServer,
	}
	closeFunc = func() {
		// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		// if err := httpServer.Shutdown(ctx); err != nil {
		// 	fmt.Println("httpSrv.Shutdown error(%v)", err)
		// }
		// cancel()
		fmt.Println("close app")
	}
	return
}
