package routes

import (
	"orderdetails/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/api/totalrevenue", controller.TotalRevenue)
	r.POST("/api/totalrevenuebyproduct", controller.TotalRevenuebyproduct)
	r.POST("/api/totalrevenuebycategory", controller.TotalRevenueByCategory)
	r.GET("/api/refresh", controller.RefreshData)

}
