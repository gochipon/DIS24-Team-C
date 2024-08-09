package schema

import "time"

// CommentResponse represents the minimal Comment data required for GitHub-like UI
type CommentResponse struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"` // Extracted from "user"
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// IssueResponse represents the minimal Issue data required for GitHub-like UI, including comments
type IssueResponse struct {
	ID          int64             `json:"id"`
	Number      int64             `json:"number"`
	Title       string            `json:"title"`
	State       string            `json:"state"`
	Locked      bool              `json:"locked"`
	Author      string            `json:"author"`    // Extracted from "user"
	Assignees   []string          `json:"assignees"` // Extracted from "assignees"
	Labels      []string          `json:"labels"`    // Extracted from "labels"
	Comments    int64             `json:"comments"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	ClosedAt    *time.Time        `json:"closed_at,omitempty"`
	Milestone   string            `json:"milestone,omitempty"` // Extracted from "milestone"
	Repository  string            `json:"repository"`
	Body        string            `json:"body"`
	CommentList []CommentResponse `json:"comments_list"` // List of associated comments
}
