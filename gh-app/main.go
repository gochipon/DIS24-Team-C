package main

import (
	"context"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v55/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"os"
)

var ghClient *github.Client

func main() {
	godotenv.Load()

	http.HandleFunc("/webhook", handleWebhook)
	log.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("GH_APP_CLIENT_SECRET")
	if token == "" {
		http.Error(w, "GH_APP_CLIENT_SECRET is not set", http.StatusInternalServerError)
		return
	}
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("Failed to parse webhook: %v", err)
		http.Error(w, "Failed to parse webhook", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	switch e := event.(type) {
	case *github.IssuesEvent:
		handleIssuesEvent(ctx, client, e)
	case *github.PullRequestEvent:
		handlePullRequestEvent(ctx, client, e)
	default:
		log.Printf("Unhandled event type: %s", github.WebHookType(r))
	}
	//response success
	w.WriteHeader(http.StatusOK)
}

func handleIssuesEvent(ctx context.Context, client *github.Client, event *github.IssuesEvent) {
	switch *event.Action {
	case "opened":
		comment := &github.IssueComment{Body: github.String("ありがとう!")}
		_, _, err := client.Issues.CreateComment(ctx, *event.Repo.Owner.Login, *event.Repo.Name, *event.Issue.Number, comment)
		if err != nil {
			log.Printf("Error creating comment: %v", err)
		}

	case "edited":
		comments, _, err := client.Issues.ListComments(ctx, *event.Repo.Owner.Login, *event.Repo.Name, *event.Issue.Number, nil)
		if err != nil {
			log.Printf("Error listing comments: %v", err)
			return
		}

		for _, comment := range comments {
			if *comment.User.Name == "OSSAssistant" {
				newBody := "編集内容を確認しました!"
				_, _, err := client.Issues.EditComment(ctx, *event.Repo.Owner.Login, *event.Repo.Name, *comment.ID, &github.IssueComment{Body: &newBody})
				if err != nil {
					log.Printf("Error editing comment: %v", err)
				}
				break
			}
			log.Printf("%+v", comment)
		}
	}
}

func handlePullRequestEvent(ctx context.Context, client *github.Client, event *github.PullRequestEvent) {
	if *event.Action == "opened" {
		comment := &github.IssueComment{Body: github.String("新しいPRを確認しました")}
		_, _, err := client.Issues.CreateComment(ctx, *event.Repo.Owner.Login, *event.Repo.Name, *event.PullRequest.Number, comment)
		if err != nil {
			log.Printf("Error creating comment: %v", err)
		}
	}
}

func init() {
	tr := http.DefaultTransport
	itr, err := ghinstallation.NewKeyFromFile(tr, 966655, 53652978, "./secret.pem")
	if err != nil {
		log.Fatal(err)
	}

	ghClient = github.NewClient(&http.Client{Transport: itr})
}
