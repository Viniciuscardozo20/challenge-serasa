package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type collection struct {
	coll *mongo.Collection
}

func newCollection(name string, db mongo.Database) (*collection, error) {
	names, err := db.ListCollectionNames(nil, bson.D{{}})
	if err != nil {
		return nil, err
	}
	var exists = false
	for _, n := range names {
		if n == name {
			exists = true
		}
	}
	var coll *mongo.Collection
	if exists != true {
		coll = db.Collection(name, nil)
		_, err = coll.Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys:    bson.D{{Key: "customerDocument", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		)
		if err != nil {
			return nil, err
		}
	} else {
		coll = db.Collection(name, nil)
	}
	return &collection{
		coll: coll,
	}, nil
}
