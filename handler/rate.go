package handler

import (
	"fmt"
	"net/http"
	"phsy_rsv_go/modules/rate"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type rateHandler struct {
	rateService rate.Service
}

func NewRateHandler(v1 *gin.RouterGroup, rateService rate.Service) {
	handler := &rateHandler{rateService}

	rate := v1.Group("rates")
	rate.GET("", handler.GetRates)
	rate.POST("", handler.PostRate)
	rate.GET(":ID", handler.GetRate)
	rate.POST(":ID", handler.UpdateRate)
	rate.DELETE(":ID", handler.DeleteRate)
}

func (h *rateHandler) GetRates(c *gin.Context) {
	rates, err := h.rateService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	var ratesResponse []rate.RateResponse

	for _, b := range rates {
		rateResponse := convertToRateResponse(b)

		ratesResponse = append(ratesResponse, rateResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ratesResponse,
	})
}

func (h *rateHandler) GetRate(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	b, err := h.rateService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	rateResponse := convertToRateResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": rateResponse,
	})
}

func (h *rateHandler) PostRate(c *gin.Context) {
	var rateInput rate.RateRequest

	err := c.ShouldBindJSON(&rateInput)
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

	rate, err := h.rateService.Create(rateInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToRateResponse(rate),
	})
}

func (h *rateHandler) UpdateRate(c *gin.Context) {
	var rateInput rate.RateRequest

	err := c.ShouldBindJSON(&rateInput)
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

	rate, err := h.rateService.Update(id, rateInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToRateResponse(rate),
	})
}

func (h *rateHandler) DeleteRate(c *gin.Context) {
	idString := c.Param("ID")
	id, _ := strconv.Atoi(idString)

	result, err := h.rateService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"data": convertToRateResponse(result),
	})
}

func convertToRateResponse(b rate.Rate) rate.RateResponse {
	return rate.RateResponse{
		ID:     int(b.ID),
		UserId: b.UserId,
		Rates:  b.Rates,
	}
}
