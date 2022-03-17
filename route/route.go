package route

import (
	"assignment_2/controller"

	"github.com/gin-gonic/gin"
)

func StartRoute() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controller.CreateItems)
	router.GET("/orders", controller.GetItems)
	router.DELETE("/orders/:orderID", controller.DeleteItems)
	router.PUT("/orders/:orderID", controller.UpdateItems)

	return router
}
