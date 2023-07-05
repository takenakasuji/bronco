package model

type Issue struct {
	Number   int    `json:"number"`
	Title    string `json:"title"`
	Comments int    `json:"comments"`
}
