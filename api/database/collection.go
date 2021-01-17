package database

import (
	"challenge-serasa/api/mainframe"
	"context"

	"github.com/pkg/errors"
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
		return nil, errors.Wrap(err, "failed to init collection list")
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
			return nil, errors.Wrap(err, "failed to create a collection")
		}
	} else {
		coll = db.Collection(name, nil)
	}
	return &collection{
		coll: coll,
	}, nil
}

func (c *collection) SaveDocuments(negativations []mainframe.Negativation) error {
	for _, negativation := range negativations {
		result := c.coll.FindOneAndUpdate(nil, bson.M{"customerDocument": negativation.CustomerDocument}, bson.M{"$set": negativation})
		if result.Err() != nil {
			_, err := c.coll.InsertOne(nil, negativation)
			if err != nil {
				return errors.Wrap(err, "failed to insert documents in collection")
			}
		}
	}
	return nil
}

func (c *collection) GetDocument(value interface{}, field string) (*mainframe.Negativation, error) {
	result := c.coll.FindOne(nil, bson.M{field: value})
	if result.Err() != nil {
		return nil, errors.Wrap(result.Err(), "failed to find document")
	}
	var negativation mainframe.Negativation
	err := result.Decode(&negativation)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode document")
	}
	return &negativation, nil
}
