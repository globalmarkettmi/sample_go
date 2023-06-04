package userapp

import (
	userapi "hoainam/gin-test/user_app/apis"

	"github.com/gin-gonic/gin"
)

func GetRoutes(routeGroup *gin.RouterGroup) {

	routeGroup.POST("/login", userapi.Login)
}
