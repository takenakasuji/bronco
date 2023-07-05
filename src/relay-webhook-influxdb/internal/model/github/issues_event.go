package model

type IssuesEvent struct {
	Action     string     `json:"action"`
	Issue      Issue      `json:"issue"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}
