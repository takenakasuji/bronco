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

type PushMetric struct{}

func (e PushMetric) NewMetric(s model.PushEvent) *write.Point {
	return influxdb2.NewPoint("push",
		map[string]string{
			"event":      "push",
			"repository": s.Repository.Repository,
			"private":    fmt.Sprintf("%v", s.Repository.Private),
			"user":       s.Sender.User,
			"admin":      fmt.Sprintf("%v", s.Sender.Admin),
		},
		map[string]interface{}{
			"stars":  s.Repository.Stars,
			"forks":  s.Repository.Forks,
			"issues": s.Repository.Issues,
			"ref":    s.Ref,
			"before": s.Before,
			"after":  s.After,
		},
		time.Now(),
	)
}

type ReleaseMetric struct{}

func (e ReleaseMetric) NewMetric(s model.ReleaseEvent) *write.Point {
	return influxdb2.NewPoint("release",
		map[string]string{
			"event":      "release",
			"repository": s.Repository.Repository,
			"private":    fmt.Sprintf("%v", s.Repository.Private),
			"user":       s.Sender.User,
			"admin":      fmt.Sprintf("%v", s.Sender.Admin),
		},
		map[string]interface{}{
			"stars":   s.Repository.Stars,
			"forks":   s.Repository.Forks,
			"issues":  s.Repository.Issues,
			"tagName": s.Release.TagName,
		},
		time.Now(),
	)
}

type CreateMetric struct{}

func (e CreateMetric) NewMetric(s model.CreateEvent) *write.Point {
	return influxdb2.NewPoint("create",
		map[string]string{
			"event":      "create",
			"repository": s.Repository.Repository,
			"private":    fmt.Sprintf("%v", s.Repository.Private),
			"user":       s.Sender.User,
			"admin":      fmt.Sprintf("%v", s.Sender.Admin),
		},
		map[string]interface{}{
			"stars":   s.Repository.Stars,
			"forks":   s.Repository.Forks,
			"issues":  s.Repository.Issues,
			"ref":     s.Ref,
			"refType": s.RefType,
		},
		time.Now(),
	)
}

type DeleteMetric struct{}

func (e DeleteMetric) NewMetric(s model.DeleteEvent) *write.Point {
	return influxdb2.NewPoint("delete",
		map[string]string{
			"event":      "delete",
			"repository": s.Repository.Repository,
			"private":    fmt.Sprintf("%v", s.Repository.Private),
			"user":       s.Sender.User,
			"admin":      fmt.Sprintf("%v", s.Sender.Admin),
		},
		map[string]interface{}{
			"stars":   s.Repository.Stars,
			"forks":   s.Repository.Forks,
			"issues":  s.Repository.Issues,
			"ref":     s.Ref,
			"refType": s.RefType,
		},
		time.Now(),
	)
}
