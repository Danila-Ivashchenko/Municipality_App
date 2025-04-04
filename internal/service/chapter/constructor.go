package chapter

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	ChapterRepository repository.ChapterRepository
}

type chapterService struct {
	ServiceParams
}

func New(params ServiceParams) service.ChapterService {
	return &chapterService{
		ServiceParams: params,
	}
}
