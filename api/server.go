package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router := gin.Default()
	router.Any("/*", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"url":    c.Request.URL,
			"method": c.Request.Method,
			"query":  c.Request.URL.Query(),
			"body":   c.Request.Body,
		})
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
