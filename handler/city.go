package handler

import (
	"fmt"
	"net/http"
	"phsy_rsv_go/modules/city"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type cityHandler struct {
	cityService city.Service
}

func NewCityHandler(v1 *gin.RouterGroup, cityService city.Service) {
	handler := &cityHandler{cityService}

	city := v1.Group("cities")
	city.GET("", handler.GetAll)
	city.POST("", handler.Post)
	city.GET(":ID", handler.GetDetail)
	city.POST(":ID", handler.Update)
	city.DELETE(":ID", handler.Delete)
}

func (h *cityHandler) GetAll(c *gin.Context) {
	cities, err := h.cityService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	var citiesResponse []city.CityResponse

	for _, b := range cities {
		cityResponse := ConvertToCityResponse(b)

		citiesResponse = append(citiesResponse, cityResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": citiesResponse,
	})
}

func (h *cityHandler) GetDetail(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	b, err := h.cityService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	cityResponse := ConvertToCityResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": cityResponse,
	})
}

func (h *cityHandler) Post(c *gin.Context) {
	var cityinput city.CityRequest

	err := c.ShouldBindJSON(&cityinput)
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

	city, err := h.cityService.Create(cityinput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ConvertToCityResponse(city),
	})
}

func (h *cityHandler) Update(c *gin.Context) {
	var cityInput city.CityRequest

	err := c.ShouldBindJSON(&cityInput)
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

	city, err := h.cityService.Update(id, cityInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ConvertToCityResponse(city),
	})
}

func (h *cityHandler) Delete(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	result, err := h.cityService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"data": ConvertToCityResponse(result),
	})
}

func ConvertToCityResponse(b city.City) city.CityResponse {
	return city.CityResponse{
		ID:         int(b.ID),
		Name:       b.Name,
		IsActive:   b.IsActive,
		ProvinceId: b.ProvinceId,
	}
}
