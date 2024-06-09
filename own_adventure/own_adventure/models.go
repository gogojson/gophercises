package ownadventure

type Book map[string]Story

type Story struct {
	Title   string   `json:"title"`
	Stories []string `json:"stories"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
