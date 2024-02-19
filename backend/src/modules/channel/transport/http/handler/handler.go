package handler

import (
	"github.com/tukangremot/bunchat/backend/src/app"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/endpoint"
)

type Handler struct {
	app      *app.App
	endpoint endpoint.EndpointInterface
}

func InitHandler(app *app.App, endpoint endpoint.EndpointInterface) *Handler {
	handler := &Handler{
		app:      app,
		endpoint: endpoint,
	}

	return handler
}
