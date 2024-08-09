package uc

import (
	"database/sql"
	"fmt"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/model"
)

type QueryIssueUC struct {
	DB *sql.DB
}

func NewQueryIssueUC(db *sql.DB) *QueryIssueUC {
	return &QueryIssueUC{
		DB: db,
	}
}

// Exec retrieves the IssuePage based on the provided issue ID
func (handler *QueryIssueUC) Exec(repository string, issueID int64) (*model.IssuePage, error) {
	// Initialize an empty IssuePage
	issuePage := &model.IssuePage{}

	// Retrieve the Issue data
	err := handler.DB.QueryRow(`
        SELECT id, url, body, "user", draft, state, title, labels, locked, number,
               node_id, user_id, assignee, comments, html_url, assignees, closed_at, 
               milestone, reactions, created_at, events_url, labels_url, repository, 
               updated_at, comments_url, pull_request, state_reason, timeline_url, 
               repository_url, active_lock_reason, author_association, performed_via_github_app,
               _airbyte_raw_id, _airbyte_extracted_at, _airbyte_generation_id, _airbyte_meta
        FROM issues
        WHERE number = $1 and repository = $2`, issueID, repository).Scan(
		&issuePage.Issue.ID,
		&issuePage.Issue.URL,
		&issuePage.Issue.Body,
		&issuePage.Issue.User,
		&issuePage.Issue.Draft,
		&issuePage.Issue.State,
		&issuePage.Issue.Title,
		&issuePage.Issue.Labels,
		&issuePage.Issue.Locked,
		&issuePage.Issue.Number,
		&issuePage.Issue.NodeID,
		&issuePage.Issue.UserID,
		&issuePage.Issue.Assignee,
		&issuePage.Issue.Comments,
		&issuePage.Issue.HTMLURL,
		&issuePage.Issue.Assignees,
		&issuePage.Issue.ClosedAt,
		&issuePage.Issue.Milestone,
		&issuePage.Issue.Reactions,
		&issuePage.Issue.CreatedAt,
		&issuePage.Issue.EventsURL,
		&issuePage.Issue.LabelsURL,
		&issuePage.Issue.Repository,
		&issuePage.Issue.UpdatedAt,
		&issuePage.Issue.CommentsURL,
		&issuePage.Issue.PullRequest,
		&issuePage.Issue.StateReason,
		&issuePage.Issue.TimelineURL,
		&issuePage.Issue.RepositoryURL,
		&issuePage.Issue.ActiveLockReason,
		&issuePage.Issue.AuthorAssociation,
		&issuePage.Issue.PerformedViaGitHubApp,
		&issuePage.Issue.AirbyteRawID,
		&issuePage.Issue.AirbyteExtractedAt,
		&issuePage.Issue.AirbyteGenerationID,
		&issuePage.Issue.AirbyteMeta,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve issue: %v", err)
	}

	// Retrieve the Comments data
	rows, err := handler.DB.Query(`
        SELECT id, url, body, "user", node_id, user_id, html_url, issue_url, reactions, 
               created_at, repository, updated_at, author_association, performed_via_github_app,
               _airbyte_raw_id, _airbyte_extracted_at, _airbyte_generation_id, _airbyte_meta
        FROM comments
        WHERE issue_url = $1 AND repository = $2`, issuePage.Issue.HTMLURL, repository)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve comments: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.URL,
			&comment.Body,
			&comment.User,
			&comment.NodeID,
			&comment.UserID,
			&comment.HTMLURL,
			&comment.IssueURL,
			&comment.Reactions,
			&comment.CreatedAt,
			&comment.Repository,
			&comment.UpdatedAt,
			&comment.AuthorAssociation,
			&comment.PerformedViaGitHubApp,
			&comment.AirbyteRawID,
			&comment.AirbyteExtractedAt,
			&comment.AirbyteGenerationID,
			&comment.AirbyteMeta,
		); err != nil {
			return nil, fmt.Errorf("failed to scan comment: %v", err)
		}
		issuePage.Comments = append(issuePage.Comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading comment rows: %v", err)
	}

	return issuePage, nil
}
