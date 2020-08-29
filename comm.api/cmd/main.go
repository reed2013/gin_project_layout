package main

import (
	"comm.api/global"
	"comm.api/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	global.Init()
	gin.SetMode(global.ServerConf.RunMode)
	router := routers.NewRouter()

	server := http.Server{
		Addr:              ":" + global.ServerConf.HttpPort,
		Handler:           router,
		ReadTimeout:       global.ServerConf.ReadTimeout,
		WriteTimeout:      global.ServerConf.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}