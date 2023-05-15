package handler

import (
	"fmt"
	"net/http"
	"phsy_rsv_go/modules/user"
	"phsy_rsv_go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Login(c *gin.Context) {
	var loginInput user.LoginRequest

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
	return
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err.Error(),
		})
	}

	var result []user.UserResponse

	for _, user := range users {
		result = append(result, convertToUserResponse(user))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
	return
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
	return
}

func (h *userHandler) PostUser(c *gin.Context) {
	var userInput user.RegisterRequest

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
	return
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
	return
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var userInput user.UpdateRequest

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
	return
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
	return
}

func convertToUserResponse(b user.User) user.UserResponse {
	return user.UserResponse{
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
	}
}
