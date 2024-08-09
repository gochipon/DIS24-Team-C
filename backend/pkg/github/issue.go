package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

// Issue represents the structure of an issue in the database
type Issue struct {
	ID          int64           `json:"id"`
	URL         string          `json:"url"`
	Body        string          `json:"body"`
	User        json.RawMessage `json:"user"`
	Draft       bool            `json:"draft"`
	Title       string          `json:"title"`
	Labels      json.RawMessage `json:"labels"`
	Locked      bool            `json:"locked"`
	Number      int64           `json:"number"`
	Assignee    json.RawMessage `json:"assignee"`
	Assignees   json.RawMessage `json:"assignees"`
	ClosedAt    *time.Time      `json:"closed_at"`
	Reactions   json.RawMessage `json:"reactions"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	PullRequest json.RawMessage `json:"pull_request"`
}

func main() {
	// Database connection string
	// localhost:5432
	godotenv.Load()
	hostDB := os.Getenv("DB_HOST")
	portDB := os.Getenv("DB_PORT")
	userDB := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", hostDB, portDB, userDB, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Successfully connected!")
	// The issue number you want to retrieve
	issueNumber := 9132

	// Query to retrieve the issue by number
	query := `SELECT id, url, body, "user", draft, title, labels, locked, number, assignee, 
				assignees, closed_at, reactions, created_at, updated_at, pull_request 
				FROM public.issues WHERE number = $1`

	var issue Issue
	err = db.QueryRow(query, issueNumber).Scan(
		&issue.ID, &issue.URL, &issue.Body, &issue.User, &issue.Draft, &issue.Title, &issue.Labels,
		&issue.Locked, &issue.Number, &issue.Assignee, &issue.Assignees, &issue.ClosedAt,
		&issue.Reactions, &issue.CreatedAt, &issue.UpdatedAt, &issue.PullRequest,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the user, labels, assignee, assignees, and pull_request JSON fields to readable strings
	var user, labels, assignee, assignees, pullRequest map[string]interface{}
	json.Unmarshal(issue.User, &user)
	json.Unmarshal(issue.Labels, &labels)
	json.Unmarshal(issue.Assignee, &assignee)
	json.Unmarshal(issue.Assignees, &assignees)
	json.Unmarshal(issue.PullRequest, &pullRequest)

	// Format the issue data to output
	fmt.Printf("Issue Number: #%d\n", issue.Number)
	fmt.Printf("Title: %s\n", issue.Title)
	fmt.Printf("URL: %s\n", issue.URL)
	fmt.Printf("State: %s\n", getState(issue.Draft, issue.ClosedAt))
	fmt.Printf("Created at: %s\n", issue.CreatedAt.Format(time.RFC3339))
	fmt.Printf("Updated at: %s\n", issue.UpdatedAt.Format(time.RFC3339))
	if issue.ClosedAt != nil {
		fmt.Printf("Closed at: %s\n", issue.ClosedAt.Format(time.RFC3339))
	}
	fmt.Printf("User: %s\n", user["login"])
	fmt.Printf("Assignee: %v\n", getAssignee(assignee))
	fmt.Printf("Labels: %v\n", labels)
	fmt.Printf("Locked: %t\n", issue.Locked)
	fmt.Printf("Pull Request: %v\n", pullRequest)
	fmt.Printf("\nDescription:\n%s\n", issue.Body)
}

func getState(draft bool, closedAt *time.Time) string {
	if closedAt != nil {
		return "closed"
	}
	if draft {
		return "draft"
	}
	return "open"
}

func getAssignee(assignee map[string]interface{}) string {
	if assignee == nil {
		return "None"
	}
	return fmt.Sprintf("%v", assignee["login"])
}
