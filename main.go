package main

import (
	"fmt"
	"net/http"
	"url_shortner/handler"
	"url_shortner/store"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Url-shortner")

	router := gin.Default()
	store.InnitializeStore()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the url-shortner",
		})
	})

	router.POST("/create_short_url", func(context *gin.Context) {
		handler.CreateShortUrl(context)
	})

	router.GET("/:short_url", func(context *gin.Context) {
		redirectURL := handler.HandelShortURLRedirection(context)
		fmt.Println("Redirecting to : ", redirectURL)
	})

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent) // Respond with no content
	})

	err := router.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Faild to start web server %v", err))
	}

}
