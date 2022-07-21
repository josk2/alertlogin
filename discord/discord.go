package discord

type Message struct {
	Content string  `json:"content"`
	Tts     bool    `json:"tts"`
	Embeds  []Embed `json:"embeds"`
}

type Embed struct {
	Title  string       `json:"title"`
	Desc   string       `json:"description"`
	Color  int          `json:"color"`
	Fields []EmbedField `json:"fields"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
