package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/geoirb/event/pkg/model"
	"github.com/geoirb/event/pkg/service"
)

type builder interface {
	NewEvent(e model.Event) Event
	UpdateEvent(e model.Event) (filter bson.M, update bson.M)
}

type Mongo struct {
	collection *mongo.Collection
	builder    builder
}

func NewMongo(
	ctx context.Context,
	connStr, databaseName, collectionName string,
	builder builder,
) (*Mongo, error) {
	opts := options.Client().ApplyURI(connStr)
	connect, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	if err = connect.Ping(ctx, nil); err != nil {
		err = fmt.Errorf("error ping mongo storage %w", err)
	}

	collection := connect.Database(databaseName).Collection(collectionName)
	if _, err := collection.Indexes().CreateOne(ctx,
		mongo.IndexModel{
			Keys:    bson.M{"type": 1},
			Options: options.Index().SetUnique(true),
		}); err != nil {
		return nil, err
	}

	return &Mongo{
		collection: collection,
		builder:    builder,
	}, err
}

func (m *Mongo) Create(ctx context.Context, e model.Event) error {
	event := m.builder.NewEvent(e)
	_, err := m.collection.InsertOne(ctx, event)
	if err == mongo.ErrInvalidIndexValue {
		err = nil
	}
	return err
}

func (m *Mongo) Update(ctx context.Context, e model.Event) error {
	filter, update := m.builder.UpdateEvent(e)
	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err == mongo.ErrNoDocuments {
		// HOWDO
		err = service.ErrNotFound
	}
	return err
}
