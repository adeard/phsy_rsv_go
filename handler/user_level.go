package handler

import (
	"fmt"
	"net/http"
	userlevel "phsy_rsv_go/modules/user_level"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userLevelHandler struct {
	userLevelService userlevel.Service
}

func NewUserLevelHandler(userLevelService userlevel.Service) *userLevelHandler {
	return &userLevelHandler{userLevelService}
}

func (h *userLevelHandler) GetUserLevels(c *gin.Context) {
	userlevels, err := h.userLevelService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	var userLevelsResponse []userlevel.UserLevelResponse

	for _, b := range userlevels {
		userLevelResponse := ConvertToUserLevelResponse(b)

		userLevelsResponse = append(userLevelsResponse, userLevelResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userLevelsResponse,
	})
}

func (h *userLevelHandler) GetUserLevel(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	b, err := h.userLevelService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	userLevelResponse := ConvertToUserLevelResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": userLevelResponse,
	})
}

func (h *userLevelHandler) PostUserLevel(c *gin.Context) {
	var userLevelInput userlevel.UserLevelRequest

	err := c.ShouldBindJSON(&userLevelInput)
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

	userlevel, err := h.userLevelService.Create(userLevelInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ConvertToUserLevelResponse(userlevel),
	})
}

func (h *userLevelHandler) UpdateUserLevel(c *gin.Context) {
	var userLevelInput userlevel.UserLevelRequest

	err := c.ShouldBindJSON(&userLevelInput)
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

	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	userlevel, err := h.userLevelService.Update(id, userLevelInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ConvertToUserLevelResponse(userlevel),
	})
}

func (h *userLevelHandler) DeleteUserLevel(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	result, err := h.userLevelService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"data": ConvertToUserLevelResponse(result),
	})
}

func ConvertToUserLevelResponse(b userlevel.UserLevel) userlevel.UserLevelResponse {
	return userlevel.UserLevelResponse{
		ID:       int(b.ID),
		Name:     b.Name,
		IsActive: b.IsActive,
	}
}
