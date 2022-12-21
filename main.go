package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": c.Param("hi"),
		})
	})

	r.GET("/v1/karma/share/link/:hash", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hash":     c.Param("hash"),
			"redirect": "magic.store",
		})
	})
	r.Run()
}
