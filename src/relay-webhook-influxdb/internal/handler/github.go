package handler

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/takenakasuji/bronco/src/relay-webhook-github/dto"
	"github.com/takenakasuji/bronco/src/relay-webhook-github/internal/application"
	model "github.com/takenakasuji/bronco/src/relay-webhook-github/internal/model/github"
	"github.com/takenakasuji/bronco/src/relay-webhook-github/library"
)

func Github(githubApp application.GithubApplicationService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		httpReq, err := adaptor.ConvertRequest(c, false)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Convert request error")
		}
		if r := library.GithubValidatePayload(httpReq, os.Getenv("GITHUB_APP_SECRET")); r == false {
			return fiber.NewError(fiber.StatusInternalServerError, "Unable to parse secret")
		}
		eventType := c.Get("X-Github-Event")
		if eventType == "" {
			return fiber.NewError(fiber.StatusBadRequest, "X-Github-Event has no value")
		}
		e, err := githubEventRouter(eventType, c)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Request parse error")
		}
		return c.JSON(githubApp.WriteEvent(e))
	}
}

func githubEventRouter(eventType string, c *fiber.Ctx) (*write.Point, error) {
	switch eventType {
	case "pull_request":
		e := model.PullRequestEvent{}
		err := c.BodyParser(&e)
		if err != nil {
			return nil, err
		}
		m := dto.PullRequestMetric{}
		return m.NewMetric(e), nil
	case "issues":
		e := model.IssuesEvent{}
		err := c.BodyParser(&e)
		if err != nil {
			return nil, err
		}
		m := dto.IssuesMetric{}
		return m.NewMetric(e), nil
	case "push":
		e := model.PushEvent{}
		err := c.BodyParser(&e)
		if err != nil {
			return nil, err
		}
		m := dto.PushMetric{}
		return m.NewMetric(e), nil
	case "release":
		e := model.ReleaseEvent{}
		err := c.BodyParser(&e)
		if err != nil {
			return nil, err
		}
		m := dto.ReleaseMetric{}
		return m.NewMetric(e), nil
	case "create":
		e := model.CreateEvent{}
		err := c.BodyParser(&e)
		if err != nil {
			return nil, err
		}
		m := dto.CreateMetric{}
		return m.NewMetric(e), nil
	case "delete":
		e := model.DeleteEvent{}
		err := c.BodyParser(&e)
		if err != nil {
			return nil, err
		}
		m := dto.DeleteMetric{}
		return m.NewMetric(e), nil
	}
	return nil, nil
}
