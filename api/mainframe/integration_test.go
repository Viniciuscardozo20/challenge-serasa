package mainframe

import (
	"challenge-serasa/api/helper_tests/mainframe"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

func TestCreateNavigation(t *testing.T) {
	g := NewGomegaWithT(t)

	nav := GenerateNegativation(
		"70170935000100",
		"ASD S.A.",
		"25124543043",
		10340.67,
		"d6628a0e-d4dd-4f14-8591-2ddc7f1bbeff",
		time.Date(2015, 07, 9, 20, 32, 51, 00, time.UTC),
		time.Date(2020, 07, 9, 20, 32, 51, 00, time.UTC),
	)
	g.Expect(nav).Should(PointTo(MatchAllFields(Fields{
		"CompanyDocument":  BeEquivalentTo("70170935000100"),
		"CompanyName":      BeEquivalentTo("ASD S.A."),
		"CustomerDocument": BeEquivalentTo("25124543043"),
		"Value":            BeEquivalentTo(10340.67),
		"Contract":         BeEquivalentTo("d6628a0e-d4dd-4f14-8591-2ddc7f1bbeff"),
		"DebtDate":         BeEquivalentTo(time.Date(2015, 07, 9, 20, 32, 51, 00, time.UTC)),
		"InclusionDate":    BeEquivalentTo(time.Date(2020, 07, 9, 20, 32, 51, 00, time.UTC)),
	})))
}

func TestGetNavigationsFromValidURL(t *testing.T) {
	g := NewGomegaWithT(t)

	server := mainframe.MockMainframeServer(g)

	navigations, err := GetNegativations(server.URL)
	g.Expect(err).Should(BeNil())
	g.Expect(navigations).ShouldNot(BeNil())
}

func TestGetNavigationsFromInvalidURL(t *testing.T) {
	g := NewGomegaWithT(t)

	navigations, err := GetNegativations("http://localhost:303/")
	g.Expect(err).ShouldNot(BeNil())
	g.Expect(navigations).Should(BeNil())

}
