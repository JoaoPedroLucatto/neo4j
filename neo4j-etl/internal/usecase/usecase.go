package usecase

import (
	"context"
	"neo4j-etl/internal/db"

	"github.com/rs/zerolog"
)

type Usecase struct {
	Context    context.Context
	Logger     *zerolog.Logger
	Repository db.Repository
}

func NewUsecaseService(logger *zerolog.Logger,
	repository db.Repository) *Usecase {
	return &Usecase{
		Logger:     logger,
		Repository: repository,
	}
}
