package endpoint

import (
	"context"
	"time"

	"github.com/fazpass/goliath/v3/helper/validator"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/entity"
	_error "github.com/tukangremot/bunchat/backend/src/modules/channel/error"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/usecase"
)

type (
	CreateChannelRequest struct {
		Name string `json:"name" validate:"required"`
	}

	CreateChannelResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)

func (endpoint *Endpoint) CreateChannel(ctx context.Context, reqData *CreateChannelRequest) (interface{}, error) {
	var (
		now       = time.Now()
		validates = validator.Validate(reqData)
	)

	if validates != nil {
		endpoint.usecase.LogError(ctx, usecase.MethodCreateChannel, _error.ErrPayloadValidation, time.Since(now), map[string]interface{}{
			"req_data": reqData,
			"errors":   validates,
		})

		return validates, _error.ErrPayloadValidation
	}

	channel := &entity.Channel{
		Name: reqData.Name,
	}

	_, err := endpoint.usecase.CreateChannel(ctx, channel)
	if err != nil {
		endpoint.usecase.LogError(ctx, usecase.MethodCreateChannel, err, time.Since(now), map[string]interface{}{
			"req_data": reqData,
		})

		return nil, err
	}

	endpoint.usecase.LogSuccess(ctx, usecase.MethodCreateChannel, "success", time.Since(now), map[string]interface{}{
		"req_data": reqData,
	})

	return CreateChannelResponse{
		ID:   channel.ID,
		Name: channel.Name,
	}, nil
}
