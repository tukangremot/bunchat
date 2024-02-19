package repository

import (
	"context"

	"github.com/tukangremot/bunchat/backend/src/app"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionChannel = "channels"
)

type (
	RepositoryInterface interface {
		CreateChannel(ctx context.Context, data *entity.Channel) (*entity.Channel, error)
	}

	Repository struct {
		app *app.App
		db  *mongo.Database
		// logger *logrus.Logger
	}
)

func Init(app *app.App) *Repository {
	return &Repository{
		app: app,
		db:  app.GetDB(),
		// logger: app.GetLogger(),
	}
}
