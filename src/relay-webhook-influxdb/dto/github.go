package dto

import (
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	model "github.com/takenakasuji/bronco/src/relay-webhook-github/internal/model/github"
)

type PullRequestMetric struct{}

func (e PullRequestMetric) NewMetric(s model.PullRequestEvent) *write.Point {
	return influxdb2.NewPoint("pull_request",
		map[string]string{
			"event":      "pull_request",
			"action":     s.Action,
			"repository": s.Repository.Repository,
			"haed":       s.PullRequest.Head.Ref,
			"private":    fmt.Sprintf("%v", s.Repository.Private),
			"user":       s.Sender.User,
			"admin":      fmt.Sprintf("%v", s.Sender.Admin),
			"prNumber":   fmt.Sprintf("%v", s.PullRequest.Number),
		},
		map[string]interface{}{
			"stars":        s.Repository.Stars,
			"forks":        s.Repository.Forks,
			"issues":       s.Repository.Issues,
			"state":        s.PullRequest.State,
			"title":        s.PullRequest.Title,
			"comments":     s.PullRequest.Comments,
			"commits":      s.PullRequest.Commits,
			"additions":    s.PullRequest.Additions,
			"deletions":    s.PullRequest.Deletions,
			"changedFiles": s.PullRequest.ChangedFiles,
			"closedAt":     s.PullRequest.ClosedAt,
		},
		time.Now(),
	)
}

type IssuesMetric struct{}

func (e IssuesMetric) NewMetric(s model.IssuesEvent) *write.Point {
	return influxdb2.NewPoint("issues",
		map[string]string{
			"event":      "issues",
			"action":     s.Action,
			"repository": s.Repository.Repository,
			"private":    fmt.Sprintf("%v", s.Repository.Private),
			"user":       s.Sender.User,
			"admin":      fmt.Sprintf("%v", s.Sender.Admin),
			"issue":      fmt.Sprintf("%v", s.Issue.Number),
		},
		map[string]interface{}{
			"stars":    s.Repository.Stars,
			"forks":    s.Repository.Forks,
			"issues":   s.Repository.Issues,
			"title":    s.Issue.Title,
			"comments": s.Issue.Comments,
		},
		time.Now(),
	)
}
