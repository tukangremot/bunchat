package usecase

import (
	"context"

	"github.com/tukangremot/bunchat/backend/src/modules/channel/entity"
)

func (usecase *UseCase) CreateChannel(ctx context.Context, data *entity.Channel) (*entity.Channel, error) {
	_, err := usecase.repo.CreateChannel(ctx, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
