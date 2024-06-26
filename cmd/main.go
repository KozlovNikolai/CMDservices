package main

import (
	"net/http"

	"github.com/KozlovNikolai/CMDservices/internal/config"
	"github.com/KozlovNikolai/CMDservices/internal/server"
	"github.com/KozlovNikolai/CMDservices/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.MustLoad()

	store.InitDB(cfg.StoragePath)
	defer store.CloseDB()
	router := gin.Default()

	router.POST("/services", server.Create)
	router.GET("/services/:id", server.Get)
	router.DELETE("/services/:id", server.Delete)
	router.GET("/services/list", server.GetList)

	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimout,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
