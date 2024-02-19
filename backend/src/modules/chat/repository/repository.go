package repository

import (
	"github.com/tukangremot/bunchat/backend/src/app"
)

type (
	RepositoryInterface interface {
	}

	Repository struct {
		app *app.App
		// db     *app.DB
		// logger *logrus.Logger
	}
)

func Init(app *app.App) *Repository {
	return &Repository{
		app: app,
		// db:     app.GetDB(),
		// logger: app.GetLogger(),
	}
}

// func (r *Repository) BeginTx(ctx context.Context) (*sql.Tx, error) {
// 	tx, err := r.db.Master.BeginTx(ctx, nil)
// 	if err != nil {

// 		return nil, _error.ErrDB
// 	}

// 	return tx, nil
// }
