package controllers

import (
	"net/http"
	"strings"

	"funculator/models"
	"funculator/services"

	"github.com/gin-gonic/gin"
)

func Calculate(c *gin.Context) {
	var request models.CalculateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	expression := strings.TrimSpace(request.Expression)

	result, err := services.Calculate(expression)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, models.CalculateResponse{
		Expression: expression,
		Result:     result,
	})
}
