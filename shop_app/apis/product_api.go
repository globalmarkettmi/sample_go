package shopapi

import (
	shopschema "hoainam/gin-test/shop_app/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	info := shopschema.ProductInfo{Name: "hoai nam"}

	c.JSON(http.StatusOK, info)
}
