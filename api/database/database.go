package database

import (
	"fmt"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type driver struct {
	db *mongo.Database
}

func NewDatabase(config Config) (Database, error) {
	client, err := mongo.Connect(nil, options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d/?connect=direct", config.Host, config.Port)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init database")
	}
	err = client.Ping(nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "fail no ping")
	}
	return &driver{db: client.Database(config.Database)}, nil
}

func (d *driver) Collection(name string) (Collection, error) {
	return newCollection(name, *d.db)
}
