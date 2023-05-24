package main

import (
	"phsy_rsv_go/config"
	"phsy_rsv_go/handler"
	"phsy_rsv_go/middlewares"
	"phsy_rsv_go/registry"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db := config.Connect()

	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("api/v1")
	handler.NewUserHandler(v1, registry.UserRegistry(db))

	v1.Use(middlewares.JwtAuthMiddleware())

	handler.NewBookHandler(v1, registry.BookRegistry(db))
	handler.NewRateHandler(v1, registry.RateRegistry(db))
	handler.NewUserTypeHandler(v1, registry.UserTypeRegistry(db))
	handler.NewProvinceHandler(v1, registry.ProvinceRegistry(db))
	handler.NewUserLevelHandler(v1, registry.UserLevelRegistry(db))

	router.Run(":85")
}
