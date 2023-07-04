package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/takenakasuji/bronco/src/relay-webhook-github/dto"
	"github.com/takenakasuji/bronco/src/relay-webhook-github/internal/application"
	model "github.com/takenakasuji/bronco/src/relay-webhook-github/internal/model/github"
)

//TODO: secretの実装

func Github(githubApp application.GithubApplicationService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		eventType := c.Get("X-Github-Event")
		if eventType == "" {
			return fiber.NewError(fiber.StatusBadRequest, "X-Github-Event has no value")
		}
		e := githubEventRouter(eventType, c)
		return c.JSON(githubApp.WriteEvent(e))
	}
}

func githubEventRouter(eventType string, c *fiber.Ctx) *write.Point {
	switch eventType {
	case "pull_request":
		e := model.PullRequestEvent{}
		_ = c.BodyParser(&e)
		m := dto.PullRequestMetric{}
		return m.NewMetric(e)
	}
	return nil
}
