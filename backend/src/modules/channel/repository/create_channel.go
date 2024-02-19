package repository

import (
	"context"

	"github.com/tukangremot/bunchat/backend/src/modules/channel/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ChannelData struct {
		ID   string `bson:"_id"`
		Name string `name:"name"`
	}
)

func (repository *Repository) CreateChannel(ctx context.Context, data *entity.Channel) (*entity.Channel, error) {
	collection := repository.db.Collection(CollectionChannel)

	data.ID = primitive.NewObjectID().Hex()

	_, err := collection.InsertOne(context.TODO(), ChannelData{
		ID:   data.ID,
		Name: data.Name,
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}
