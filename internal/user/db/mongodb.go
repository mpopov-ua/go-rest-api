package db

import (
	"context"
	"github.com/pritunl/mongo-go-driver/mongo"
	"rest-app/internal/user"
	"rest-app/package/logging"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.collection.InsertOne()
}

func (d *db) FindOne(ctx context.Context, id string) (user.User, error) {
	panic("implement me")
}

func (d *db) Update(ctx context.Context, user user.User) error {
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}