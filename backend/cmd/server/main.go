package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/handler"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	e := gin.Default()
	godotenv.Load()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
		return
	}
	h := handler.NewInfoHandler(db)
	// /:org/:repo/issues/:number
	e.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
	})
	e.GET("/api/v1/:org/:repo/issue/:number", h.Issue())
	e.GET("/api/v1/:org/:repo/pull/:number", h.PullRequest())
	e.NoRoute(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.JSON(200, "Options Request!")
		}
	})
	if err := e.Run(":8080"); err != nil {
		return
	}
}
