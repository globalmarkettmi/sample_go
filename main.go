package main

import (
	shopapp "hoainam/gin-test/shop_app"
	userapp "hoainam/gin-test/user_app"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Perform authentication logic here
		// For example, check if the user is authenticated or if the provided token is valid

		// If authentication succeeds, proceed to the next middleware or route handler
		c.Next()

		// If authentication fails, you can return an error response or perform any desired action
		// For example:
		// c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	}
}
func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "Welcome server 02",
			},
		)
	})

	return e
}

var (
	g errgroup.Group
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
	v1.Use(AuthMiddleware())
	shopapp.GetRoutes(v1)
	//r.GET("/getproduct", papi.ShowProduct)
	// r.Run(":8000") // listen and serve on 0.0.0.0:8080
	vuser := r.Group("/user")
	userapp.GetRoutes(vuser)

	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
