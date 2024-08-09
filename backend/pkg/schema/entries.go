package schema

type SearchEntry struct {
	Type    string      `json:"type"`
	Score   float32     `json:"score"`
	Content interface{} `json:"content"`
}
