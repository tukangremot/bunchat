package app

import (
	"context"
	"log"
	"net/http"

	"github.com/fazpass/goliath/v3/config"
	"github.com/fazpass/goliath/v3/router"
	"github.com/go-chi/chi"
	"github.com/jabardigitalservice/golog/logger"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/tukangremot/bunchat/backend/src/constant"
	"go.elastic.co/apm/module/apmhttp"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	App struct {
		router *chi.Mux
		logger *logger.Logger
		mongo  *mongo.Client
		redis  *redis.Client
	}
)

func Init() *App {
	ctx := context.Background()

	err := config.Init()
	if err != nil {
		log.Panic(err)
	}

	redisClient, err := InitRedis(ctx)
	if err != nil {
		log.Panic(err)
	}

	app := &App{
		router: router.InitChi(router.Config{
			Debug: viper.GetBool("APP_DEBUG"),
		}),
		logger: logger.Init(),
		mongo:  InitMongo(ctx),
		redis:  redisClient,
	}

	return app
}

func (app *App) GetHttpRouter() *chi.Mux {
	return app.router
}

func (app *App) GetLogger() *logger.Logger {
	return app.logger
}

func (app *App) GetVersion() string {
	return viper.GetString("APP_VERSION")
}

func (app *App) GetDB() *mongo.Database {
	dbName := viper.GetString("MONGO_DB_NAME")

	return app.mongo.Database(dbName)
}

func (app *App) GetRedis() *redis.Client {
	return app.redis
}

func (app *App) RunHttp() error {
	var port = viper.GetString("APP_PORT")
	if port == "" {
		port = "8181"
	}

	var host = "0.0.0.0:" + port

	app.logger.Info(&logger.LoggerData{
		Category: logger.LoggerApp,
		Service:  constant.ServiceName,
		Method:   "startup",
		Version:  app.GetVersion(),
		AdditionalInfo: map[string]interface{}{
			"host": host,
		},
	}, "running")

	return http.ListenAndServe(host, apmhttp.Wrap(app.router))
}
