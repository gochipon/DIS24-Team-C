package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/config"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/handler"
)

func main() {
	e := gin.Default()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.C.DBHost,
		config.C.DBPort,
		config.C.DBUser,
		config.C.DBPassword,
		config.C.DBName,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
		return
	}
	h := handler.NewInfoHandler(db)
	pc := handler.NewPineconeSearchHandler(db)
	// /:org/:repo/issues/:number
	e.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
	})
	e.GET("/api/v1/:org/:repo/issue/:number", h.Issue())
	e.GET("/api/v1/:org/:repo/pull/:number", h.PullRequest())
	e.POST("/api/v1/:org/:repo/search", pc.Search())
	e.NoRoute(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.JSON(200, "Options Request!")
		}
	})
	if err := e.Run(":8080"); err != nil {
		return
	}
}
