package model

type Discord struct {
	Username  string   `json:"username"`
	AvatarURL string   `json:"avatar_url"`
	Content   string   `json:"content"`
	Embeds    []Embeds `json:"embeds"`
}

type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Embeds struct {
	Title       string  `json:"title"`
	URL         string  `json:"url"`
	Description string  `json:"description"`
	Color       int     `json:"color"`
	Fields      []Field `json:"fields"`
	Thumbnail   struct {
		URL string `json:"url"`
	} `json:"thumbnail"`
	Footer struct {
		Text    string `json:"text"`
		IconURL string `json:"icon_url"`
	} `json:"footer"`
}