package main

import (
	"phsy_rsv_go/config"
	"phsy_rsv_go/middlewares"
	"phsy_rsv_go/modules/book"
	"phsy_rsv_go/modules/city"
	"phsy_rsv_go/modules/province"
	"phsy_rsv_go/modules/rate"
	"phsy_rsv_go/modules/user"
	"phsy_rsv_go/modules/userlevel"
	"phsy_rsv_go/modules/usertype"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db := config.Connect()

	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("api/v1")
	user.NewUserHandler(v1, user.UserRegistry(db))

	v1.Use(middlewares.JwtAuthMiddleware())

	book.NewBookHandler(v1, book.BookRegistry(db))
	rate.NewRateHandler(v1, rate.RateRegistry(db))
	city.NewCityHandler(v1, city.CityRegistry(db))
	usertype.NewUserTypeHandler(v1, usertype.UserTypeRegistry(db))
	province.NewProvinceHandler(v1, province.ProvinceRegistry(db))
	userlevel.NewUserLevelHandler(v1, userlevel.UserLevelRegistry(db))

	router.Run(":85")
}
