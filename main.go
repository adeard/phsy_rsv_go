package main

import (
	"phsy_rsv_go/config"
	"phsy_rsv_go/handler"
	"phsy_rsv_go/middlewares"
	"phsy_rsv_go/modules/book"
	"phsy_rsv_go/modules/city"
	"phsy_rsv_go/modules/province"
	"phsy_rsv_go/modules/rate"
	"phsy_rsv_go/modules/user"
	userlevel "phsy_rsv_go/modules/user_level"
	usertype "phsy_rsv_go/modules/user_type"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db := config.Connect()

	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("api/v1")
	handler.NewUserHandler(v1, user.UserRegistry(db))

	v1.Use(middlewares.JwtAuthMiddleware())

	handler.NewBookHandler(v1, book.BookRegistry(db))
	handler.NewRateHandler(v1, rate.RateRegistry(db))
	handler.NewCityHandler(v1, city.CityRegistry(db))
	handler.NewUserTypeHandler(v1, usertype.UserTypeRegistry(db))
	handler.NewProvinceHandler(v1, province.ProvinceRegistry(db))
	handler.NewUserLevelHandler(v1, userlevel.UserLevelRegistry(db))

	router.Run(":85")
}
