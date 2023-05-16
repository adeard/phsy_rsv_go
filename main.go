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

	bookRegistry := registry.BookRegistry(db)
	bookHandler := handler.NewBookHandler(bookRegistry)
	userRegistry := registry.UserRegistry(db)
	userHandler := handler.NewUserHandler(userRegistry)
	userTypeRegistry := registry.UserTypeRegistry(db)
	userTypeHandler := handler.NewUserTypeHandler(userTypeRegistry)
	userLevelRegistry := registry.UserLevelRegistry(db)
	userLevelHandler := handler.NewUserLevelHandler(userLevelRegistry)

	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("api/v1")
	v1.POST("/login", userHandler.Login)
	v1.POST("/register", userHandler.PostUser)

	v1.Use(middlewares.JwtAuthMiddleware())

	user := v1.Group("users")
	user.GET("", userHandler.GetUsers)
	user.GET(":ID", userHandler.GetUser)
	user.POST(":ID", userHandler.UpdateUser)
	user.DELETE(":ID", userHandler.DeleteUser)
	user.GET("logged", userHandler.CurrentUser)

	book := v1.Group("books")
	book.GET("", bookHandler.GetBooks)
	book.POST("", bookHandler.PostBook)
	book.GET(":ID", bookHandler.GetBook)
	book.PUT(":ID", bookHandler.UpdateBook)
	book.DELETE(":ID", bookHandler.DeleteBook)

	userType := v1.Group("user-types")
	userType.GET("", userTypeHandler.GetUserTypes)
	userType.POST("", userTypeHandler.PostUserType)
	userType.GET(":ID", userTypeHandler.GetUserType)
	userType.POST(":ID", userTypeHandler.UpdateUserType)
	userType.DELETE(":ID", userTypeHandler.DeleteUserType)

	userLevel := v1.Group("user-levels")
	userLevel.GET("", userLevelHandler.GetUserLevels)
	userLevel.POST("", userLevelHandler.PostUserLevel)
	userLevel.GET(":ID", userLevelHandler.GetUserLevel)
	userLevel.POST(":ID", userLevelHandler.UpdateUserLevel)
	userLevel.DELETE(":ID", userLevelHandler.DeleteUserLevel)

	router.Run(":85")
}
