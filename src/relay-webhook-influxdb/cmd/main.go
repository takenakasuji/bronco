package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/takenakasuji/bronco/src/relay-webhook-github/internal/application"
	"github.com/takenakasuji/bronco/src/relay-webhook-github/internal/handler"
	"github.com/takenakasuji/bronco/src/relay-webhook-github/internal/repository"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	tsClient := influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rawClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@mongodb:27017"))
	err = rawClient.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = rawClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	gr := repository.NewGithubRepository(tsClient)
	rr := repository.NewRawRepository(rawClient)

	app.Post("/github", handler.Github(application.NewGithubApplicationService(gr, rr)))

	log.Fatal(app.Listen(":3000"))
}
