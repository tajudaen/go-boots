package services

import (
	"strings"

	"github.com/tajud99n/go-micro/src/api/config"
	"github.com/tajud99n/go-micro/src/api/domain/github"
	"github.com/tajud99n/go-micro/src/api/domain/repositories"
	providers "github.com/tajud99n/go-micro/src/api/providers/github"
	"github.com/tajud99n/go-micro/src/api/utils/errors"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := providers.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
