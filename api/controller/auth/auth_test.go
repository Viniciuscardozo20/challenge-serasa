package auth

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCreateJWTToken(t *testing.T) {
	g := NewGomegaWithT(t)
	token, err := CreateToken("51537476467", "secretkeytest")
	g.Expect(err).Should(BeNil())
	g.Expect(string(token)).ShouldNot(BeNil())
}
