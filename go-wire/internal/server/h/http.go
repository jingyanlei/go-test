package h

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func New() (server *http.Server, cf func(), err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	server = &http.Server{
		Addr:         ":1111",
		WriteTimeout: time.Second * 3, //设置3秒的写超时
		Handler:      mux,
	}
	go func() {
		err = server.ListenAndServe()
	}()
	fmt.Println("start http server")
	cf = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		if err := server.Shutdown(ctx); err != nil {
			fmt.Println("httpSrv.Shutdown error(%v)", err)
		}
		cancel()
		fmt.Println("close http")
	}
	return
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(i)
	}
}
