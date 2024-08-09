package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/model"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/schema"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/uc"
	"net/http"
	"strconv"
)

type InfoHandler struct {
	issueHandler *uc.IssueHandler
	pullHandler  *uc.PullRequestHandler
}

func NewInfoHandler(db *sql.DB) *InfoHandler {
	return &InfoHandler{
		issueHandler: uc.NewIssueHandler(db),
		pullHandler:  uc.NewPullRequestHandler(db),
	}
}

func (h *InfoHandler) Issue() gin.HandlerFunc {
	return func(c *gin.Context) {
		org := c.Param("org")
		repo := c.Param("repo")
		number := c.Param("number")
		parseInt, err := strconv.
			ParseInt(number, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		repository := org + "/" + repo
		issue, err := h.issueHandler.Exec(repository, parseInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, issue)
	}
}

func (h *InfoHandler) PullRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		org := c.Param("org")
		repo := c.Param("repo")
		number := c.Param("number")
		parseInt, err := strconv.
			ParseInt(number, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		repository := org + "/" + repo
		pull, err := h.pullHandler.Exec(repository, parseInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, pull)
	}
}

func convertModelToSchema(issue model.Issue, comments []model.Comment) (schema.IssueResponse, error) {
	var user map[string]interface{}
	var labels []string
	var assignees []string

	// Unmarshal JSON fields
	if err := json.Unmarshal(issue.User, &user); err != nil {
		return schema.IssueResponse{}, err
	}
	if err := json.Unmarshal(issue.Labels, &labels); err != nil {
		return schema.IssueResponse{}, err
	}
	if issue.Assignees != nil {
		if err := json.Unmarshal([]byte(*issue.Assignees), &assignees); err != nil {
			return schema.IssueResponse{}, err
		}
	}

	// Convert comments
	var commentList []schema.CommentResponse
	for _, comment := range comments {
		var commentUser map[string]interface{}
		if err := json.Unmarshal(comment.User, &commentUser); err != nil {
			return schema.IssueResponse{}, err
		}
		commentList = append(commentList, schema.CommentResponse{
			ID:        comment.ID,
			Author:    commentUser["login"].(string),
			Body:      comment.Body,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	// Populate schema.IssueResponse
	resp := schema.IssueResponse{
		ID:          issue.ID,
		Number:      issue.Number,
		Title:       issue.Title,
		State:       issue.State,
		Locked:      issue.Locked,
		Author:      user["login"].(string),
		Assignees:   assignees,
		Labels:      labels,
		Comments:    issue.Comments,
		CreatedAt:   issue.CreatedAt,
		UpdatedAt:   issue.UpdatedAt,
		ClosedAt:    issue.ClosedAt,
		Milestone:   *issue.Milestone,
		Repository:  issue.Repository,
		Body:        issue.Body,
		CommentList: commentList,
	}

	return resp, nil
}

func convertPullRequestModelToSchema(pr model.PullRequest, reviews []model.Review, reviewComments []model.ReviewComment) (schema.PullRequestResponse, error) {
	var user map[string]interface{}
	var labels []string
	var assignees []string
	var head map[string]interface{}
	var base map[string]interface{}

	// Unmarshal JSON fields
	if err := json.Unmarshal(pr.User, &user); err != nil {
		return schema.PullRequestResponse{}, err
	}
	if err := json.Unmarshal(pr.Labels, &labels); err != nil {
		return schema.PullRequestResponse{}, err
	}
	if err := json.Unmarshal([]byte(pr.Assignees), &assignees); err != nil {
		return schema.PullRequestResponse{}, err
	}
	if err := json.Unmarshal(pr.Head, &head); err != nil {
		return schema.PullRequestResponse{}, err
	}
	if err := json.Unmarshal(pr.Base, &base); err != nil {
		return schema.PullRequestResponse{}, err
	}

	// Convert reviews
	var reviewList []schema.ReviewResponse
	for _, review := range reviews {
		var reviewUser map[string]interface{}
		if err := json.Unmarshal(review.User, &reviewUser); err != nil {
			return schema.PullRequestResponse{}, err
		}
		reviewList = append(reviewList, schema.ReviewResponse{
			ID:        review.ID,
			Author:    reviewUser["login"].(string),
			Body:      review.Body,
			State:     review.State,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		})
	}

	// Convert review comments
	var reviewCommentList []schema.ReviewCommentResponse
	for _, reviewComment := range reviewComments {
		var reviewCommentUser map[string]interface{}
		if err := json.Unmarshal(reviewComment.User, &reviewCommentUser); err != nil {
			return schema.PullRequestResponse{}, err
		}
		reviewCommentList = append(reviewCommentList, schema.ReviewCommentResponse{
			ID:        reviewComment.ID,
			Author:    reviewCommentUser["login"].(string),
			Body:      reviewComment.Body,
			Path:      reviewComment.Path,
			Position:  reviewComment.Position,
			CreatedAt: reviewComment.CreatedAt,
			UpdatedAt: reviewComment.UpdatedAt,
		})
	}

	// Populate schema.PullRequestResponse
	resp := schema.PullRequestResponse{
		ID:                pr.ID,
		Number:            pr.Number,
		Title:             pr.Title,
		State:             pr.State,
		Locked:            pr.Locked,
		Draft:             pr.Draft,
		Author:            user["login"].(string),
		Assignees:         assignees,
		Labels:            labels,
		CreatedAt:         pr.CreatedAt,
		UpdatedAt:         pr.UpdatedAt,
		ClosedAt:          pr.ClosedAt,
		MergedAt:          pr.MergedAt,
		Milestone:         *pr.Milestone,
		Repository:        pr.Repository,
		Body:              pr.Body,
		MergeCommitSHA:    pr.MergeCommitSHA,
		HeadBranch:        head["ref"].(string),
		BaseBranch:        base["ref"].(string),
		ReviewList:        reviewList,
		ReviewCommentList: reviewCommentList,
	}

	return resp, nil
}
