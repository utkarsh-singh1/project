package handler

import (
	"net/http"
	"github.com/utkarsh-singh1/URLShortener/storage"
	"github.com/utkarsh-singh1/URLShortener/shortener"
	"github.com/gin-gonic/gin"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}


func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	storage.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host +shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := storage.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
