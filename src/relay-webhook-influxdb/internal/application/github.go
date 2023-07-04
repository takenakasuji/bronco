package application

import (
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/takenakasuji/bronco/src/relay-webhook-github/internal/repository"
)

type githubApplicationService struct {
	githubRepository repository.GithubRepository
}

type GithubApplicationService interface {
	WriteEvent(event *write.Point) error
}

func NewGithubApplicationService(repository repository.GithubRepository) GithubApplicationService {
	return &githubApplicationService{
		githubRepository: repository,
	}
}

func (app *githubApplicationService) WriteEvent(event *write.Point) error {
	return app.githubRepository.Write(event)
}
