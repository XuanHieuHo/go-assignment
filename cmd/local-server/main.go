package main

import (
	"log"

	"github.com/XuanHieuHo/go-assignment/config"
	"github.com/XuanHieuHo/go-assignment/controllers"
	"github.com/XuanHieuHo/go-assignment/routers"
	"github.com/XuanHieuHo/go-assignment/uow"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	db := config.DatabaseConnect()
	// set up controller
	uow := uow.NewUnitOfWorkImpl(db)
	registry := controllers.NewGormControllerRegistery(uow)

	router := gin.Default()
	api := router.Group("/api")

	routers.UserRouter(api, registry.UserController)
	routers.FriendshipRouter(api, registry.FriendController)
	err = router.Run(c.LocalServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
