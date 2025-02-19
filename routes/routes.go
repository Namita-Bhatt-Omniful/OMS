package routes

import (
	"OMS/controllers"

	"github.com/omniful/go_commons/http"
)

// func Test(srvr *http.Server) {
// 	response := gin.H{"test": "Hello World!"}
// 	srvr.GET("/test", func(ctx *gin.Context) {
// 		ctx.JSON(200, response)
// 	})
// }

func GetRouter(srvr *http.Server) {
	orders := srvr.Group("/orders")
	{
		// orders.GET("/view", controllers.ViewOrders)
		orders.POST("/createBulk", controllers.CreateBulkOrder)
	}
}
