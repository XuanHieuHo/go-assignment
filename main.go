package main

import (
	"log"

	"github.com/XuanHieuHo/go-assignment/config"
	"github.com/XuanHieuHo/go-assignment/controllers"
	"github.com/XuanHieuHo/go-assignment/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DatabaseConnect()
	// set up controller
	registry := controllers.NewGormControllerRegistery(db)

	router := gin.Default()
	api := router.Group("/api")

	routers.UserRouter(api, registry.UserController)
	routers.FriendshipRouter(api, registry.FriendController)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
