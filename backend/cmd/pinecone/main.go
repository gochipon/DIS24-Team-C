package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"os"
	"strings"

	"github.com/pinecone-io/go-pinecone/pinecone"
)

type SearchResult struct {
	Repository string  `json:"repository"`
	StreamName string  `json:"name"`
	ID         string  `json:"id"`
	Score      float32 `json:"score"`
}

func main() {
	godotenv.Load()
	openaiKey := os.Getenv("OPENAI_KEY")
	repository := "cli/cli"
	pineconeIndex := os.Getenv("PINECONE_INDEX")
	pineconeAPIKey := os.Getenv("PINECONE_KEY")
	query := "git authentication does not work"
	ctx := context.Background()

	pc, err := pinecone.NewClient(pinecone.NewClientParams{
		ApiKey: pineconeAPIKey,
	})
	if err != nil {
		log.Fatalf("Failed to create Client: %v", err)
	}

	idx, err := pc.DescribeIndex(ctx, pineconeIndex)
	if err != nil {
		log.Fatalf("Failed to describe Index: %v", err)
	}

	idxConn, err := pc.Index(pinecone.NewIndexConnParams{Host: idx.Host})
	if err != nil {
		log.Fatalf("Failed to create Index: %v", err)
	}

	// Configure OpenAI client
	openAIClient := openai.NewClient(openaiKey)

	// Get embeddings from OpenAI
	embeddingResp, err := openAIClient.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
		Model: openai.AdaEmbeddingV2,
		Input: []string{query},
	})
	if err != nil {
		log.Fatalf("Failed to get embeddings: %v", err)
	}
	embeddings := embeddingResp.Data[0].Embedding

	// Execute similarity search
	metadataMap := map[string]interface{}{
		"repository": repository,
	}
	newStruct, err := structpb.NewStruct(metadataMap)
	if err != nil {
		log.Fatalf("Failed to create metadata struct: %v", err)
		return
	}
	searchResults, err := idxConn.QueryByVectorValues(
		ctx,
		&pinecone.QueryByVectorValuesRequest{
			Vector:          embeddings,
			TopK:            4,
			MetadataFilter:  newStruct,
			IncludeMetadata: true,
		},
	)
	if err != nil {
		log.Fatalf("Failed to execute similarity search: %v", err)
	}

	// Process search results
	var topKID []SearchResult

	fmt.Println(searchResults.Matches)
	for _, result := range searchResults.Matches {

		meta := result.Vector.Metadata.AsMap()
		stream := meta["_ab_stream"].(string)
		info := SearchResult{
			Repository: meta["repository"].(string),
			Score:      result.Score,
		}

		switch stream {
		case "releases":
			info.StreamName = "releases"
			info.ID = meta["name"].(string)
		case "comments":
			info.StreamName = "issues"
			info.ID = getLastPartOfURL(meta["issue_url"].(string))
		case "issues":
			info.StreamName = "issues"
			info.ID = fmt.Sprintf("%v", meta["number"])
		case "pull_requests":
			info.StreamName = "pull_requests"
			info.ID = fmt.Sprintf("%v", meta["number"])
		case "review_comments":
			info.StreamName = "pull_requests"
			info.ID = getLastPartOfURL(meta["pull_request_url"].(string))
		default:
			fmt.Println("Unknown stream:", stream)
			continue
		}

		topKID = append(topKID, info)
	}
	fmt.Println(topKID)
}

func getLastPartOfURL(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
