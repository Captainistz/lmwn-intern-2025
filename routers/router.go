package routers

import (
	"github.com/Captainistz/lmwn-intern-2025/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/")
	covid := r.Group("/covid")
	{
		covid.GET("/summary", controller.GetCovidSummary)
	}
}
