package handler

import (
	"fmt"
	"net/http"
	"phsy_rsv_go/modules/province"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type provinceHandler struct {
	provinceService province.Service
}

func NewProvinceHandler(v1 *gin.RouterGroup, provinceService province.Service) {
	handler := &provinceHandler{provinceService}

	province := v1.Group("provinces")
	province.GET("", handler.GetAll)
	province.POST("", handler.Post)
	province.GET(":ID", handler.GetDetail)
	province.POST(":ID", handler.Update)
	province.DELETE(":ID", handler.Delete)
}

func (h *provinceHandler) GetAll(c *gin.Context) {
	provinces, err := h.provinceService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	var provincesResponse []province.ProvinceResponse

	for _, b := range provinces {
		provinceResponse := ConvertToProvinceResponse(b)

		provincesResponse = append(provincesResponse, provinceResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": provincesResponse,
	})
}

func (h *provinceHandler) GetDetail(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	b, err := h.provinceService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	provinceResponse := ConvertToProvinceResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": provinceResponse,
	})
}

func (h *provinceHandler) Post(c *gin.Context) {
	var provinceInput province.ProvinceRequest

	err := c.ShouldBindJSON(&provinceInput)
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

	province, err := h.provinceService.Create(provinceInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ConvertToProvinceResponse(province),
	})
}

func (h *provinceHandler) Update(c *gin.Context) {
	var provinceInput province.ProvinceRequest

	err := c.ShouldBindJSON(&provinceInput)
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

	province, err := h.provinceService.Update(id, provinceInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ConvertToProvinceResponse(province),
	})
}

func (h *provinceHandler) Delete(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	result, err := h.provinceService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"data": ConvertToProvinceResponse(result),
	})
}

func ConvertToProvinceResponse(b province.Province) province.ProvinceResponse {
	return province.ProvinceResponse{
		ID:       int(b.ID),
		Name:     b.Name,
		IsActive: b.IsActive,
	}
}
