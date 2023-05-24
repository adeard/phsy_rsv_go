package handler

import (
	"fmt"
	"net/http"
	usertype "phsy_rsv_go/modules/user_type"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userTypeHandler struct {
	userTypeService usertype.Service
}

func NewUserTypeHandler(v1 *gin.RouterGroup, userTypeService usertype.Service) {
	handler := &userTypeHandler{userTypeService}

	userType := v1.Group("user-types")
	userType.GET("", handler.GetUserTypes)
	userType.POST("", handler.PostUserType)
	userType.GET(":ID", handler.GetUserType)
	userType.POST(":ID", handler.UpdateUserType)
	userType.DELETE(":ID", handler.DeleteUserType)
}

func (h *userTypeHandler) GetUserTypes(c *gin.Context) {
	usertypes, err := h.userTypeService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	var userTypesResponse []usertype.UserTypeResponse

	for _, b := range usertypes {
		userTypeResponse := ConvertToUserTypeResponse(b)

		userTypesResponse = append(userTypesResponse, userTypeResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userTypesResponse,
	})
}

func (h *userTypeHandler) GetUserType(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	b, err := h.userTypeService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	userTypeResponse := ConvertToUserTypeResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": userTypeResponse,
	})
}

func (h *userTypeHandler) PostUserType(c *gin.Context) {
	var userTypeInput usertype.UserTypeRequest

	err := c.ShouldBindJSON(&userTypeInput)
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

	usertype, err := h.userTypeService.Create(userTypeInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ConvertToUserTypeResponse(usertype),
	})
}

func (h *userTypeHandler) UpdateUserType(c *gin.Context) {
	var userTypeInput usertype.UserTypeRequest

	err := c.ShouldBindJSON(&userTypeInput)
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

	usertype, err := h.userTypeService.Update(id, userTypeInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ConvertToUserTypeResponse(usertype),
	})
}

func (h *userTypeHandler) DeleteUserType(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	result, err := h.userTypeService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"data": ConvertToUserTypeResponse(result),
	})
}

func ConvertToUserTypeResponse(b usertype.UserType) usertype.UserTypeResponse {
	return usertype.UserTypeResponse{
		ID:       int(b.ID),
		Name:     b.Name,
		IsActive: b.IsActive,
	}
}
