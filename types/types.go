package types

type Request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type Content struct {
	Type     string   `json:"type"`
	Text     string   `json:"text"`
	ImageURL ImageURL `json:"-"`
}

type ImageURL struct {
	URL string `json:"url"`
}
