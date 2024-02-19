package app

import (
	"context"
	"time"

	"github.com/fazpass/goliath/v3/database"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(ctx context.Context) *mongo.Client {

	var mongoClient, err = database.Init(ctx, database.Config{
		Driver: "mongodb",
		Source: viper.GetString("MONGO_URI"),
		MongoOptions: options.Client().
			ApplyURI(viper.GetString("MONGO_URI")).
			SetMinPoolSize(viper.GetUint64("MONGO_POOL_MIN")).
			SetMaxPoolSize(viper.GetUint64("MONGO_POOL_MAX")).
			SetMaxConnIdleTime(time.Duration(viper.GetUint64("MONGO_MAX_IDLE_TIME_SECOND")) * time.Second),
	})

	if err != nil {
		panic(err)
	}

	return mongoClient.(*mongo.Client)
}
