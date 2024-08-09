package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/handler"
)

func main() {
	e := gin.Default()
	h := handler.NewInfoHandler()
	// /:org/:repo/issues/:number
	e.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
	})
	e.GET("/api/v1/:org/:repo/issue/:number", h.Issue())
	e.GET("/api/v1/:org/:repo/pull/:number", h.PullRequest())
	err := e.Run(":8080")
	if err != nil {
		return
	}
}
