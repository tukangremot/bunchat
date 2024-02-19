package endpoint

import (
	"context"

	"github.com/jabardigitalservice/golog/logger"
	"github.com/tukangremot/bunchat/backend/src/app"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/usecase"
)

type (
	EndpointInterface interface {
		CreateChannel(ctx context.Context, reqData *CreateChannelRequest) (interface{}, error)
	}

	Endpoint struct {
		app     *app.App
		usecase usecase.UseCaseInterface
		logger  *logger.Logger
	}
)

func Init(app *app.App, usecase usecase.UseCaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
		logger:  app.GetLogger(),
	}
}
