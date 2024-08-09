package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/schema"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/uc"
	"log"
	"strconv"
)

type PineconeSearchHandler struct {
	pcUC    *uc.PineconeTopKUseCase
	issueUC *uc.QueryIssueUC
	pullUC  *uc.QueryPullUC
}

func NewPineconeSearchHandler(db *sql.DB) *PineconeSearchHandler {
	return &PineconeSearchHandler{
		issueUC: uc.NewQueryIssueUC(db),
		pullUC:  uc.NewQueryPullUC(db),
		pcUC:    uc.NewPineconeTopKUseCase(),
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
		resp := make([]schema.SearchEntry, 0, len(exec))
		for _, result := range exec {
			var out schema.SearchEntry
			switch result.StreamName {
			case "issue":
				iid, _ := strconv.ParseInt(result.ID, 10, 64)
				issue, err := p.issueUC.Exec(repository, iid)
				if err != nil {
					log.Printf("failed to get issue, %s %d: %v", repository, iid, err)
					continue
				}
				out.Content = issue
				out.Type = "issue"
			case "pull":
				pid, _ := strconv.ParseInt(result.ID, 10, 64)
				pull, err := p.pullUC.Exec(repository, pid)
				if err != nil {
					log.Printf("failed to get pull request, %s %d: %v", repository, pid, err)
					continue
				}
				out.Content = pull
				out.Type = "pull"
			default:
				log.Printf("unknown stream: %s", result.StreamName)
			}
			out.Score = result.Score
			resp = append(resp, out)
		}
		c.JSON(200, resp)
		return
	}
}
