package user

import (
	"fmt"
	"net/http"
	"phsy_rsv_go/domain"
	"phsy_rsv_go/middlewares"
	"phsy_rsv_go/utils"
	"strconv"

	"phsy_rsv_go/modules/userlevel"
	"phsy_rsv_go/modules/usertype"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService Service
}

func NewUserHandler(v1 *gin.RouterGroup, userService Service) {

	handler := &userHandler{userService}

	v1.POST("/login", handler.Login)
	v1.POST("/register", handler.PostUser)

	v1.Use(middlewares.JwtAuthMiddleware())

	user := v1.Group("users")
	user.GET("", handler.GetUsers)
	user.GET(":ID", handler.GetUser)
	user.POST(":ID", handler.UpdateUser)
	user.DELETE(":ID", handler.DeleteUser)
	user.GET("logged", handler.CurrentUser)
}

func (h *userHandler) Login(c *gin.Context) {
	var loginInput domain.LoginRequest

	c.ShouldBind(&loginInput)

	token, err := h.userService.Login(loginInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})

		return
	}

	var result []domain.UserResponse

	for _, user := range users {
		result = append(result, convertToUserResponse(user))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *userHandler) GetUser(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	u, err := h.userService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})

		return
	}

	userResponse := convertToUserResponse(u)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (h *userHandler) PostUser(c *gin.Context) {
	var userInput domain.RegisterRequest

	err := c.ShouldBindJSON(&userInput)
	if err != nil {

		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s , condition : %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": errorMessages,
		})

		return
	}

	user, err := h.userService.Create(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
		return
	}

	userResponse := convertToUserResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (h *userHandler) CurrentUser(c *gin.Context) {

	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.userService.FindByID(int(user_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse := convertToUserResponse(u)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": userResponse})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var userInput domain.UpdateRequest

	user_id := c.Param("ID")
	id, _ := strconv.Atoi(user_id)

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.userService.Update(id, userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse := convertToUserResponse(u)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": userResponse})
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	user_id := c.Param("ID")
	id, _ := strconv.Atoi(user_id)

	u, err := h.userService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse := convertToUserResponse(u)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": userResponse})
}

func convertToUserResponse(b domain.User) domain.UserResponse {
	userTypeResponse := usertype.ConvertToUserTypeResponse(b.UserType)
	userLevelResponse := userlevel.ConvertToUserLevelResponse(b.UserLevel)
	return domain.UserResponse{
		ID:          int(b.ID),
		Username:    b.Username,
		UserTypeId:  b.UserTypeId,
		UserLevelId: b.UserLevelId,
		Name:        b.Name,
		Email:       b.Email,
		ImgProfile:  b.ImgProfile,
		Address:     b.Address,
		Gender:      b.Gender,
		BirthDate:   b.BirthDate,
		IsActive:    b.IsActive,
		UserType:    userTypeResponse,
		UserLevel:   userLevelResponse,
	}
}
