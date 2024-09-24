package handler

import (
	"log"
	"net/http"
	shortner "url_shortner/shortener"
	"url_shortner/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"` // `` is a struct tag that is used to define metadata for the struct field.

	UserId string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindBodyWithJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
         

        log.Printf("Request received for short url creation {long_url : %s,user_id:%s}: ", creationRequest.LongUrl, creationRequest.UserId) 


	shortURL := shortner.GenerateShortURLLink(creationRequest.LongUrl, creationRequest.UserId)

	store.SaveURLMapping(shortURL, creationRequest.LongUrl, creationRequest.UserId)

	c.JSON(http.StatusOK, gin.H{
		"mesage":    "short url created sucessfully",
		"short_url": c.Request.Host + "/" + shortURL,
                "long_url":  creationRequest.LongUrl, 
	})

}

func HandelShortURLRedirection(c *gin.Context) string {
	shortURL := c.Param("short_url")
	originalURL := store.RetriveOriginalUrl(shortURL)
	c.Redirect(http.StatusPermanentRedirect, originalURL)

	return originalURL

}
