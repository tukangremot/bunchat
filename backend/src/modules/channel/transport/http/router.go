package http

import (
	"github.com/go-chi/chi"

	"github.com/tukangremot/bunchat/backend/src/app"

	"github.com/jabardigitalservice/golog/http/middleware"
	gologlogger "github.com/jabardigitalservice/golog/logger"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/constant"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/endpoint"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/transport/http/handler"
)

func Init(app *app.App, endpoint endpoint.EndpointInterface) *chi.Mux {
	var (
		router = chi.NewRouter()
		h      = handler.InitHandler(app, endpoint)
	)

	router.Use(middleware.Logger(app.GetLogger(), &gologlogger.LoggerData{
		Service: constant.ServiceName,
		Module:  constant.ModuleName,
		Version: app.GetVersion(),
	}, false))

	router.Post("/", h.CreateChannel)

	return router
}
