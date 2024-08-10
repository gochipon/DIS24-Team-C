package schema

type SearchEntry struct {
	Type    string      `json:"type"`
	Score   float32     `json:"score"`
	Summary string      `json:"summary"`
	Content interface{} `json:"content"`
}
