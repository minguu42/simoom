package handler

import (
	"log"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	"github.com/minguu42/simoom/api/domain/auth"
	"github.com/minguu42/simoom/api/domain/model"
	"github.com/minguu42/simoom/api/domain/repository"
	"github.com/minguu42/simoom/api/usecase"
)

var th handler

func TestMain(m *testing.M) {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatalf("failed to create validator: %s", err)
	}
	authn := &auth.AuthenticatorMock{}
	repo := &repository.RepositoryMock{}
	idgen := &model.IDGeneratorMock{}
	th = handler{
		validator:  v,
		auth:       usecase.NewAuth(authn, repo, idgen),
		monitoring: usecase.Monitoring{},
		project:    usecase.NewProject(repo, idgen),
		step:       usecase.NewStep(repo, idgen),
		tag:        usecase.NewTag(repo, idgen),
		task:       usecase.NewTask(repo, idgen),
	}

	m.Run()
}
