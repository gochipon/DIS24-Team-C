package uc

import (
	"database/sql"
	"fmt"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/model"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// PullRequestHandler handles the retrieval of PullRequestPage data from the database
type PullRequestHandler struct {
	DB *sql.DB
}

// NewPullRequestHandler creates a new PullRequestHandler
func NewPullRequestHandler(db *sql.DB) *PullRequestHandler {
	return &PullRequestHandler{
		DB: db,
	}
}

// Exec retrieves the PullRequestPage based on the provided pull request ID
func (handler *PullRequestHandler) Exec(repository string, pullRequestID int64) (*model.PullRequestPage, error) {
	// Initialize an empty PullRequestPage
	pullRequestPage := &model.PullRequestPage{}

	// Retrieve the PullRequest data
	err := handler.DB.QueryRow(`
        SELECT id, url, base, body, head, "user", draft, state, title, _links, labels, locked, number,
               node_id, assignee, diff_url, html_url, assignees, closed_at, issue_url, merged_at, 
               milestone, patch_url, auto_merge, created_at, repository, updated_at, commits_url, 
               comments_url, statuses_url, requested_teams, merge_commit_sha, active_lock_reason, 
               author_association, review_comment_url, requested_reviewers, review_comments_url,
               _airbyte_raw_id, _airbyte_extracted_at, _airbyte_generation_id, _airbyte_meta
        FROM pull_requests
        WHERE number = $1 and repository = $2`,
		pullRequestID, repository).Scan(
		&pullRequestPage.PullRequest.ID,
		&pullRequestPage.PullRequest.URL,
		&pullRequestPage.PullRequest.Base,
		&pullRequestPage.PullRequest.Body,
		&pullRequestPage.PullRequest.Head,
		&pullRequestPage.PullRequest.User,
		&pullRequestPage.PullRequest.Draft,
		&pullRequestPage.PullRequest.State,
		&pullRequestPage.PullRequest.Title,
		&pullRequestPage.PullRequest.Links,
		&pullRequestPage.PullRequest.Labels,
		&pullRequestPage.PullRequest.Locked,
		&pullRequestPage.PullRequest.Number,
		&pullRequestPage.PullRequest.NodeID,
		&pullRequestPage.PullRequest.Assignee,
		&pullRequestPage.PullRequest.DiffURL,
		&pullRequestPage.PullRequest.HTMLURL,
		&pullRequestPage.PullRequest.Assignees,
		&pullRequestPage.PullRequest.ClosedAt,
		&pullRequestPage.PullRequest.IssueURL,
		&pullRequestPage.PullRequest.MergedAt,
		&pullRequestPage.PullRequest.Milestone,
		&pullRequestPage.PullRequest.PatchURL,
		&pullRequestPage.PullRequest.AutoMerge,
		&pullRequestPage.PullRequest.CreatedAt,
		&pullRequestPage.PullRequest.Repository,
		&pullRequestPage.PullRequest.UpdatedAt,
		&pullRequestPage.PullRequest.CommitsURL,
		&pullRequestPage.PullRequest.CommentsURL,
		&pullRequestPage.PullRequest.StatusesURL,
		&pullRequestPage.PullRequest.RequestedTeams,
		&pullRequestPage.PullRequest.MergeCommitSHA,
		&pullRequestPage.PullRequest.ActiveLockReason,
		&pullRequestPage.PullRequest.AuthorAssociation,
		&pullRequestPage.PullRequest.ReviewCommentURL,
		&pullRequestPage.PullRequest.RequestedReviewers,
		&pullRequestPage.PullRequest.ReviewCommentsURL,
		&pullRequestPage.PullRequest.AirbyteRawID,
		&pullRequestPage.PullRequest.AirbyteExtractedAt,
		&pullRequestPage.PullRequest.AirbyteGenerationID,
		&pullRequestPage.PullRequest.AirbyteMeta,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve pull request: %v", err)
	}

	// Retrieve the Reviews data
	reviewsRows, err := handler.DB.Query(`
        SELECT id, body, "user", state, _links, node_id, html_url, commit_id, created_at, repository, 
               updated_at, submitted_at, pull_request_url, author_association, 
               _airbyte_raw_id, _airbyte_extracted_at, _airbyte_generation_id, _airbyte_meta
        FROM reviews
        WHERE pull_request_url = $1`, pullRequestPage.PullRequest.HTMLURL)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve reviews: %v", err)
	}
	defer reviewsRows.Close()

	for reviewsRows.Next() {
		var review model.Review
		if err := reviewsRows.Scan(
			&review.ID,
			&review.Body,
			&review.User,
			&review.State,
			&review.Links,
			&review.NodeID,
			&review.HTMLURL,
			&review.CommitID,
			&review.CreatedAt,
			&review.Repository,
			&review.UpdatedAt,
			&review.SubmittedAt,
			&review.PullRequestURL,
			&review.AuthorAssociation,
			&review.AirbyteRawID,
			&review.AirbyteExtractedAt,
			&review.AirbyteGenerationID,
			&review.AirbyteMeta,
		); err != nil {
			return nil, fmt.Errorf("failed to scan review: %v", err)
		}
		pullRequestPage.Reviews = append(pullRequestPage.Reviews, review)
	}

	if err := reviewsRows.Err(); err != nil {
		return nil, fmt.Errorf("error reading review rows: %v", err)
	}

	// Retrieve the ReviewComments data
	commentsRows, err := handler.DB.Query(`
        SELECT id, url, body, line, path, side, "user", _links, node_id, html_url, position, commit_id, 
               diff_hunk, reactions, created_at, repository, start_line, start_side, updated_at, subject_type,
               original_line, in_reply_to_id, pull_request_url, original_position, author_association, 
               original_commit_id, original_start_line, pull_request_review_id, 
               _airbyte_raw_id, _airbyte_extracted_at, _airbyte_generation_id, _airbyte_meta
        FROM review_comments
        WHERE pull_request_url = $1`, pullRequestPage.PullRequest.HTMLURL)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve review comments: %v", err)
	}
	defer commentsRows.Close()

	for commentsRows.Next() {
		var comment model.ReviewComment
		if err := commentsRows.Scan(
			&comment.ID,
			&comment.URL,
			&comment.Body,
			&comment.Line,
			&comment.Path,
			&comment.Side,
			&comment.User,
			&comment.Links,
			&comment.NodeID,
			&comment.HTMLURL,
			&comment.Position,
			&comment.CommitID,
			&comment.DiffHunk,
			&comment.Reactions,
			&comment.CreatedAt,
			&comment.Repository,
			&comment.StartLine,
			&comment.StartSide,
			&comment.UpdatedAt,
			&comment.SubjectType,
			&comment.OriginalLine,
			&comment.InReplyToID,
			&comment.PullRequestURL,
			&comment.OriginalPosition,
			&comment.AuthorAssociation,
			&comment.OriginalCommitID,
			&comment.OriginalStartLine,
			&comment.PullRequestReviewID,
			&comment.AirbyteRawID,
			&comment.AirbyteExtractedAt,
			&comment.AirbyteGenerationID,
			&comment.AirbyteMeta,
		); err != nil {
			return nil, fmt.Errorf("failed to scan review comment: %v", err)
		}
		pullRequestPage.ReviewComments = append(pullRequestPage.ReviewComments, comment)
	}

	if err := commentsRows.Err(); err != nil {
		return nil, fmt.Errorf("error reading review comment rows: %v", err)
	}

	return pullRequestPage, nil
}
