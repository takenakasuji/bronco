package model

import "time"

type PullRequest struct {
	Number       int       `json:"number"`
	State        string    `json:"state"`
	Title        string    `json:"title"`
	Comments     int       `json:"comments"`
	Commits      int       `json:"commits"`
	Additions    int       `json:"additions"`
	Deletions    int       `json:"deletions"`
	ChangedFiles int       `json:"changed_files"`
	Head         Head      `json:"head"`
	ClosedAt     time.Time `json:"closed_at"`
}
