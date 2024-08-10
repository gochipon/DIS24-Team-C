package main

import (
	"context"
	"fmt"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v55/github"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"net/http"
	"os"
)

var ghClient *github.Client
var gptUC *IssueCheck

func main() {
	godotenv.Load()
	openaiClient := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	gptUC = NewIssueCheck(*openaiClient)
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
	switch e := event.(type) {
	case *github.IssuesEvent:
		handleIssuesEvent(ctx, ghClient, e)
	case *github.PullRequestEvent:
		handlePullRequestEvent(ctx, ghClient, e)
	default:
		log.Printf("Unhandled event type: %s", github.WebHookType(r))
	}
	//response success
	w.WriteHeader(http.StatusOK)
}

func handleIssuesEvent(ctx context.Context, client *github.Client, event *github.IssuesEvent) {
	flag := gptUC.CheckIssue(*event.Issue.Title, *event.Issue.Body)
	var comment *github.IssueComment
	var label string
	switch flag {
	case ResultOK:
		label = "Ready"
		comment = &github.IssueComment{Body: github.String("Issueを投稿頂きありがとうございます!\n**Issueの内容基準を満たしました🥳**\nMaintainerからの返信をお待ちください")}
	case ResultNG:
		label = "Needs Clarification"
		feedback, err := gptUC.ProvideFeedback(*event.Issue.Title, *event.Issue.Body)
		if err != nil {
			log.Printf("Error providing feedback: %v", err)
			feedback = "エラーが発生しました。"
		}
		comment = &github.IssueComment{Body: github.String("Issueを投稿頂きありがとうございます!\n**Issueに不明瞭な点があるようです??**\nIssueの内容を見直して明瞭にしてください🙇\n\n" +
			"以下、自動生成されたアドバイスです。\n```\n" + feedback + "\n```\n")}
	case ResultSpam:
		label = "Spam"
		comment = &github.IssueComment{Body: github.String("Issueを投稿頂きありがとうございます!\n**Issueがスパムと判断されました🚫**\n" +
			"# 何も編集がない場合、1時間後にこのIssueは自動的にCloseされます。\n\nスパムでない場合は、再度説明文を修正してください。")}
	}
	switch *event.Action {
	case "opened":
		_, _, err := client.Issues.CreateComment(ctx, *event.Repo.Owner.Login, *event.Repo.Name, *event.Issue.Number, comment)
		if err != nil {
			log.Printf("Error creating comment: %v", err)
		}
		client.Issues.ReplaceLabelsForIssue(ctx, *event.Repo.Owner.Login, *event.Repo.Name, *event.Issue.Number, []string{label})
	case "edited":
		comments, _, err := client.Issues.ListComments(ctx, *event.Repo.Owner.Login, *event.Repo.Name, *event.Issue.Number, nil)
		if err != nil {
			log.Printf("Error listing comments: %v", err)
			return
		}
		client.Issues.ReplaceLabelsForIssue(ctx, *event.Repo.Owner.Login, *event.Repo.Name, *event.Issue.Number, []string{label})
		for _, cmt := range comments {
			fmt.Printf("%+v\n", cmt)
			if cmt == nil {
				continue
			}
			if cmt.User == nil {
				fmt.Println("cmt.User is nil")
				fmt.Println(cmt)
				continue
			}
			if *cmt.User.Type == "Bot" {
				_, _, err := client.Issues.EditComment(ctx, *event.Repo.Owner.Login, *event.Repo.Name, *cmt.ID, comment)
				if err != nil {
					log.Printf("Error editing cmt: %v", err)
				}
				break
			}
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
	itr, err := ghinstallation.NewKeyFromFile(tr, 966655, 53652978, "/app/credential/secret.pem")
	if err != nil {
		log.Fatal(err)
	}

	ghClient = github.NewClient(&http.Client{Transport: itr})
}
