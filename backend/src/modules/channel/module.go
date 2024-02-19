package channel

import (
	"github.com/fazpass/goliath/v3/module"
	"github.com/go-chi/chi"
	"github.com/tukangremot/bunchat/backend/src/app"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/endpoint"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/repository"
	transporthttp "github.com/tukangremot/bunchat/backend/src/modules/channel/transport/http"

	"github.com/tukangremot/bunchat/backend/src/modules/channel/usecase"
)

type Module struct {
	usecase    usecase.UseCaseInterface
	endpoint   endpoint.EndpointInterface
	repository *repository.Repository
	httpRouter *chi.Mux
}

func Init(app *app.App) module.ModuleInterface {
	var (
		repository = repository.Init(app)
		usecase    = usecase.Init(app, repository)
		endpoint   = endpoint.Init(app, usecase)
		httpRouter = transporthttp.Init(app, endpoint)
	)

	return &Module{
		repository: repository,
		usecase:    usecase,
		endpoint:   endpoint,
		httpRouter: httpRouter,
	}
}

func (module *Module) GetHttpRouter() *chi.Mux {
	return module.httpRouter
}
