package endpoint

import (
	"context"
	"errors"
	"time"

	"github.com/fazpass/goliath/v3/helper/validator"
	"github.com/tukangremot/bunchat/backend/src/modules/chat/usecase"
)

type (
	ClientConnectedRequest struct {
		ChannelID string `json:"channelId" validate:"required"`
	}

	ClientConnectedResponse struct {
	}
)

func (e *Endpoint) ClientConnected(ctx context.Context, data ClientConnectedRequest) (interface{}, error) {
	now := time.Now()

	validates := validator.Validate(data)
	if validates != nil {
		return validates, errors.New("error")
	}

	e.usecase.LogSuccess(ctx, usecase.MethodClientConnected, "connected", time.Since(now), map[string]interface{}{})

	return ClientConnectedResponse{}, nil
}
