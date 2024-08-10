package model

import (
	"encoding/json"
	"time"
)

// PullRequest represents a pull request from the pull_requests table
type PullRequest struct {
	ID                  int64           `json:"id"`
	URL                 string          `json:"url"`
	Base                json.RawMessage `json:"base"`
	Body                string          `json:"body"`
	Head                json.RawMessage `json:"head"`
	User                json.RawMessage `json:"user"`
	Draft               bool            `json:"draft"`
	State               string          `json:"state"`
	Title               string          `json:"title"`
	Links               json.RawMessage `json:"_links"`
	Labels              json.RawMessage `json:"labels"`
	Locked              bool            `json:"locked"`
	Number              int64           `json:"number"`
	NodeID              string          `json:"node_id"`
	Assignee            *string         `json:"assignee"`
	DiffURL             string          `json:"diff_url"`
	HTMLURL             string          `json:"html_url"`
	Assignees           string          `json:"assignees"`
	ClosedAt            *time.Time      `json:"closed_at"`
	IssueURL            string          `json:"issue_url"`
	MergedAt            *time.Time      `json:"merged_at"`
	Milestone           *string         `json:"milestone"`
	PatchURL            string          `json:"patch_url"`
	AutoMerge           *string         `json:"auto_merge"`
	CreatedAt           time.Time       `json:"created_at"`
	Repository          string          `json:"repository"`
	UpdatedAt           time.Time       `json:"updated_at"`
	CommitsURL          string          `json:"commits_url"`
	CommentsURL         string          `json:"comments_url"`
	StatusesURL         string          `json:"statuses_url"`
	RequestedTeams      json.RawMessage `json:"requested_teams"`
	MergeCommitSHA      string          `json:"merge_commit_sha"`
	ActiveLockReason    *string         `json:"active_lock_reason"`
	AuthorAssociation   string          `json:"author_association"`
	ReviewCommentURL    string          `json:"review_comment_url"`
	RequestedReviewers  json.RawMessage `json:"requested_reviewers"`
	ReviewCommentsURL   string          `json:"review_comments_url"`
	AirbyteRawID        string          `json:"_airbyte_raw_id"`
	AirbyteExtractedAt  time.Time       `json:"_airbyte_extracted_at"`
	AirbyteGenerationID int64           `json:"_airbyte_generation_id"`
	AirbyteMeta         json.RawMessage `json:"_airbyte_meta"`
}

// Review represents a review from the reviews table
type Review struct {
	ID                  int64           `json:"id"`
	Body                string          `json:"body"`
	User                json.RawMessage `json:"user"`
	State               string          `json:"state"`
	Links               json.RawMessage `json:"_links"`
	NodeID              string          `json:"node_id"`
	HTMLURL             string          `json:"html_url"`
	CommitID            string          `json:"commit_id"`
	CreatedAt           time.Time       `json:"created_at"`
	Repository          string          `json:"repository"`
	UpdatedAt           time.Time       `json:"updated_at"`
	SubmittedAt         time.Time       `json:"submitted_at"`
	PullRequestURL      string          `json:"pull_request_url"`
	AuthorAssociation   string          `json:"author_association"`
	AirbyteRawID        string          `json:"_airbyte_raw_id"`
	AirbyteExtractedAt  time.Time       `json:"_airbyte_extracted_at"`
	AirbyteGenerationID int64           `json:"_airbyte_generation_id"`
	AirbyteMeta         json.RawMessage `json:"_airbyte_meta"`
}

// ReviewComment represents a review comment from the review_comments table
type ReviewComment struct {
	ID                  int64           `json:"id"`
	URL                 string          `json:"url"`
	Body                string          `json:"body"`
	Line                *int64          `json:"line"`
	Path                string          `json:"path"`
	Side                string          `json:"side"`
	User                json.RawMessage `json:"user"`
	Links               json.RawMessage `json:"_links"`
	NodeID              string          `json:"node_id"`
	HTMLURL             string          `json:"html_url"`
	Position            *int64          `json:"position"`
	CommitID            string          `json:"commit_id"`
	DiffHunk            string          `json:"diff_hunk"`
	Reactions           json.RawMessage `json:"reactions"`
	CreatedAt           time.Time       `json:"created_at"`
	Repository          string          `json:"repository"`
	StartLine           *int64          `json:"start_line"`
	StartSide           *string         `json:"start_side"`
	UpdatedAt           time.Time       `json:"updated_at"`
	SubjectType         string          `json:"subject_type"`
	OriginalLine        int64           `json:"original_line"`
	InReplyToID         *int64          `json:"in_reply_to_id"`
	PullRequestURL      string          `json:"pull_request_url"`
	OriginalPosition    int64           `json:"original_position"`
	AuthorAssociation   string          `json:"author_association"`
	OriginalCommitID    string          `json:"original_commit_id"`
	OriginalStartLine   *int64          `json:"original_start_line"`
	PullRequestReviewID int64           `json:"pull_request_review_id"`
	AirbyteRawID        string          `json:"_airbyte_raw_id"`
	AirbyteExtractedAt  time.Time       `json:"_airbyte_extracted_at"`
	AirbyteGenerationID int64           `json:"_airbyte_generation_id"`
	AirbyteMeta         json.RawMessage `json:"_airbyte_meta"`
}

// PullRequestPage represents the combined data of a pull request, its reviews, and review comments
type PullRequestPage struct {
	PullRequest    PullRequest     `json:"pull_request"`
	Reviews        []Review        `json:"reviews"`
	ReviewComments []ReviewComment `json:"review_comments"`
}
