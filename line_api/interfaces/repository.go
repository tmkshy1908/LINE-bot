package interfaces

import (
	"context"

	"github.com/tmkshy1908/LINE-bot/domain"
)

type CommonInteractor interface {
	UseCaseSampleRepository(context.Context) ([]*domain.Schedule, error)
}
