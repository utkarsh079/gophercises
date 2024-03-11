package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func JsonStory(f io.Reader) (Story, error) {
	var story Story
	d := json.NewDecoder(f)
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
