package model

type Sender struct {
	User  string `json:"login"`
	Admin bool   `json:"site_admin"`
}
