package controller

import (
	"net/http"

	"github.com/Captainistz/lmwn-intern-2025/services"

	"github.com/gin-gonic/gin"
)

func GetCovidSummary(ctx *gin.Context) {
	covidCases, err := services.GetCases()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	covidSummary, err := services.GetCovidSummary(covidCases)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, covidSummary)
}
