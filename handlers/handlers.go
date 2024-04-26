package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var UrlMap = make(map[string]string)

func URLshorter(c *gin.Context) {
	OrginalURL := c.PostForm("url")
	if OrginalURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there is no url"})
		return
	}
	shortKey := GenerateShortKey()

	shortenedUrl := fmt.Sprintf("http://localhost:7000/short/%s", shortKey)
	UrlMap[shortKey] = OrginalURL
	// c.JSON(http.StatusOK, gin.H{
	// 	"shortenedURL": shortenedUrl,
	// })
	c.HTML(http.StatusOK, "response.html", gin.H{
		"URL": shortenedUrl,
	})
}

func GenerateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6
	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func Redirect(c *gin.Context) {
	shortKey := c.Param("shortKey")
	OrginalURL, exists := UrlMap[shortKey]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "Url not Found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, OrginalURL)
}
