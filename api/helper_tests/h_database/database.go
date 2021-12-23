package helper_tests

import (
	"fmt"

	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBHostTest = "localhost"
const DBNameTest = "Dummy-db"
const DBPassTest = "dummyPass"
const DBUserTest = "dummyUser"
const DBPortTest = 27017

const Negativation = "dummyNegativation-collection"

func MockCollection(g *GomegaWithT, collName string) *mongo.Collection {
	db := MockClient(g).Database(DBNameTest)
	coll := db.Collection(collName)
	return coll
}

func MockClient(g *GomegaWithT) *mongo.Client {
	client, err := mongo.Connect(nil, options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d/?connect=direct", DBHostTest, DBPortTest)),
	)
	g.Expect(err).ToNot(HaveOccurred())
	return client
}
