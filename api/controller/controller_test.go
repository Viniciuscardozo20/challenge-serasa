package controller

import (
	"challenge-serasa/api/database"
	. "challenge-serasa/api/helper_tests/h_database"
	"challenge-serasa/api/helper_tests/h_mainframe"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"go.mongodb.org/mongo-driver/bson"
)

const testCollection = "dummy-collection"

func TestController(t *testing.T) {
	g := NewGomegaWithT(t)
	client := MockClient(g)
	coll := MockCollection(g, testCollection)
	server := h_mainframe.MockMainframeServer(g)
	err := client.Database(DBNameTest).Drop(nil)
	g.Expect(err).ShouldNot(HaveOccurred())
	db, err := database.NewDatabase(FakeDbConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(db).ShouldNot(BeNil())
	intfColl, err := db.Collection(testCollection)
	g.Expect(err).ShouldNot(HaveOccurred())
	controller := NewController(intfColl, server.URL)

	t.Run("validate update negativations", func(t *testing.T) {
		coll.DeleteMany(nil, bson.M{})
		err := controller.UpdateNegativations()
		g.Expect(err).ShouldNot(HaveOccurred())
	})

	t.Run("validate update and read negativations", func(t *testing.T) {
		coll.DeleteMany(nil, bson.M{})
		err := controller.UpdateNegativations()
		g.Expect(err).ShouldNot(HaveOccurred())
		navs, err := controller.GetNegativationByCustomer("51537476467")
		g.Expect(err).ShouldNot(HaveOccurred())
		for _, d := range *navs {
			if "bc063153-fb9e-4334-9a6c-0d069a42065b" == d.Contract {
				g.Expect(&d).Should(PointTo(MatchAllFields(Fields{
					"CompanyDocument":  BeEquivalentTo("59291534000167"),
					"CompanyName":      BeEquivalentTo("ABC S.A."),
					"CustomerDocument": BeEquivalentTo("51537476467"),
					"Value":            BeEquivalentTo(1235.23),
					"Contract":         BeEquivalentTo("bc063153-fb9e-4334-9a6c-0d069a42065b"),
					"DebtDate":         BeEquivalentTo(time.Date(2015, 11, 13, 23, 32, 51, 00, time.UTC)),
					"InclusionDate":    BeEquivalentTo(time.Date(2020, 11, 13, 23, 32, 51, 00, time.UTC)),
				})))
			} else {
				g.Expect(&d).Should(PointTo(MatchAllFields(Fields{
					"CompanyDocument":  BeEquivalentTo("77723018000146"),
					"CompanyName":      BeEquivalentTo("123 S.A."),
					"CustomerDocument": BeEquivalentTo("51537476467"),
					"Value":            BeEquivalentTo(400.00),
					"Contract":         BeEquivalentTo("5f206825-3cfe-412f-8302-cc1b24a179b0"),
					"DebtDate":         BeEquivalentTo(time.Date(2015, 10, 12, 23, 32, 51, 00, time.UTC)),
					"InclusionDate":    BeEquivalentTo(time.Date(2020, 10, 12, 23, 32, 51, 00, time.UTC)),
				})))
			}
		}
	})
}

func FakeDbConfig() database.Config {
	return database.Config{
		Host:     DBHostTest,
		Port:     DBPortTest,
		User:     DBUserTest,
		Password: DBPassTest,
		Database: DBNameTest,
	}
}
