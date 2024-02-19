package endpoint

import (
	"context"

	"github.com/jabardigitalservice/golog/logger"
	"github.com/tukangremot/bunchat/backend/src/app"
	"github.com/tukangremot/bunchat/backend/src/modules/chat/usecase"
)

type (
	EndpointInterface interface {
		ClientConnected(ctx context.Context, data ClientConnectedRequest) (interface{}, error)
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
