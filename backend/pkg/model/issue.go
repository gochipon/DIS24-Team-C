package model

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Issue struct {
	ID                    int64           `json:"id"`
	URL                   string          `json:"url"`
	Body                  string          `json:"body"`
	User                  json.RawMessage `json:"user"`
	Draft                 sql.NullBool    `json:"draft"`
	State                 string          `json:"state"`
	Title                 string          `json:"title"`
	Labels                json.RawMessage `json:"labels"`
	Locked                bool            `json:"locked"`
	Number                int64           `json:"number"`
	NodeID                string          `json:"node_id"`
	UserID                sql.NullInt64   `json:"user_id"`
	Assignee              *string         `json:"assignee"`
	Comments              int64           `json:"comments"`
	HTMLURL               string          `json:"html_url"`
	Assignees             *string         `json:"assignees"`
	ClosedAt              *time.Time      `json:"closed_at"`
	Milestone             *string         `json:"milestone"`
	Reactions             json.RawMessage `json:"reactions"`
	CreatedAt             time.Time       `json:"created_at"`
	EventsURL             string          `json:"events_url"`
	LabelsURL             string          `json:"labels_url"`
	Repository            string          `json:"repository"`
	UpdatedAt             time.Time       `json:"updated_at"`
	CommentsURL           string          `json:"comments_url"`
	PullRequest           *string         `json:"pull_request"`
	StateReason           *string         `json:"state_reason"`
	TimelineURL           string          `json:"timeline_url"`
	RepositoryURL         string          `json:"repository_url"`
	ActiveLockReason      *string         `json:"active_lock_reason"`
	AuthorAssociation     string          `json:"author_association"`
	PerformedViaGitHubApp *string         `json:"performed_via_github_app"`
	AirbyteRawID          string          `json:"_airbyte_raw_id"`
	AirbyteExtractedAt    time.Time       `json:"_airbyte_extracted_at"`
	AirbyteGenerationID   int64           `json:"_airbyte_generation_id"`
	AirbyteMeta           json.RawMessage `json:"_airbyte_meta"`
}

// Comment represents a comment from the comments table
type Comment struct {
	ID                    int64           `json:"id"`
	URL                   string          `json:"url"`
	Body                  string          `json:"body"`
	User                  json.RawMessage `json:"user"`
	NodeID                string          `json:"node_id"`
	UserID                *int64          `json:"user_id"`
	HTMLURL               string          `json:"html_url"`
	IssueURL              string          `json:"issue_url"`
	Reactions             json.RawMessage `json:"reactions"`
	CreatedAt             time.Time       `json:"created_at"`
	Repository            string          `json:"repository"`
	UpdatedAt             time.Time       `json:"updated_at"`
	AuthorAssociation     string          `json:"author_association"`
	PerformedViaGitHubApp *string         `json:"performed_via_github_app"`
	AirbyteRawID          string          `json:"_airbyte_raw_id"`
	AirbyteExtractedAt    time.Time       `json:"_airbyte_extracted_at"`
	AirbyteGenerationID   int64           `json:"_airbyte_generation_id"`
	AirbyteMeta           json.RawMessage `json:"_airbyte_meta"`
}

// IssuePage represents the combined data of an issue and its comments
type IssuePage struct {
	Issue    Issue     `json:"issue"`
	Comments []Comment `json:"comments"`
}
