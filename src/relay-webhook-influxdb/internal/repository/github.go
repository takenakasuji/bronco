package repository

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type GithubRepository interface {
	Write(event *write.Point) error
}

type githubRepository struct {
	writeAPI api.WriteAPIBlocking
}

func NewGithubRepository(client influxdb2.Client) GithubRepository {
	writeAPI := client.WriteAPIBlocking("my-org", "github")
	return &githubRepository{
		writeAPI: writeAPI,
	}
}

func (gr *githubRepository) Write(event *write.Point) error {
	err := gr.writeAPI.WritePoint(context.Background(), event)
	if err != nil {
		return err
	}
	return err
}
