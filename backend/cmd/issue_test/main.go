package main

import (
	"database/sql"
	"fmt"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/uc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	issue := uc.NewQueryIssueUC(db)
	exec, err := issue.Exec("cli/cli", 925)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println((*exec).Issue.Body)
	fmt.Println((*exec).Issue.Title)
	fmt.Println((*exec).Issue.Comments)
}
