package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"github.com/takenakasuji/bronco/src/relay-webhook-github/internal/application"
	"github.com/takenakasuji/bronco/src/relay-webhook-github/internal/handler"
	"github.com/takenakasuji/bronco/src/relay-webhook-github/internal/repository"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	client := influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"))

	githubRepository := repository.NewGithubRepository(client)
	app.Post("/github", handler.Github(application.NewGithubApplicationService(githubRepository)))

	log.Fatal(app.Listen(":3000"))
}
