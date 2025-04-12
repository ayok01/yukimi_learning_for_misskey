package model

type User struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Host     string `json:"host"`
}
