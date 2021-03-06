package main

import (
	"net/http"
	"time"

	"github.com/taniwhy/mochi-match-rest/config"
	"github.com/taniwhy/mochi-match-rest/infrastructure/dao"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/router"
)

func init() {
	config.InitConf()
}

func main() {
	conn := dao.NewDatabase()
	defer conn.Close()

	routers := router.InitRouter(conn)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routers,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
