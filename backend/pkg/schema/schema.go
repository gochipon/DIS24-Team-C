package schema

import "time"

type ReviewResponse struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"` // Extracted from "user"
	Body      string    `json:"body"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ReviewCommentResponse represents the minimal Review Comment data required for GitHub-like UI
type ReviewCommentResponse struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"` // Extracted from "user"
	Body      string    `json:"body"`
	Path      string    `json:"path"`
	Position  int64     `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PullRequestResponse represents the minimal Pull Request data required for GitHub-like UI, including reviews and review comments
type PullRequestResponse struct {
	ID                int64                   `json:"id"`
	Number            int64                   `json:"number"`
	Title             string                  `json:"title"`
	State             string                  `json:"state"`
	Locked            bool                    `json:"locked"`
	Draft             bool                    `json:"draft"`
	Author            string                  `json:"author"`    // Extracted from "user"
	Assignees         []string                `json:"assignees"` // Extracted from "assignees"
	Labels            []string                `json:"labels"`    // Extracted from "labels"
	CreatedAt         time.Time               `json:"created_at"`
	UpdatedAt         time.Time               `json:"updated_at"`
	ClosedAt          *time.Time              `json:"closed_at,omitempty"`
	MergedAt          *time.Time              `json:"merged_at,omitempty"`
	Milestone         string                  `json:"milestone,omitempty"` // Extracted from "milestone"
	Repository        string                  `json:"repository"`
	Body              string                  `json:"body"`
	MergeCommitSHA    string                  `json:"merge_commit_sha,omitempty"`
	HeadBranch        string                  `json:"head_branch"`         // Extracted from "head"
	BaseBranch        string                  `json:"base_branch"`         // Extracted from "base"
	ReviewList        []ReviewResponse        `json:"review_list"`         // List of associated reviews
	ReviewCommentList []ReviewCommentResponse `json:"review_comment_list"` // List of associated review comments
}
