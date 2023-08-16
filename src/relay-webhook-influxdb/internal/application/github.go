package application

import (
	"fmt"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/takenakasuji/bronco/src/relay-webhook-github/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type githubApplicationService struct {
	githubRepository repository.GithubRepository
	rawRepository    repository.RawRepository
}

type GithubApplicationService interface {
	WriteEvent(event *write.Point, rawBody []byte) error
}

func NewGithubApplicationService(gr repository.GithubRepository, rr repository.RawRepository) GithubApplicationService {
	return &githubApplicationService{
		githubRepository: gr,
		rawRepository:    rr,
	}
}

func (app *githubApplicationService) WriteEvent(event *write.Point, rawBody []byte) error {
	if err := app.githubRepository.Write(event); err != nil {
		return err
	}

	bsonMap := bson.M{}
	err := bson.UnmarshalExtJSON(rawBody, false, &bsonMap)
	fmt.Println(bsonMap)
	if err != nil {
		return err
	}
	if err := app.rawRepository.Insert(bsonMap); err != nil {
		return err
	}
	return nil
}
