package handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/schema"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/uc"
	"log"
	"strconv"
	"sync"
)

type PineconeSearchHandler struct {
	pcUC      *uc.PineconeTopKUseCase
	issueUC   *uc.QueryIssueUC
	pullUC    *uc.QueryPullUC
	summaryUC *uc.SummarizeUseCase
}

func NewPineconeSearchHandler(db *sql.DB) *PineconeSearchHandler {
	return &PineconeSearchHandler{
		issueUC:   uc.NewQueryIssueUC(db),
		pullUC:    uc.NewQueryPullUC(db),
		pcUC:      uc.NewPineconeTopKUseCase(),
		summaryUC: uc.NewSummarizeUseCase(),
	}
}

type PineconeSearchRequest struct {
	Query string `json:"query"`
}

func (p *PineconeSearchHandler) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		org := c.Param("org")
		repo := c.Param("repo")
		if org == "" || repo == "" {
			c.JSON(400, gin.H{"error": "org and repo are required"})
			return
		}
		repository := org + "/" + repo
		var req PineconeSearchRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		exec, err := p.pcUC.Exec(repository, req.Query)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		resp := make([]schema.SearchEntry, len(exec))
		wg := sync.WaitGroup{}
		for i, result := range exec {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				var out schema.SearchEntry
				var fullContent string
				switch result.StreamName {
				case "issue":
					iid, _ := strconv.ParseInt(result.ID, 10, 64)
					issue, err := p.issueUC.Exec(repository, iid)
					if err != nil {
						log.Printf("failed to get issue, %s %d: %v", repository, iid, err)
						return
					}
					out.Content = issue
					out.Type = "issue"
					var fullComments string
					for _, comment := range issue.Comments {
						fullComments += fmt.Sprintf("Author: %s\n\nBody: %s\n\n", comment.User, comment.Body)
					}
					fullContent = fmt.Sprintf("Title: %s\n\nBody: %s\n\nComments: %s", issue.Issue.Title, issue.Issue.Body, fullComments)
				case "pull":
					pid, _ := strconv.ParseInt(result.ID, 10, 64)
					pull, err := p.pullUC.Exec(repository, pid)
					if err != nil {
						log.Printf("failed to get pull request, %s %d: %v", repository, pid, err)
						return
					}
					out.Content = pull
					out.Type = "pull"
					var fullComments string
					for _, review := range pull.Reviews {
						fullComments += fmt.Sprintf("---\nAuthor: %s\n\nBody: %s\n\n", review.User, review.Body)
					}
					fullContent = fmt.Sprintf("Title: %s\n\nBody: %s\n\nReviews: %s", pull.PullRequest.Title, pull.PullRequest.Body, fullComments)
				default:
					log.Printf("unknown stream: %s", result.StreamName)
				}
				out.Score = result.Score
				summary, err := p.summaryUC.Summarize(fullContent)
				if err != nil {
					log.Printf("failed to summarize: %v", err)
					return
				}
				out.Summary = summary
				resp[i] = out
			}(i)
		}
		wg.Wait()
		c.JSON(200, resp)
		return
	}
}
