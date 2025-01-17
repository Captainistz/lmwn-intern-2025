package routers

import (
	"net/http"

	"github.com/Captainistz/lmwn-intern-2025/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "🚀 Ready to serve.",
		})
	})
	covid := r.Group("/covid")
	{
		covid.GET("/summary", controller.GetCovidSummary)
	}
}
