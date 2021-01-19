package database

import (
	. "challenge-serasa/api/helper_tests/h_database"
	"challenge-serasa/api/mainframe"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"go.mongodb.org/mongo-driver/bson"
)

const testCollection = "dummy-collection"

func TestCollection(t *testing.T) {
	g := NewGomegaWithT(t)
	client := MockClient(g)
	err := client.Database(DBNameTest).Drop(nil)
	g.Expect(err).ShouldNot(HaveOccurred())
	db, err := NewDatabase(FakeDbConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(db).ShouldNot(BeNil())

	coll := MockCollection(g, testCollection)

	t.Run("validate save negativations", func(t *testing.T) {
		intfColl, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		coll.DeleteMany(nil, bson.M{})
		err = intfColl.SaveDocuments([]mainframe.Negativation{fakeNegativation("d6628a0e-d4dd-4f14-8591-2ddc7f1bbeff")})
		g.Expect(err).ShouldNot(HaveOccurred())
	})

	t.Run("validate save and read negativation", func(t *testing.T) {
		intfColl, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		coll.DeleteMany(nil, bson.M{})
		err = intfColl.SaveDocuments([]mainframe.Negativation{fakeNegativation("d6628a0e-d4dd-4f14-8591-2ddc7f1bbeff")})
		g.Expect(err).ShouldNot(HaveOccurred())
		documents, err := intfColl.GetDocuments("25124543043", "customerDocument")
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(&documents[0]).Should(PointTo(MatchAllFields(Fields{
			"CompanyDocument":  BeEquivalentTo("70170935000100"),
			"CompanyName":      BeEquivalentTo("ASD S.A."),
			"CustomerDocument": BeEquivalentTo("25124543043"),
			"Value":            BeEquivalentTo(10340.67),
			"Contract":         BeEquivalentTo("d6628a0e-d4dd-4f14-8591-2ddc7f1bbeff"),
			"DebtDate":         BeEquivalentTo(time.Date(2015, 07, 9, 20, 32, 51, 00, time.UTC)),
			"InclusionDate":    BeEquivalentTo(time.Date(2020, 07, 9, 20, 32, 51, 00, time.UTC)),
		})))
	})

	t.Run("validate save many negativations", func(t *testing.T) {
		intfColl, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		coll.DeleteMany(nil, bson.M{})
		negativations := make([]mainframe.Negativation, 0)
		negativations = append(negativations, fakeNegativation("d6628a0e-d4dd-4f14-8591-2ddc7f1bbeff"))
		negativations = append(negativations, fakeNegativation("5f206825-3cfe-412f-8302-cc1b24a179b0"))
		err = intfColl.SaveDocuments(negativations)
		g.Expect(err).ShouldNot(HaveOccurred())
		documents, err := intfColl.GetDocuments("25124543043", "customerDocument")
		g.Expect(err).ShouldNot(HaveOccurred())
		for _, d := range documents {
			if "d6628a0e-d4dd-4f14-8591-2ddc7f1bbeff" == d.Contract {
				g.Expect(&d).Should(PointTo(MatchAllFields(Fields{
					"CompanyDocument":  BeEquivalentTo("70170935000100"),
					"CompanyName":      BeEquivalentTo("ASD S.A."),
					"CustomerDocument": BeEquivalentTo("25124543043"),
					"Value":            BeEquivalentTo(10340.67),
					"Contract":         BeEquivalentTo("d6628a0e-d4dd-4f14-8591-2ddc7f1bbeff"),
					"DebtDate":         BeEquivalentTo(time.Date(2015, 07, 9, 20, 32, 51, 00, time.UTC)),
					"InclusionDate":    BeEquivalentTo(time.Date(2020, 07, 9, 20, 32, 51, 00, time.UTC)),
				})))
			} else {
				g.Expect(&d).Should(PointTo(MatchAllFields(Fields{
					"CompanyDocument":  BeEquivalentTo("70170935000100"),
					"CompanyName":      BeEquivalentTo("ASD S.A."),
					"CustomerDocument": BeEquivalentTo("25124543043"),
					"Value":            BeEquivalentTo(10340.67),
					"Contract":         BeEquivalentTo("5f206825-3cfe-412f-8302-cc1b24a179b0"),
					"DebtDate":         BeEquivalentTo(time.Date(2015, 07, 9, 20, 32, 51, 00, time.UTC)),
					"InclusionDate":    BeEquivalentTo(time.Date(2020, 07, 9, 20, 32, 51, 00, time.UTC)),
				})))
			}
		}
	})

	t.Run("validate read nonexistent document", func(t *testing.T) {
		intfColl, err := db.Collection(testCollection)
		g.Expect(err).ShouldNot(HaveOccurred())
		coll.DeleteMany(nil, bson.M{})
		err = intfColl.SaveDocuments([]mainframe.Negativation{fakeNegativation("5f206825-3cfe-412f-8302-cc1b24a179b0")})
		g.Expect(err).ShouldNot(HaveOccurred())
		nav, err := intfColl.GetDocuments("26658236674", "customerDocument")
		g.Expect(nav).Should(BeEmpty())
		g.Expect(err).Should(BeNil())
	})

}

func fakeNegativation(contract string) mainframe.Negativation {
	negativation := mainframe.GenerateNegativation(
		"70170935000100",
		"ASD S.A.",
		"25124543043",
		10340.67,
		contract,
		time.Date(2015, 07, 9, 20, 32, 51, 00, time.UTC),
		time.Date(2020, 07, 9, 20, 32, 51, 00, time.UTC),
	)
	return *negativation
}
