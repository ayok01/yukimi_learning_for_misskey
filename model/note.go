package model

type Note struct {
	ID         string         `json:"id"`
	Text       string         `json:"text"`
	Visibility string         `json:"visibility"`
	LocalOnly  bool           `json:"localOnly"`
	User       User           `json:"user"`
	Reactions  map[string]int `json:"reactions"`
	MyReaction string         `json:"myReaction"`
}
