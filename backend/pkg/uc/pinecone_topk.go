package uc

import (
	"context"
	"fmt"
	"github.com/gochipon/DIS24-Team-C/backend/pkg/config"
	"github.com/pinecone-io/go-pinecone/pinecone"
	"github.com/sashabaranov/go-openai"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

type SearchResult struct {
	Repository string  `json:"repository"`
	StreamName string  `json:"name"`
	ID         string  `json:"id"`
	Score      float32 `json:"score"`
}
type PineconeTopKUseCase struct {
}

func NewPineconeTopKUseCase() *PineconeTopKUseCase {
	return &PineconeTopKUseCase{}
}

func (uc PineconeTopKUseCase) Exec(repository, query string) ([]SearchResult, error) {
	ctx := context.Background()
	pc, err := pinecone.NewClient(pinecone.NewClientParams{
		ApiKey: config.C.PineconeKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Client: %v", err)
	}
	idx, err := pc.DescribeIndex(ctx, config.C.PineconeIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to describe Index: %v", err)
	}

	idxConn, err := pc.Index(pinecone.NewIndexConnParams{Host: idx.Host})
	if err != nil {
		return nil, fmt.Errorf("failed to create Index: %v", err)
	}

	// Configure OpenAI client
	openAIClient := openai.NewClient(config.C.OpenaiKey)

	// Get embeddings from OpenAI
	embeddingResp, err := openAIClient.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
		Model: openai.AdaEmbeddingV2,
		Input: []string{query},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get embeddings: %v", err)
	}
	embeddings := embeddingResp.Data[0].Embedding

	// Execute similarity search
	metadataMap := map[string]interface{}{
		"repository": repository,
	}
	newStruct, err := structpb.NewStruct(metadataMap)
	if err != nil {
		return nil, fmt.Errorf("failed to create metadata struct: %v", err)
	}
	searchResults, err := idxConn.QueryByVectorValues(
		ctx,
		&pinecone.QueryByVectorValuesRequest{
			Vector:          embeddings,
			TopK:            5,
			MetadataFilter:  newStruct,
			IncludeMetadata: true,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query by vector values: %v", err)
	}

	// Process search results
	var topKID []SearchResult

	for _, result := range searchResults.Matches {

		meta := result.Vector.Metadata.AsMap()
		stream := meta["_ab_stream"].(string)
		info := SearchResult{
			Repository: meta["repository"].(string),
			Score:      result.Score,
		}

		switch stream {
		case "releases":
			info.StreamName = "release"
			info.ID = meta["name"].(string)
		case "comments":
			info.StreamName = "issue"
			info.ID = getLastPartOfURL(meta["issue_url"].(string))
		case "issues":
			info.StreamName = "issue"
			info.ID = fmt.Sprintf("%v", meta["number"])
		case "pull_requests":
			info.StreamName = "pull"
			info.ID = fmt.Sprintf("%v", meta["number"])
		case "review_comments":
			info.StreamName = "pull"
			info.ID = getLastPartOfURL(meta["pull_request_url"].(string))
		default:
			fmt.Println("Unknown stream:", stream)
			continue
		}
		topKID = append(topKID, info)
	}
	return topKID, nil
}

func getLastPartOfURL(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
