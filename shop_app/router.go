package shopapp

import (
	shopapi "hoainam/gin-test/shop_app/apis"

	"github.com/gin-gonic/gin"
)

func GetRoutes(routeGroup *gin.RouterGroup) {

	routeGroup.GET("/getinfo", shopapi.GetProduct)
}
